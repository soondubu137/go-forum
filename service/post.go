package service

import (
	"github.com/SoonDubu923/go-forum/dao/mysql"
	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/pkg/snowflake"
)

// Publish publishes a post.
func Publish(p *model.Post) (err error) {
    // generate post ID
    p.ID = snowflake.GenID()
    // insert post data into database
    return mysql.SavePost(p)
}

// GetPostDetail gets the post details by post ID.
func GetPostDetail(postID int64) (*model.PostDetail, error) {
    return mysql.GetPostByID(postID)
}

// GetPostList returns a list of posts.
func GetPostList(pageNum, pageSize int64) ([]*model.Post, error) {
    return mysql.GetPostList(pageNum, pageSize)
}
