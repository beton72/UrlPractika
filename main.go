package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Article struct {
    Keyword string
    URL     string
    Content string
}

var articles = []Article{
    // Golang статьи
    {"golang", "/articles/golang-basics", "Introduction to Golang: Learn the basics of Go programming language, including syntax, types, and control structures."},
    {"golang", "/articles/golang-concurrency", "Understanding concurrency in Go: Explore Goroutines, Channels, and the power of Go’s concurrency model."},
    {"golang", "/articles/golang-web-development", "Web development with Go: Learn how to build web applications using the Go programming language and popular frameworks like Gin."},
    // HTTP статьи
    {"http", "/articles/http-introduction", "Introduction to HTTP: Learn the basics of HTTP, the foundational protocol for the web, including methods like GET, POST, PUT, and DELETE."},
    {"http", "/articles/http-status-codes", "HTTP Status Codes: Understand the meaning behind HTTP status codes such as 200, 404, 500, and more."},
    {"http", "/articles/http-headers", "HTTP Headers: Learn how HTTP headers work, including Content-Type, User-Agent, and Authorization."},
    // PostgreSQL статьи
    {"postgresql", "/articles/postgresql-introduction", "Introduction to PostgreSQL: Learn the basics of PostgreSQL, including installation, basic queries, and database setup."},
    {"postgresql", "/articles/postgresql-performance-tuning", "Performance Tuning in PostgreSQL: Learn how to optimize PostgreSQL for better performance by tuning queries and indexes."},
    {"postgresql", "/articles/postgresql-advanced-queries", "Advanced PostgreSQL Queries: Dive deeper into complex SQL queries, including joins, subqueries, and window functions."},
}

func main() {
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

    var urls []string
    for _, article := range articles {
        if article.Keyword == keyword {
            urls = append(urls, article.URL)
        }
    }

    c.HTML(http.StatusOK, "result.html", gin.H{
        "Keyword": keyword,
        "URLs":    urls,
    })
}

func getArticle(c *gin.Context) {
    articleURL := "/articles/" + c.Param("article")

    var content string
    for _, article := range articles {
        if article.URL == articleURL {
            content = article.Content
            break
        }
    }

    if content == "" {
        c.JSON(http.StatusNotFound, gin.H{"error": "Статья не найдена"})
        return
    }

    c.HTML(http.StatusOK, "article.html", gin.H{
        "URL":     articleURL,
        "Content": content,
    })
}
