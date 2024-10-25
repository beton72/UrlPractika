package main

import (
    "database/sql"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=postgres password=xog321766 dbname=PractikaUrls sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
    initDB()

    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    r.GET("/static/", showIndex)
    r.GET("/search", searchByKeyword)
    r.GET("/articles/:article", getArticle)
    r.Run(":8080")
}

func showIndex(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}


func searchByKeyword(c *gin.Context) {
    keyword := c.Query("keyword")
    rows, err := db.Query("SELECT url FROM articles WHERE keyword = $1", keyword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
        return
    }
    defer rows.Close()

    var urls []string
    for rows.Next() {
        var url string
        if err := rows.Scan(&url); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
            return
        }
        urls = append(urls, url)
    }
    c.HTML(http.StatusOK, "result.html", gin.H{
        "Keyword": keyword,
        "URLs":    urls,
    })
}


func getArticle(c *gin.Context) {
    article := c.Param("article")

    row := db.QueryRow("SELECT content FROM articles WHERE url = $1", "/articles/"+article)

    var content string
    if err := row.Scan(&content); err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Статья не найдена"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
        }
        return
    }

    c.HTML(http.StatusOK, "article.html", gin.H{
        "URL":     article,
        "Content": content,
    })
}

