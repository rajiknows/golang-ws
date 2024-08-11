package utils

import "github.com/rajiknows/vedashala/internal/database"
import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func DatabaseUserToUser(db database.User) User {
	return User{
		ID:        db.ID,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		Name:      db.Name,
	}
}
