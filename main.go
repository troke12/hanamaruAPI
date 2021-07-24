package main

//import "os"
//import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "time"
import "github.com/gin-gonic/gin"

type Data struct {
	Code  int     `json:"code"`
	Posts []Posts `json:"posts"`
}
type Posts struct {
	ID            string    `json:"id"`
	URL           string    `json:"url"`
	Title         string    `json:"title"`
	ContentHTML   string    `json:"content_html"`
	DatePublished time.Time `json:"date_published"`
	Category      string    `json:"category"`
	PostThumbnail string    `json:"post_thumbnail"`
	PostAuthor    string    `json:"post_author"`
	PostAuthorURL string    `json:"post_author_url"`
	Wordcount     int       `json:"wordcount"`
}

func get_content(c *gin.Context) {
    // json data
    url := "https://datenshi.pw/feed.json"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

    var data Data
    json.Unmarshal(body, &data)
    c.IndentedJSON(http.StatusOK, data)
    //fmt.Printf("Results: %v\n", data)
}

func main() {
    router := gin.Default()
    router.GET("/blogh", get_content)
    router.Run(":4343")
    //get_content()
}
