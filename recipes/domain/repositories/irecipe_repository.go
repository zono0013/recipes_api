package repositories

import (
	"context"

	"github.com/zono0013/recipes_api.git/recipes/domain/models"
)

type RecipeRepository interface {
	GetAllRecipes(ctx context.Context) ([]models.Recipe, error)
	GetRecipeByID(ctx context.Context, id uint) (*models.Recipe, error)
	CreateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	UpdateRecipe(ctx context.Context, id uint, recipe *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, id uint) error
}
