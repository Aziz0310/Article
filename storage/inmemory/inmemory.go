package inmemory 

import "github.com/Aziz0310/bootcamp/article/models"

// InMemory ...
type InMemory struct {
	Db *DB
}

// DB mock
type DB struct {
	// InMemoryArticleData ...
	InMemoryArticleData []models.Article
	// InMemoryAuthorData ...
	InMemoryAuthorData []models.Author
}