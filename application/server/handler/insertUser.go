package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarf/sample_layer/application/server/response"
)

type insertUserRequest struct {
	Name string `json:"name"`
}

type insertUserResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// HandleGetUser はDBからユーザー情報を取得
func HandleInsertUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			response.BadRequest(w, "Invalid Request Body")
			return
		}

		var requestBody insertUserRequest
		json.Unmarshal(body, &requestBody)

		if requestBody.Name == "" {
			response.BadRequest(w, "Invalid Request Body")
			return
		}

		user, err := userService.InsertUserService(requestBody.Name)
		if err != nil {
			response.BadRequest(w, "Invalid Request Body")
			return
		}
		response.Success(w, insertUserResponse{UserID: user.UserID, Name: user.Name, CreatedAt: user.CreatedAt})
	}
}
