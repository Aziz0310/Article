package inmemory

import (
	"errors"
	"time"

	"github.com/Aziz0310/bootcamp/article/models"
)

// AddAuthor ...
func (im InMemory) AddAuthor(id string, entity models.CreateAuthorModel) error {
	var author models.Author
	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)
	return nil
}

// GetAuthorByID ...
func (im InMemory) GetAuthorByID(id string) (models.Author, error) {
	var result models.Author
	for _, v := range im.Db.InMemoryAuthorData {
		if v.ID == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

// GetAuthorList ...
func (im InMemory) GetAuthorList() (resp []models.Author, err error) {
	resp = im.Db.InMemoryAuthorData
	return resp, err
}

// UpdateAuthor ...
func (im InMemory) UpdateAuthor(entity models.UpdateAuthorModel) error {
	var author models.Author
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == entity.ID {
			author.CreatedAt = v.CreatedAt
			author.Firstname = entity.Firstname
			author.Lastname = entity.Lastname
			t := time.Now()
			author.UpdatedAt = &t
			im.Db.InMemoryAuthorData[i] = author
		}
	}
	return nil
}

// DeleteAuthor ...
func (im InMemory) DeleteAuthor(id string) (models.Author, error) {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == id {
			im.Db.InMemoryAuthorData = rem(im.Db.InMemoryAuthorData, i)

			return v, nil
		}
	}
	return models.Author{}, errors.New("article not found")
}

func rem(slice []models.Author, s int) []models.Author {
	return append(slice[:s], slice[s+1:]...)
}
