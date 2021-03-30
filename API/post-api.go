package API

import (
	"Coursekita-chi/controller"
	router "Coursekita-chi/http"
	"Coursekita-chi/repository"
	"Coursekita-chi/service"
)

type postHandler struct{}

var (
	repo repository.CrudInterface             = repository.NewMongoDBRepository()
	postService  service.IPostService         = service.NewPostService(repo)
	postController controller.IPostController = controller.NewPostController(postService)
	httpRouter router.IRouter
)

func NewPostHandler(router router.IRouter) IAPI{
	httpRouter = router
	return &postHandler{}
}

func (*postHandler) Handler(){
	httpRouter.GET("/getpost", postController.GetPosts)
	httpRouter.POST("/sendpost", postController.AddPost)
	httpRouter.DELETE("/deletepost",postController.DeletePost)
	httpRouter.PUT("/updatepost",postController.UpdatePost)
	httpRouter.GET("/getpost/{id}",postController.GetPostsByID)
}


