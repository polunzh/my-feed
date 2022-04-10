package dal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"polunzh/my-feed/ent"

	_ "github.com/mattn/go-sqlite3"
)

func InitClient(name string) (*ent.Client) {
	homedir, _ := os.UserHomeDir()
	dir := path.Join(homedir, ".my-feed")
	os.Mkdir(dir, os.ModePerm)
	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc&cache=shared&_fk=1", path.Join(dir, "db")))
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlite3: %v", err))
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed to create schema: %v", err))
	}

	return client
}

