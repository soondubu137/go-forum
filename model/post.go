package model

import "time"

type Post struct {
    ID          int64     `json:"id,string" db:"post_id"`
    Votes       int64     `json:"votes" db:"votes"`
    AuthorID    int64     `json:"author_id,string" db:"author_id"`
    CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
    Status      int32     `json:"status" db:"status"`
    Title       string    `json:"title" db:"title" binding:"required"`
    Content     string    `json:"content" db:"content" binding:"required"`
    CreatedTime time.Time `json:"created_time" db:"created_time"`
}

type PostDetail struct {
    Author string    `json:"author_name"`
    *Post
    *CommunityDetail `json:"community"`
}