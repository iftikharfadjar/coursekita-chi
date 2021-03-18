package main

import (
	"Coursekita-chi/controller"
	"Coursekita-chi/http"
	"Coursekita-chi/repository"
	"Coursekita-chi/service"
	"net/http"
)

var (
	repo repository.IPostRepo = repository.NewMongoDBRepository()
	postService  service.IPostService = service.NewPostService(repo)
	postController controller.IPostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	const port string = ":3000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	httpRouter.GET("/getpost", postController.GetPosts)
	httpRouter.POST("/sendpost", postController.AddPost)
	httpRouter.SERVE(port)


}
