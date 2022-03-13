package logic

import (
	"encoding/json"
	"net/http"
)

type ResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int)
	CreateErrorResponse(err error) []byte
	CreateErrorStringResponse(errMessage string) []byte
}

type responseLogic struct{}

func NewResponseLogic() ResponseLogic {
	return &responseLogic{}
}

/*
	APIレスポンス送信処理
*/
func (rl *responseLogic) SendResponse(w http.ResponseWriter, response []byte, code int) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Access-Control-Expose-Headers", "X-CSRF-Token")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(code)
	w.Write(response)
}

/*
	エラーレスポンス作成
*/
func (rl *responseLogic) CreateErrorResponse(err error) []byte {
	response := map[string]interface{}{
		"error": err,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

/*
	エラーレスポンス作成（エラーメッセージはstring）
*/
func (rl *responseLogic) CreateErrorStringResponse(errMessage string) []byte {
	// response := map[string]interface{}{
	// 	"error": errMessage,
	// }
	responseBody, _ := json.Marshal(errMessage)

	return responseBody
}
