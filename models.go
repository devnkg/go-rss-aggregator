package main

// Feed represents an RSS feed
type Feed struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	UserID    int    `json:"user_id"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// User represents a user
type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	APIKey    string `gorm:"uniqueIndex" json:"api_key"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// Article represents an article from an RSS feed
type Article struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PublishedAt int64  `json:"published_at"`
	FeedID      int    `json:"feed_id"`
	Feed        Feed   `gorm:"foreignKey:FeedID" json:"-"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

// TableName specifies the table name for Feed
func (Feed) TableName() string {
	return "feeds"
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}

// TableName specifies the table name for Article
func (Article) TableName() string {
	return "articles"
}
