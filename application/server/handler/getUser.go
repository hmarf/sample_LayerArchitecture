package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarf/sample_layer/application/server/response"
)

type getUserRequest struct {
	UserID string `json:"userId"`
}

type getUserResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// HandleGetUser はDBからユーザー情報を取得
func HandleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			response.BadRequest(w, "Invalid Request Body")
			return
		}

		var requestBody getUserRequest
		json.Unmarshal(body, &requestBody)

		if requestBody.UserID == "" {
			response.BadRequest(w, "Invalid Request Body")
			return
		}

		user, err := userService.GetUserService(requestBody.UserID)
		response.Success(w, getUserResponse{UserID: user.UserID, Name: user.Name, CreatedAt: user.CreatedAt})
	}
}
