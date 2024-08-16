package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pandfun/blog/config"
	"github.com/pandfun/blog/service/auth"
	"github.com/pandfun/blog/types"
	"github.com/pandfun/blog/utils"
)

type Handler struct {
	store types.PostStore
}

func NewHandler(store types.PostStore) *Handler {
	return &Handler{store: store}
}

// Register post routes
func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/posts", h.handleGetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", h.handleGetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts", h.handleCreatePost).Methods(http.MethodPost)
	// router.HandleFunc("/posts/{id}", h.handleUpdatePost).Methods(http.MethodPut)
	// router.HandleFunc("/posts/{id}", h.handleDeletePost).Methods(http.MethodDelete)
}

// Return all posts
func (h *Handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := h.store.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get posts: %w", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

// Return a single post
func (h *Handler) handleGetPost(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid post ID"))
		return
	}

	post, err := h.store.GetPostByID(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get post: %w", err))
		return
	}

	if post == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("post not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

// Create a new post
func (h *Handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {

	// get token from header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing authorization token"))
		return
	}

	// validate token
	claims, err := auth.ValidateJWT(tokenString)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: %w", err))
		return
	}

	// get user ID from token
	id, err := strconv.Atoi(claims["userID"].(string))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid user ID"))
		return
	}

	// get JSON payload
	var payload types.CreatePostPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload: %w", err))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %w", errors))
		return
	}

	// set image URL
	imageURL := config.Envs.ImageURL
	if payload.ImageURL != "" {
		imageURL = payload.ImageURL
	}

	// create post
	err = h.store.CreatePost(types.Post{
		UserID:   id,
		Title:    payload.Title,
		Content:  payload.Content,
		ImageURL: imageURL,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create post: %w", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Post created"})
}

// // Update a post
// func (h *Handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {

// }

// // Delete a post
// func (h *Handler) handleDeletePost(w http.ResponseWriter, r *http.Request) {

// }
