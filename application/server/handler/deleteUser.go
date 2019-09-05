package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarf/sample_layer/application/server/response"
)

type deleteUserRequest struct {
	UserID string `json:"userId"`
}

type deleteUserResponse struct {
	Result string `json:"results"`
}

// HandleGetUser はDBからユーザー情報を取得
func HandleDeleteUser() http.HandlerFunc {
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

		err = userService.DeleteUserService(requestBody.UserID)
		if err != nil {
			response.Success(w, deleteUserResponse{Result: "failure"})
		}
		response.Success(w, deleteUserResponse{Result: "success"})
	}
}
