package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Content string `json:"content"`
		}{ID: "16340239", Name: "WuLianghao", Content: "Hello World!"}) //配合js脚本在静态网页中加上我们想要的内容
	}
}
