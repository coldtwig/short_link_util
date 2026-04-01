package auth

import (
	"fmt"
	"go/http-api/configs"
	"go/http-api/pkg/req"
	"go/http-api/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		fmt.Println(payload)

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		handler.AuthService.Register(payload.Email, payload.Password, payload.Name)
	}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
