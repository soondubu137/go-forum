package model

type Community struct {
    ID   int64   `json:"id" db:"community_id"`
    Name string  `json:"name" db:"name"`
}