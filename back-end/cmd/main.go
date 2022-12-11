package main

import (
	"github.com/fi9ish/filminator/pkg/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/:id", func(ctx *gin.Context) {
		movies := models.GetMovieById(ctx.Param("id"))
		ctx.JSON(200, movies)
	})

	r.GET("/getAll", func(ctx *gin.Context) {
		movies := models.GetAllMovies()
		ctx.JSON(200, gin.H{
			"movies": movies,
		})
	})

	r.POST("/getRestrictions", func(ctx *gin.Context) {

		var jsonBody [2]models.Movie
		if err := ctx.BindJSON(&jsonBody); err != nil {
			panic(err)
		}
		question := models.GetNewQuestionWithRestrictions(jsonBody)
		ctx.JSON(200, gin.H{
			"question":    question,
			"requestJson": jsonBody,
		})
	})

	r.Run()
}
