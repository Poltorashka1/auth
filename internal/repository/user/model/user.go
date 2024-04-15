package repoUserModel

type SignUpUser struct {
	Name     string
	Email    string
	Password string
}

type SignInUser struct {
	Email    string
	Password string
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      int
	Active    bool
	CreatedAt string
}
