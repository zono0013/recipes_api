package dao

import (
	"context"
	"github.com/zono0013/recipes_api.git/recipes/domain/models"
	"gorm.io/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) *recipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) GetAllRecipes(ctx context.Context) ([]models.Recipe, error) {
	var recipes []models.Recipe
	if err := r.db.WithContext(ctx).Order("created_at desc").Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *recipeRepository) GetRecipeByID(ctx context.Context, id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := r.db.WithContext(ctx).First(&recipe, id).Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (r *recipeRepository) CreateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	if err := r.db.WithContext(ctx).Create(&recipe).Error; err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipeRepository) UpdateRecipe(ctx context.Context, id uint, recipe *models.Recipe) (*models.Recipe, error) {
	if err := r.db.WithContext(ctx).Model(&models.Recipe{}).Where("id = ?", id).Updates(recipe).Error; err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipeRepository) DeleteRecipe(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Recipe{}, id).Error; err != nil {
		return err
	}
	return nil
}
