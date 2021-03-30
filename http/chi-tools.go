package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

type chiTools struct{}

func NewChiTools() IChiTools{
	return &chiTools{}
}

func (*chiTools) GetParams(r *http.Request, key string) string{
	return chi.URLParam(r,key)
}