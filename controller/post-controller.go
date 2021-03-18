package controller

import (
	"Coursekita-chi/models"
	"Coursekita-chi/service"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	postService service.IPostService
)

type controller struct {}

type IPostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.IPostService) IPostController{
	postService = service
	return &controller{}
}

//function for GET post data
func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {

	//find all in mongo db
	posts , err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error get data from db"}`))
		return
	}

	//set header
	resp.Header().Set("Content-type", "application/json")
	//marshal json
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error marshalling the posts array"}`))
		return
	}

	//send status
	resp.WriteHeader(http.StatusOK)
	//send result
	resp.Write(result)
}

//function for POST post data
func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	//set header
	resp.Header().Set("Content-type", "application/json")
	var post models.Post
	fmt.Println(req.Body)
	//unmarshal json
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error unmarshal the request"}`))
		return
	}

	err = postService.Validate(&post)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error unmarshal the request"}`))
		return
	}

	result , err1 := postService.Create(&post)

	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error save to db"}`))
		return
	}

	//status ok
	resp.WriteHeader(http.StatusOK)
	//marshal json
	json.NewEncoder(resp).Encode(result)
}
