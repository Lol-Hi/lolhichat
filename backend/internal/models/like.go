package models

type CommentsLike struct {
	UserID			int 		`json:"user_id"`
	CommentID		int 		`json:"comment_id"`
}
