package domain

// Tweet is tweet model
type Tweet struct {
	id      string `db:"tweet_id"`
	userID  string `db:"user_id"`
	text    string `db:"text"`
	query   string `db:"query"`
	created string `db:"created_at"`
}

// Tweets represetns Tweet slice
type Tweets []Tweet
