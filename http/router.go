package router

import (
	"net/http"
)

type IRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type IChiTools interface {
	GetParams(r *http.Request, key string) string
}
