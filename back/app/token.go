package app

import (
	"fmt"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w = SetHeader(w)
	cookie := &http.Cookie{
		Name:  "_c_session", // ここにcookieの名前を記述
		Value: "bar",        // ここにcookieの値を記述
	}
	http.SetCookie(w, cookie)

	fmt.Println(cookie)
}
