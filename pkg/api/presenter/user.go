package presenter

// User data
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// Client    Client    `json:"client"`
}
