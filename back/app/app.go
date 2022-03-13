package app

import (
	"github.com/joho/godotenv"
	"github.com/obrkn/twitter/controllers"
	"github.com/obrkn/twitter/db"
	"github.com/obrkn/twitter/repositories"
	"github.com/obrkn/twitter/router"
	"github.com/obrkn/twitter/services"
	"github.com/obrkn/twitter/utils/logic"
	"github.com/obrkn/twitter/utils/validation"
)

func App() {
	// .envファイル読み込み
	godotenv.Load()

	// DB接続
	db := db.Init()
	defer db.Close()

	// logic層
	responseLogic := logic.NewResponseLogic()

	// validation層
	authValidate := validation.NewAuthValidation()
	tweetValidate := validation.NewTweetValidation()

	// repository層
	userRepo := repositories.NewUserRepository(db)
	tweetRepo := repositories.NewTweetRepository(db)
	// service層
	authService := services.NewAuthService(userRepo, responseLogic, authValidate)
	tweetService := services.NewTweetService(tweetRepo, responseLogic, tweetValidate)
	// controller層
	authController := controllers.NewAuthController(authService)
	tweetController := controllers.NewTweetController(tweetService, authService)

	// router設定
	authRouter := router.NewAuthRouter(authController)
	tweetRouter := router.NewTweetRouter(tweetController)
	mainRouter := router.NewMainRouter(authRouter, tweetRouter)

	// 起動
	mainRouter.StartWebServer()
}
