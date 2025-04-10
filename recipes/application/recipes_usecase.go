package application

import (
	"context"
	"github.com/zono0013/recipes_api.git/recipes/domain/models"
	"github.com/zono0013/recipes_api.git/recipes/domain/repository"
)

type IRecipeUsecase interface {
	GetAllRecipes(ctx context.Context) ([]models.Recipe, error)
	GetRecipeByID(ctx context.Context, id uint) (*models.Recipe, error)
	CreateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	UpdateRecipe(ctx context.Context, id uint, recipe *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, id uint) error
}

func NewRecipesUsecase(recipeRepo repository.RecipeRepository) IRecipeUsecase {
	return &recipeUsecase{recipeRepo: recipeRepo}
}

type recipeUsecase struct {
	recipeRepo repository.RecipeRepository
}

func (r *recipeUsecase) GetAllRecipes(ctx context.Context) ([]models.Recipe, error) {
	recipes, err := r.recipeRepo.GetAllRecipes(ctx)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *recipeUsecase) GetRecipeByID(ctx context.Context, id uint) (recipe *models.Recipe, err error) {
	recipe, err = r.recipeRepo.GetRecipeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipeUsecase) CreateRecipe(ctx context.Context, rec *models.Recipe) (*models.Recipe, error) {
	recipe, err := r.recipeRepo.CreateRecipe(ctx, rec)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipeUsecase) UpdateRecipe(ctx context.Context, id uint, rec *models.Recipe) (*models.Recipe, error) {
	recipe, err := r.recipeRepo.UpdateRecipe(ctx, id, rec)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipeUsecase) DeleteRecipe(ctx context.Context, id uint) error {
	err := r.recipeRepo.DeleteRecipe(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
