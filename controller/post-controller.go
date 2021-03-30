package controller

import (
	"Coursekita-chi/models"
	"Coursekita-chi/service"
	chitools "Coursekita-chi/http"
	"encoding/json"
	"net/http"
)

var (
	postService service.IPostService
	tools chitools.IChiTools = chitools.NewChiTools()
)

type controller struct {}

type IPostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
	DeletePost(resp http.ResponseWriter, req *http.Request)
	UpdatePost(resp http.ResponseWriter, req *http.Request)
	GetPostsByID(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.IPostService) IPostController{
	postService = service
	return &controller{}
}

//function for GET post data
func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	posts , err := postService.FindAll()

	//find all in mongo db
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

//function for GET post data
func (*controller) GetPostsByID(resp http.ResponseWriter, req *http.Request) {

	postID := tools.GetParams(req, "id")

	post , err := postService.FindByID(postID)

	//find all in mongo db
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error get data from db"}`))
		return
	}

	//set header
	resp.Header().Set("Content-type", "application/json")
	//marshal json
	result, err := json.Marshal(post)
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

	err1 := postService.Create(&post)

	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error save to db"}`))
		return
	}

	//status ok
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message" : "save to db success"}`))
}

func (*controller) DeletePost(resp http.ResponseWriter, req *http.Request) {
	//set header
	resp.Header().Set("Content-type", "application/json")
	var postsID models.IDCollection

	//unmarshal json
	err := json.NewDecoder(req.Body).Decode(&postsID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error unmarshal the request"}`))
		return
	}

	err = postService.Remove(postsID)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error remove data in db"}`))
		return
	}

	//p, _ := json.Marshal(postsID)

	//status ok
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message" : "remove data in db success"}`))
	//resp.Write(p)
}


//function for POST post data
func (*controller) UpdatePost(resp http.ResponseWriter, req *http.Request) {
	//set header
	resp.Header().Set("Content-type", "application/json")
	var post models.Post

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
		resp.Write([]byte(`{"error" : "error validate the request"}`))
		return
	}

	err1 := postService.Update(&post)

	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error update db"}`))
		return
	}

	//status ok
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message" : "update success"}`))
}