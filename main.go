// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"github.com/zono0013/recipes_api.git/recipes/application"
	"github.com/zono0013/recipes_api.git/recipes/infrastructure/dao"
	"github.com/zono0013/recipes_api.git/recipes/interface/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := os.Getenv("ENV")
	if env == "staging" {
		fmt.Println("environment: staging")
	} else if env == "local" {
		fmt.Println("environment: local")
	} else if env == "production-migrating" {
		fmt.Println("environment: production-migrating")
	} else {
		fmt.Println("Error loading .env file")
	}

	// データベース接続の初期化
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	if env == "production-migrating" {
		fmt.Println("OS Exit")
		os.Exit(0)
	}

	recipeRepo := dao.NewRecipeRepository(db)

	recipeUsecase := application.NewRecipesUsecase(recipeRepo)

	recipeHandler := handler.NewRecipesHandler(recipeUsecase)

	// ルーティング
	router := gin.Default()

	// recipes := router.Group("/recipes")
	router.GET("recipes/", recipeHandler.GetAll)
	router.GET("recipes/:id", recipeHandler.GetByID)
	router.POST("recipes/", recipeHandler.Create)
	router.PATCH("recipes/:id", recipeHandler.Update)
	router.DELETE("recipes/:id", recipeHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

// initDBは別ファイルの方がいいのかな\(´ω` \)
func initDB() (*gorm.DB, error) {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// 環境変数から接続情報を取得
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// 接続文字列の構築
	//dsn := fmt.Sprintf(
	//"host=%s user=%s password=%s dbname=%s port=%s sslmode=false",
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&multiStatements=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// データベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		rollbar.Error(err)
		panic(err)
	}
	dbDriver, err := migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	if err != nil {
		rollbar.Error(err)
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", dbDriver)
	if err != nil {
		rollbar.Error(err)
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		rollbar.Error(err)
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	fmt.Println("DB migrated")

	return db, nil
}
