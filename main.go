package main

import (
	"fmt"
	"net/http"

	"github.com/Aziz0310/bootcamp/article/config"
	"github.com/Aziz0310/bootcamp/article/docs"
	_ "github.com/Aziz0310/bootcamp/article/docs"
	"github.com/Aziz0310/bootcamp/article/handlers"
	"github.com/Aziz0310/bootcamp/article/storage"
	"github.com/Aziz0310/bootcamp/article/storage/postgres"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @license.name  Apache 2.0
// @license.url   https://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	cfg := config.Load()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "2.0"

	var err error
	var stg storage.StorageI
	stg, err = postgres.InitDB(psqlConnString)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	h := handlers.Handler{
		Stg: stg,
	}

	v1 := r.Group("/v1")

	{

		v1.POST("/article", h.CreateArticle)
		v1.GET("/article", h.GetArticleList)
		v1.GET("/article/:id", h.GetArticleByID)
		v1.PUT("/article", h.UpdateArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)

		v1.POST("/author", h.CreateAuthor)
		v1.GET("/author", h.GetAuthorList)
		v1.GET("/author/:id", h.GetAuthorByID)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000") // listen
}
