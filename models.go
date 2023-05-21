package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerBedir/go-learning/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string	`json:"name"`
	Url       string	`json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}
type Feed_Follow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}
func databaseUserToUser(dbUser database.User) User {
	return User{
        ID:        dbUser.ID,
        CreatedAt: dbUser.CreatedAt,
        UpdatedAt: dbUser.UpdatedAt,
        Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
    }
}
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
        ID:        dbFeed.ID,
        CreatedAt: dbFeed.CreatedAt,
        UpdatedAt: dbFeed.UpdatedAt,
        Name:      dbFeed.Name,
		Url:       dbFeed.Url,
        UserID:    dbFeed.UserID,
    }
}
func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeed))
    for i, dbFeed := range dbFeed {
        feeds[i] = databaseFeedToFeed(dbFeed)
    }
    return feeds
}
func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) Feed_Follow {
	return Feed_Follow{
        ID:        dbFeedFollow.ID,
        CreatedAt: dbFeedFollow.CreatedAt,
        UpdatedAt: dbFeedFollow.UpdatedAt,
        UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
    }
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []Feed_Follow {
	feeds := make([]Feed_Follow, len(dbFeedFollows))
    for i, dbFeed := range dbFeedFollows {
        feeds[i] = databaseFeedFollowToFeedFollow(dbFeed)
    }
    return feeds
}



type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}