package middleware

import (
	"net/http"
	"pzn-restful-api/helper"
	"pzn-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") == "1118020"{
		middleware.Handler.ServeHTTP(w,r)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}