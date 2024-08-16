package types

// User Types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=2,max=32"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Post Types

type PostStore interface {
	GetPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(post Post) error
	// UpdatePost(post Post) error
	// DeletePost(id int) error
}

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatePostPayload struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ImageURL string `json:"image_url"`
}

type UpdatePostPayload struct {
	ID       int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ImageURL string `json:"image_url"`
}
