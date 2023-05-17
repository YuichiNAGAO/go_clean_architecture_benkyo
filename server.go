package main

import (
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/controller"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/infrastructure"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/initializer"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/router"
)

func init() {
	// .env fileをloadし、環境変数として扱えるようにする
	initializer.LoadEnvVariables()
	infrastructure.ConnectDB()
}

var (
	httpRouter     = router.NewMuxRouter()
	postConrtoller = controller.NewPostController()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/posts", postConrtoller.GetPosts)
	httpRouter.POST("/posts", postConrtoller.AddPost)
	// curl -v -X POST -H "Content-Type: application/json" -d '{"id":1, "title":"タイトル", "text":"本文"}'  localhost:8080/posts

	httpRouter.SERVE(port)
}
