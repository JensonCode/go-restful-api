package auth

import (
	"net/http"
	"strings"
	"sync"

	"github.com/JensonCode/go-restful-api/internal/response"
)

type User struct {
	ID       int `json:"id"`
	User     string `json:"user"`
	Password []byte `json:"-"`
}

type authHandler struct {
	mu sync.Mutex
}

func AuthRoutes(){
    http.Handle("/api/", new(authHandler))
}

func (auth *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	//url: baseURL/api/login POST{username[username], [password]=[password]} (check password)
	//url: baseURL/api/register POST{username=[username], [password]=[password], [password2]=[password2]} (post one)
	pathSlice := strings.Split(r.URL.String(), "/")

	if r.Method != "POST"{
		response.ResponseWithError(w, http.StatusBadRequest, "Invalid http method")
		return
	}

	switch pathSlice[2]{
		case "login":
			auth.Login(w, r)
		case "register":
			auth.Register(w, r)
		default:
			response.ResponseWithError(w, http.StatusBadRequest, "Invalid http method")
	}
	

}

