package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mofodox.com/go-recipes-api/docs"
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

	file, err := os.ReadFile("./internal/recipes.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(file), &recipes)
	if err != nil {
		log.Println(err)
	}
}

// @title Recipes API
// @version 0.0.1
// @description This is a project on building applications using Gin

// @contact.name  Khairul Akmal
// @contact.url https://mofodox.com
// @contact.email kai@mofodox.com

// @Schemes http
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := router.Group("/api/v1")
	{
		recipeEndPoint := apiV1.Group("/recipes")
		{
			recipeEndPoint.GET("/", ListAllRecipesHandler)
			recipeEndPoint.POST("/", NewRecipeHandler)
			recipeEndPoint.PUT("/:id", UpdateRecipeHandler)
			recipeEndPoint.DELETE("/:id", DeleteRecipeHandler)
			recipeEndPoint.GET("/search", SearchRecipesHandler)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(router.Run())
}

// @BasePath /api/v1

// @Summary To create a new recipe
// @Description Creates a new recipe
// @Accepts json
// @Produce json
// @Param data body Recipe true "recipe data"
// @Success 201 {object} Recipe
// @Failure 400 {string} string
// @Router /recipes [post]
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

// @Summary Returns a list of recipes
// @Description Returns a list of recipes
// @Accepts json
// @Produce json
// @Success 201 {array} Recipe
// @Router /recipes [get]
func ListAllRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// @Summary Updates an existing recipe
// @Description Updates an existing recipe
// @Accepts json
// @Produce json
// @Param data body Recipe true "recipe data"
// @Success 200 {object} Recipe
// @Failure 400 {string} string
// @Router /recipes/:id [put]
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})

		return
	}

	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)
}

// @Summary Deletes an existing recipe
// @Description Deletes an existing recipe
// @Accepts json
// @Produce json
// @Success 200 {string} ok
// @Failure 400 {string} Error
// @Router /recipes/:id [delete]
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = 1
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})

		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted",
	})
}

// @Summary Searches for recipes by tags
// @Description Searches for recipes by tags
// @Accepts json
// @Produce json
// @Param tag query string false "recipe search by tags" Format(tag)
// @Success 200 {string} ok
// @Failure 400 {string} Error
// @Router /recipes/search [GET]
func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")

	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false

		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}

		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}
