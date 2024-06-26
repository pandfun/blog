package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/pandfun/blog/config"
	"github.com/pandfun/blog/service/auth"
	"github.com/pandfun/blog/types"
	"github.com/pandfun/blog/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// Register user routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

// Handle user registration
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", &errors))
		return
	}

	// check if user already exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to hash password: %w", err))
		return
	}

	// create user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user: %w", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}

// Handle user login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse JSON: %w", err))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", &errors))
		return
	}

	// check if user exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s not found", payload.Email))
		return
	}

	// compare password
	if !auth.ComparePassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid password"))
		return
	}

	// generate token
	jwtToken, err := auth.CreateJWT([]byte(config.Envs.JWTSecret), u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate jwt token"))
		return
	}

	// return token
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": jwtToken})
}
