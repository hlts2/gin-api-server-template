package domain

// User is tweet model
type User struct {
	ID      string `db:"tweet_id"`
	Name    string `db:"name"`
	Created string `db:"created_at"`
}

// Users represetns User slice
type Users []User
