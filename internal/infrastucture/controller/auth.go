package controller

import (
	"Kode_test/internal/domain/auth/model"
	"Kode_test/pkg/api/response"
	"Kode_test/pkg/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type AuthUseCase interface {
	SignUP(userDto model.SignUpRequestDTO) error
	Login(email, password string) (string, error)
}

type AuthController struct {
	authUseCase AuthUseCase
	log         *slog.Logger
}

func NewAuthController(handler *chi.Mux, userUseCase AuthUseCase, log *slog.Logger) {
	controller := &AuthController{authUseCase: userUseCase, log: log}

	handler.Post("/auth/signup", controller.signUP)
	handler.Post("/auth/login", controller.login)
}

func (c *AuthController) signUP(w http.ResponseWriter, r *http.Request) {
	var userDto model.SignUpRequestDTO

	err := render.DecodeJSON(r.Body, &userDto)

	if err != nil {
		c.log.Error("could not decode body", sl.Err(err))
		render.JSON(w, r, response.Error("could not decode body"))
		return
	}

	if err = c.authUseCase.SignUP(userDto); err != nil {
		c.log.Error("failed to create auth", sl.Err(err))
		render.JSON(w, r, response.Error("failed to create auth"))
		return
	}
	render.JSON(w, r, response.OK())
}

func (c *AuthController) login(w http.ResponseWriter, r *http.Request) {
	var loginDTO model.LoginRequestDTO

	err := render.DecodeJSON(r.Body, &loginDTO)

	if err != nil {
		c.log.Error("could not decode body", sl.Err(err))
		render.JSON(w, r, response.Error("could not decode body"))
		return
	}

	var token string
	if token, err = c.authUseCase.Login(loginDTO.Email, loginDTO.Password); err != nil {
		c.log.Error("could not login", sl.Err(err))
		render.JSON(w, r, response.Error("could not login"))
		return
	}
	render.JSON(w, r, response.OkWithData(token))
}
