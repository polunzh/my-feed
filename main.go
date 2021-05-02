package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/mmcdole/gofeed"
	mail "github.com/xhit/go-simple-mail/v2"
)

const ZONE string = "Asia/Shanghai"
const DATETIME_FORMAT = "2006-01-02 15:04:05 China"

type source struct {
	Link         string `json:"Link"`
	LastItemLink string `json:"LastItemLink"`
}

type article struct {
	Title   string
	Link    string
	Updated string
}

type blog struct {
	Title       string
	Description string
	Link        string
	Updated     string
	Articles    []article
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatTimeString(str string) string {
	if strings.Trim(str, " ") == "" {
		return ""
	}

	t, err := dateparse.ParseAny(str)
	checkError(err)
	location, err := time.LoadLocation(ZONE)

	return t.In(location).Format(DATETIME_FORMAT)
}

func fileExists(filepath string) (bool, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, errors.New(fmt.Sprintf("%s is a directory", filepath))
	}

	return true, nil
}

func checkUpdates(sources *[]source) ([]byte, error) {
	log.Print("checking updates...")

	var blogs []blog
	fp := gofeed.NewParser()
	newSources := make([]source, len(*sources))
	copy(newSources, *sources)

	for index, sourceItem := range *sources {
		log.Printf("checking items: %s...\n", sourceItem.Link)
		feed, err := fp.ParseURL(sourceItem.Link)
		if err != nil {
			return nil, err
		}

		if len(feed.Items) == 0 {
			log.Printf("link %s has no items\n", sourceItem.Link)
			continue
		}

		updatedBlog := blog{Title: feed.Title, Description: feed.Description, Updated: formatTimeString(feed.Updated), Link: feed.Link}

		lastItemLink := strings.Trim(sourceItem.LastItemLink, "")
		newSources[index].LastItemLink = strings.Trim(feed.Items[0].Link, "")

		for _, tempItem := range feed.Items {
			link := strings.Trim(tempItem.Link, "")
			if lastItemLink == link {
				break
			}

			updatedArticle := article{Title: tempItem.Title, Link: link, Updated: formatTimeString(tempItem.Updated)}
			updatedBlog.Articles = append(updatedBlog.Articles, updatedArticle)
		}

		if len(updatedBlog.Articles) > 0 {
			blogs = append(blogs, updatedBlog)
		}
	}

	if len(blogs) == 0 {
		return nil, nil
	}

	sendNotification(&blogs)

	newSourcesBytes, err := json.MarshalIndent(newSources, "", "  ")

	return newSourcesBytes, err
}

func sendNotification(blogs *[]blog) {
	if len(*blogs) == 0 {
		checkError(errors.New("no notifications"))
	}

	sendToEmail(blogs)
}

type emailData struct {
	Blogs *[]blog
}

func sendToEmail(blogs *[]blog) {
	emailData := emailData{Blogs: blogs}
	t := template.Must(template.ParseFiles("email-layout.html"))
	var html bytes.Buffer

	t.Execute(&html, emailData)

	server := mail.NewSMTPClient()
	server.Host = "smtp.163.com"
	server.Port = 465
	server.Username = os.Getenv("EMAIL_USERNAME")
	server.Password = os.Getenv("EMAIL_PASSWORD")
	server.Encryption = mail.EncryptionSSL
	server.KeepAlive = false
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom("博客订阅 <965076377@163.com>").
		AddTo(os.Getenv("EMAIL_TO")).
		SetSubject("最近更新")

	email.SetBody(mail.TextHTML, html.String())
	err = email.Send(smtpClient)
	checkError(err)
}

func readDB(dbFile string) map[string]string {
	db := make(map[string]string)

	return db
}

func downloadFromGist(url string, targetFile string) string {
	exists, _ := fileExists(targetFile)

	if exists == true {
		err := os.Remove(targetFile)
		checkError(err)
	}

	out, err := os.Create(targetFile)
	checkError(err)
	defer out.Close()

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return targetFile
}

func main() {
	cwd, err := os.Getwd()
	checkError(err)

	sourcesPath := filepath.Join(cwd, "sources.json")
	downloadFromGist(os.Getenv("GIST_SOURCE_FILE"), sourcesPath)

	_, err = fileExists(sourcesPath)
	checkError(err)

	sourcesStr, err := ioutil.ReadFile(sourcesPath)
	checkError(err)
	sources := []source{}
	err = json.Unmarshal([]byte(sourcesStr), &sources)
	checkError(err)

	if len(sources) == 0 {
		checkError(errors.New("rss sources is empty"))
	}

	newSourcesStr, err := checkUpdates(&sources)
	checkError(err)

	if newSourcesStr == nil {
		log.Println("no updates")
		os.Exit(0)
	}

	err = ioutil.WriteFile(sourcesPath, newSourcesStr, 0644)
	checkError(err)
	log.Println("rss source file updated")
}
