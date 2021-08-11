# hanamaruAPI
simpel api for fetch our [feed](https://datenshi.pw/feed.json) json build by golang

# why
because jekyll json-to-feed is not working or blocked by CORS

# usage
```bash
go get github.com/troke12/hanamaruAPI
go build
./hanamaruAPI
```

# result
```json
{
  "code": 200,
  "posts": [

  {
    "id": "/maintenance-bancho-server/",
    "url": "https://datenshi.pw/maintenance-bancho-server/",
    "title": "Maintenance Bancho Server",
    "content_html": "tonight we will do some live testing of Bancho server around 18.00 to 24.00 WIB (+7)",
    "date_published": "2021-08-09T02:25:48+07:00",
    "category": "osu",
    "post_thumbnail": "https://cdn.discordapp.com/attachments/874251888357441537/874251910624972800/229958147_677421896508962_8511576767945961180_n.png",
    "post_author": "troke",
    "post_author_url": "/authors/troke12/",
    "wordcount": 102
  }]
}
  ```
