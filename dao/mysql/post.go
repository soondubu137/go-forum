package mysql

import (
	"github.com/SoonDubu923/go-forum/model"
	"go.uber.org/zap"
)

// SavePost saves a post to the database.
func SavePost(p *model.Post) (err error) {
    _, err = db.NamedExec("INSERT INTO post (post_id, title, content, author_id, community_id) VALUES (:post_id, :title, :content, :author_id, :community_id)", p)
    return
}

// GetPostByID gets a post by post ID.
func GetPostByID(postID int64) (data *model.PostDetail, err error) {
    // first get the post, details of which are to be fetched
    p := new(model.Post)
    err = db.Get(p, "SELECT post_id, title, content, author_id, community_id, created_time FROM post WHERE post_id = ?", postID)
    if err != nil {
        zap.L().Error("GetPostByID failed", zap.Int64("postID", postID), zap.Error(err))
        return
    }

    data = new(model.PostDetail)
    // populate embedded fields
    data.Post = p
    data.CommunityDetail, err = GetCommunityDetailByID(p.CommunityID)
    if err != nil {
        zap.L().Error("GetPostByID failed", zap.Int64("postID", postID), zap.Error(err))
        return
    }
    data.Author, err = GetUsernameByID(p.AuthorID)
    if err != nil {
        zap.L().Error("GetPostByID failed", zap.Int64("postID", postID), zap.Error(err))
        return
    }
    
    return
}