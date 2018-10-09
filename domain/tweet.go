package domain

// Tweet is tweet model
type Tweet struct {
	ID      string `db:"tweet_id"`
	UserID  string `db:"user_id"`
	Text    string `db:"text"`
	Query   string `db:"query"`
	Created string `db:"created_at"`
}

// Tweets represetns Tweet slice
type Tweets []Tweet
