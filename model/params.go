package model

type ParamRegister struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=8"`
    Reenter  string `json:"reenter" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// directions:
// 1: upvote
// 0: cancel vote
// -1: downvote
type ParamVoteData struct {
    PostID    int64  `json:"post_id,string" binding:"required"`
    Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"`
}