package main

import (
	"Coursekita-chi/API"
	router "Coursekita-chi/http"
)

var (
	chiRouter   router.IRouter = router.NewChiRouter()
	postHandler API.IAPI       = API.NewPostHandler(chiRouter)
)

func main() {
	const port string = ":3000"

	postHandler.Handler()

	chiRouter.SERVE(port)


}
