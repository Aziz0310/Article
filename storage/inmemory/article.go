package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/Aziz0310/bootcamp/article/models"
)

// AddArticle ...
func (im InMemory) AddArticle(id string, entity models.CreateArticleModel) error {
	var article models.Article
	article.ID = id
	article.Content = entity.Content

	author, err := im.GetAuthorByID(entity.AuthorID)
	if err != nil {
		return err
	}

	article.AuthorID = author.ID
	article.CreatedAt = time.Now()

	im.Db.InMemoryArticleData = append(im.Db.InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func (im InMemory) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	for _, v := range im.Db.InMemoryArticleData {
		if v.ID == id && v.DeletedAt == nil {
			author, err := im.GetAuthorByID(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			result.DeletedAt = v.DeletedAt
			return result, nil

		}
	}
	return result, errors.New("article not found")
}

// GetArticleList ...
func (im InMemory) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	off := 0
	c := 0
	for _, v := range im.Db.InMemoryArticleData {
		if v.DeletedAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {
			if offset <= off {
				c++
				resp = append(resp, v)
			}

			if c >= limit {
				break
			}

			off++
		}
	}
	return resp, err
}

// UpdateArticle ...
func (im InMemory) UpdateArticle(entity models.UpdateArticleModel) error {
	for i, v := range im.Db.InMemoryArticleData {
		if v.ID == entity.ID && v.DeletedAt == nil {
			v.Content = entity.Content
			t := time.Now()
			v.UpdatedAt = &t
			im.Db.InMemoryArticleData[i] = v

			return nil
		}
	}
	return errors.New("article not found")
}

// DeleteArticle ...
func (im InMemory) DeleteArticle(id string) error {
	for i, v := range im.Db.InMemoryArticleData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			im.Db.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}
