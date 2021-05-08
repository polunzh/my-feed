# My Feed

GO 练习项目。

从 gist 读取订阅 rss 的信息，使用 github actions 定时将更新的 rss 内容发送到邮箱。


## 配置

gist 数据格式(json):

```
[
  {
    "Link": "https://draveness.me/feed.xml"
  },
  {
    "Link": "https://rss.app/feeds/AiZJgzirUsi9Bqws.xml"
  },
]
```

环境变量:

- EMAIL_USERNAME 发送邮箱用户名
- EMAIL_PASSWORD 发送邮箱用户密码
- EMAIL_TO 发送到的邮箱地址
- GIST_SOURCE_FILE gist 文件地址
- GIST_TOKEN Github Personal Access Token, 用来更新 gist
- GIST_ID gist id, 用来更新 gist
