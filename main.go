package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, err := ioutil.ReadFile("./internal/recipes.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(file), &recipes)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/recipes", NewRecipeHandler)
	router.GET("/api/v1/recipes", ListAllRecipesHandler)

	router.Run()
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()

	recipes = append(recipes, recipe)

	c.JSON(http.StatusCreated, recipe)
}

func ListAllRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}
