package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/JensonCode/go-restful-api/internal/database"
	"github.com/JensonCode/go-restful-api/internal/response"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "asdasd"

func (auth *authHandler) Login(w http.ResponseWriter, r *http.Request){
	
	auth.mu.Lock()
	defer auth.mu.Unlock()

	loginUser, err := getRequestBody(w, r)
	if err != nil {
		return
	}

	var count int
	if err := database.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE user=?", loginUser.User).Scan(&count); err != nil{
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	if count == 0 {
		response.ResponseWithError(w, http.StatusBadRequest, "User name not found")
		return 
	}

	query := "SELECT id, user, password FROM Users WHERE user=?"
	row := database.DB.QueryRow(query, loginUser.User)

	var user User
	if err:= row.Scan(&user.ID, &user.User, &user.Password); err != nil{
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, loginUser.Password); err != nil{
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}


	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer: strconv.Itoa(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour * 1).Unix(), 0)),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.ResponseWithJSON(w, http.StatusOK, token)

}

func (auth *authHandler) Register( w http.ResponseWriter, r *http.Request){

	auth.mu.Lock()
	defer auth.mu.Unlock()

	registerUser, err := getRequestBody(w, r)
	if err != nil {
		return
	}

	var count int
	if err := database.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE user=?", registerUser.User).Scan(&count); err != nil{
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	if count > 0 {
		response.ResponseWithError(w, http.StatusBadRequest, "User name has already been used")
		return 
	}
	
	password, _ := bcrypt.GenerateFromPassword(registerUser.Password, 10)
	registerUser.Password = password

	query := "INSERT INTO Users (user, password) VALUES (?, ?)"

	if 	_, err:= database.DB.Exec(query,registerUser.User, password); err != nil {
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
        return 
    }

	response.ResponseWithJSON(w, http.StatusCreated, registerUser)

}

func getRequestBody(w http.ResponseWriter,r *http.Request, ) (User, error) {
	var user User

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return user, err
	}

	if r.Header.Get("content-type") != "application/json"{
		response.ResponseWithError(w, http.StatusUnsupportedMediaType, "content type must be application/json" )
		return user, err
	}

	if err := json.Unmarshal(body, &user); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error() )
		return user, err
	}

	return user, nil
}

