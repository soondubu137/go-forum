package mysql

import (
	"strings"

	"github.com/SoonDubu923/go-forum/model"
	"github.com/jmoiron/sqlx"
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

// GetPostList returns a list of posts.
func GetPostList(pageNum, pageSize int64) (data []*model.Post, err error) {
    data = make([]*model.Post, 0, pageSize)
    err = db.Select(&data, "SELECT post_id, title, content, author_id, community_id, created_time FROM post ORDER BY created_time DESC LIMIT ?, ?", (pageNum - 1) * pageSize, pageSize)
    if err != nil {
        zap.L().Error("GetPostList failed", zap.Error(err))
    }
    return
}

// GetPostListByID gets a list of posts by post IDs.
func GetPostListByID(ids []string) (data []*model.Post, err error) {
    data = make([]*model.Post, 0, len(ids))
    query, args, err := sqlx.In("SELECT post_id, title, content, author_id, community_id, created_time FROM post WHERE post_id IN (?) ORDER BY FIND_IN_SET(post_id, ?)", ids, strings.Join(ids, ","))
    if err != nil {
        zap.L().Error("GetPostListByID failed", zap.Error(err))
        return
    }
    query = db.Rebind(query)
    err = db.Select(&data, query, args...)
    if err != nil {
        zap.L().Error("GetPostListByID failed", zap.Error(err))
    }
    return
}