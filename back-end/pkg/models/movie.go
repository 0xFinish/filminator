package models

import (
	"fmt"
	"strconv"

	"math/rand"

	"github.com/fi9ish/filminator/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	Show_id      int    `gorm:"primaryKey;column:show_id" json:"show_id"`
	Show_type    string `gorm:"column:show_type" json:"show_type"`
	Title        string `gorm:"column:title" json:"title"`
	Director     string `gorm:"column:director" json:"director"`
	Casting      string `gorm:"column:casting" json:"casting"`
	Country      string `gorm:"column:country" json:"country"`
	Date_added   string `gorm:"column:date_added" json:"date_added"`
	Release_year int    `gorm:"column:release_year" json:"release_year"`
	Rating       string `gorm:"column:rating" json:"rating"`
	Duration     string `gorm:"column:duration" json:"duration"`
	Listed_in    string `gorm:"column:listed_in" json:"listed_in"`
}

type Question struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
	RawName  string `json:"raw_name,omitempty"`
}

func (Movie) TableName() string {
	return "filminator"
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func GetMovieById(id string) Movie {
	var Movies Movie
	db.First(&Movies, "show_id = ?", id)
	fmt.Println("LEN MOVIES")
	fmt.Println(Movies)
	return Movies
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Limit(10).Find(&Movies)
	return Movies
}

func GetNewQuestionWithRestrictions(movieRestrictions [2]Movie) Question {
	var Movies []Movie
	var count int
	db.Not(movieRestrictions[1]).Where(movieRestrictions[0]).Find(&Movies).Count(&count)
	randMovie := Movies[rand.Intn(len(Movies))]
	var question Question

	switch randNumber := rand.Intn(9) + 1; randNumber {

	case 1:
		question.Question = fmt.Sprintf("Your show type is %s?", randMovie.Show_type)
		question.Answer = randMovie.Show_type
		question.RawName = "show_type"
	case 2:
		question.Question = fmt.Sprintf("Your show title is %s", randMovie.Title)
		question.Answer = randMovie.Title
		question.RawName = "title"
	case 3:
		question.Question = fmt.Sprintf("%s directed your show?", randMovie.Director)
		question.Answer = randMovie.Director
		question.RawName = "director"
	case 4:
		question.Question = fmt.Sprintf("Have this people: `%s` played in your movie?", randMovie.Casting)
		question.Answer = randMovie.Casting
		question.RawName = "casting"
	case 5:
		question.Question = fmt.Sprintf("Was your movie filmed in %s", randMovie.Country)
		question.Answer = randMovie.Country
		question.RawName = "country"
	case 6:
		question.Question = fmt.Sprintf("Was your film added at %s", randMovie.Date_added)
		question.Answer = randMovie.Date_added
		question.RawName = "date_added"
	case 7:
		question.Question = fmt.Sprintf("Your movie was released in %d", randMovie.Release_year)
		question.Answer = strconv.Itoa(randMovie.Release_year)
		question.RawName = "release_year"
	case 8:
		question.Question = fmt.Sprintf("Is your movie rated %s ?", randMovie.Rating)
		question.Answer = randMovie.Rating
		question.RawName = "rating"
	case 9:
		question.Question = fmt.Sprintf("Is your movie duration %s", randMovie.Duration)
		question.Answer = randMovie.Duration
		question.RawName = "duration"
	case 10:
		question.Question = fmt.Sprintf("Is your movie listed in %s", randMovie.Listed_in)
		question.Answer = randMovie.Listed_in
		question.RawName = "listed_in"
	}
	if count == 1 {
		question.Question = fmt.Sprintf("You movie name is %s", randMovie.Title)
		question.Answer = fmt.Sprintf(randMovie.Title)
		question.Answer = "end"
		return question
	}
	fmt.Printf("Amount of rows counted: %d", count)
	return question
}
