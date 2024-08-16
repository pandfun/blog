package post

import (
	"database/sql"

	"github.com/pandfun/blog/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoPost(row *sql.Rows) (*types.Post, error) {
	p := new(types.Post)
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.Title,
		&p.Content,
		&p.ImageURL,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Store) GetPosts() ([]types.Post, error) {

	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := make([]types.Post, 0)

	for rows.Next() {
		post, err := scanRowIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

func (s *Store) GetPostByID(id int) (*types.Post, error) {

	rows, err := s.db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Post)
	for rows.Next() {
		p, err = scanRowIntoPost(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, nil
	}

	return p, nil
}

func (s *Store) CreatePost(post types.Post) error {

	_, err := s.db.Exec("INSERT INTO posts (user_id, title, content, image_url) VALUES (?, ?, ?, ?)",
		post.UserID, post.Title, post.Content, post.ImageURL)

	if err != nil {
		return err
	}

	return nil
}
