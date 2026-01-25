// Package models contains the structs that correspond to the rows of the database.
package models

// CommentsLike is the format of an entry in the "comments_likes" table.
type CommentsLike struct {
	UserID    int `json:"user_id"`    // The ID of the user who liked the comment.
	CommentID int `json:"comment_id"` // The ID of the comment that the user liked.
}
