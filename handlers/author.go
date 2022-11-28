package handlers

import (
	"net/http"
	"strconv"

	"github.com/Aziz0310/bootcamp/article/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.Stg.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.Stg.GetAuthorByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | GetList",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     update author by id
// @Description an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/author/{id} [get]
func (h Handler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	author, err := h.Stg.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResponse{data=[]models.Author}
// @Router      /v1/author [get]
func (h Handler) GetAuthorList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	authorList, err := h.Stg.GetAuthorList(offset,limit,searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    authorList,
	})
}

// UpdateAuthor godoc
// @Summary     Update Author
// @Description update Authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       Author body     models.UpdateAuthorModel true "Author body"
// @Success     200    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Failure     404    {object} models.JSONErrorResponse
// @Router      /v1/author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var data models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateAuthor(data)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleByID(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: article,
	})

}

// DeleteAuthor godoc
// @Summary     Delete author by id
// @Description delete authors by id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "authors id"
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.Stg.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResponse{
		Data: author,
	})
}
