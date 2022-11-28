package storage

import "github.com/Aziz0310/bootcamp/article/models"

// StorageI ...
type StorageI interface {
	AddArticle(id string, entity models.CreateArticleModel) error
	GetArticleByID(id string) (models.PackedArticleModel, error)
	GetArticleList(offset, limit int, search string) (resp []models.Article, err error)
	UpdateArticle(entity models.UpdateArticleModel) error
	DeleteArticle(id string) error

	AddAuthor(id string, entity models.CreateAuthorModel) error
	GetAuthorByID(id string) (models.GetAuthorByIdResp, error)
	GetAuthorList(offset, limit int, search string) (resp []models.Author, err error)
	UpdateAuthor(entity models.UpdateAuthorModel) error
	DeleteAuthor(id string) error
}
