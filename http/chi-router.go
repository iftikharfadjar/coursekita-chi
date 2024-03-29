package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type chiRouter struct {}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() IRouter {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Get(uri,f)
}
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Delete(uri, f)
}

func (*chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Put(uri, f)
}

func (*chiRouter) SERVE(port string){
	fmt.Printf("chi HTTP Server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher )
}
