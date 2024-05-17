package repoUserModel

import "time"

type SignUpUser struct {
	Name     string
	Email    string
	Password string
}

type SignInUser struct {
	Email    string
	Password string
}

// User - user from data base
type User struct {
	// from table Users
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	// from table UserRoles
	Roles string `db:"roles"`
}

type Users struct {
	Users []User
}
