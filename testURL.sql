CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    keyword TEXT NOT NULL,
    url TEXT NOT NULL,
    content TEXT NOT NULL
);

INSERT INTO articles (keyword, url, content) VALUES
-- Golang статьи
('golang', '/articles/golang-basics', 'Introduction to Golang: Learn the basics of Go programming language, including syntax, types, and control structures.'),
('golang', '/articles/golang-concurrency', 'Understanding concurrency in Go: Explore Goroutines, Channels, and the power of Go’s concurrency model.'),
('golang', '/articles/golang-web-development', 'Web development with Go: Learn how to build web applications using the Go programming language and popular frameworks like Gin.'),
-- HTTP статьи
('http', '/articles/http-introduction', 'Introduction to HTTP: Learn the basics of HTTP, the foundational protocol for the web, including methods like GET, POST, PUT, and DELETE.'),
('http', '/articles/http-status-codes', 'HTTP Status Codes: Understand the meaning behind HTTP status codes such as 200, 404, 500, and more.'),
('http', '/articles/http-headers', 'HTTP Headers: Learn how HTTP headers work, including Content-Type, User-Agent, and Authorization.'),
-- PostgreSQL статьи
('postgresql', '/articles/postgresql-introduction', 'Introduction to PostgreSQL: Learn the basics of PostgreSQL, including installation, basic queries, and database setup.'),
('postgresql', '/articles/postgresql-performance-tuning', 'Performance Tuning in PostgreSQL: Learn how to optimize PostgreSQL for better performance by tuning queries and indexes.'),
('postgresql', '/articles/postgresql-advanced-queries', 'Advanced PostgreSQL Queries: Dive deeper into complex SQL queries, including joins, subqueries, and window functions.');

SELECT * FROM articles WHERE url = '/articles/http-status-codes';