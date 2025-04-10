package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zono0013/recipes_api.git/recipes/application"
	"github.com/zono0013/recipes_api.git/recipes/domain/models"
)

func NewRecipesHandler(recipeUsecase application.IRecipeUsecase) *recipeHandler {
	return &recipeHandler{recipeUsecase: recipeUsecase}
}

type recipeHandler struct {
	recipeUsecase application.IRecipeUsecase
}

func (h *recipeHandler) GetAll(ctx *gin.Context) {
	recipes, err := h.recipeUsecase.GetAllRecipes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]map[string]interface{}, len(recipes))
	for i, r := range recipes {
		response[i] = map[string]interface{}{
			"id":          r.ID,
			"title":       r.Title,
			"making_time": r.MakingTime,
			"serves":      r.Serves,
			"ingredients": r.Ingredients,
			"cost":        strconv.Itoa(r.Cost),
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"recipes": response,
	})
}

func (h *recipeHandler) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	recipe, err := h.recipeUsecase.GetRecipeByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No Recipe found"})
		return
	}

	res := map[string]interface{}{
		"id":          recipe.ID,
		"title":       recipe.Title,
		"making_time": recipe.MakingTime,
		"serves":      recipe.Serves,
		"ingredients": recipe.Ingredients,
		"cost":        strconv.Itoa(recipe.Cost),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Recipe details by id",
		"recipe":  []map[string]interface{}{res},
	})
}

func (h *recipeHandler) Create(ctx *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		MakingTime  string `json:"making_time" binding:"required"`
		Serves      string `json:"serves" binding:"required"`
		Ingredients string `json:"ingredients" binding:"required"`
		Cost        int    `json:"cost" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Recipe creation failed!",
			"required": "title, making_time, serves, ingredients, cost",
		})
		return
	}

	recipe := &models.Recipe{
		Title:       req.Title,
		MakingTime:  req.MakingTime,
		Serves:      req.Serves,
		Ingredients: req.Ingredients,
		Cost:        req.Cost,
	}

	res, err := h.recipeUsecase.CreateRecipe(ctx, recipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Recipe creation failed!",
			"required": "title, making_time, serves, ingredients, cost",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Recipe successfully created!",
		"recipe": []map[string]interface{}{
			{
				"id":          res.ID,
				"title":       res.Title,
				"making_time": res.MakingTime,
				"serves":      res.Serves,
				"ingredients": res.Ingredients,
				"cost":        strconv.Itoa(res.Cost),
				"created_at":  res.CreatedAt.Format("2006-01-02 15:04:05"),
				"updated_at":  res.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	})
}

func (h *recipeHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		MakingTime  string `json:"making_time"`
		Serves      string `json:"serves"`
		Ingredients string `json:"ingredients"`
		Cost        int    `json:"cost"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to decode request body"})
		return
	}

	recipe := &models.Recipe{
		Title:       req.Title,
		MakingTime:  req.MakingTime,
		Serves:      req.Serves,
		Ingredients: req.Ingredients,
		Cost:        req.Cost,
	}

	updated, err := h.recipeUsecase.UpdateRecipe(ctx, uint(id), recipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Recipe successfully updated!",
		"recipe": []map[string]interface{}{
			{
				"title":       updated.Title,
				"making_time": updated.MakingTime,
				"serves":      updated.Serves,
				"ingredients": updated.Ingredients,
				"cost":        strconv.Itoa(updated.Cost),
			},
		},
	})
}

func (h *recipeHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := h.recipeUsecase.DeleteRecipe(ctx, uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No Recipe found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Recipe successfully removed!"})
}
