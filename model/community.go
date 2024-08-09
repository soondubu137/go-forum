package model

import "time"

type Community struct {
    ID   int64   `json:"id" db:"community_id"`
    Name string  `json:"name" db:"name"`
}

type CommunityDetail struct {
    ID          int64     `json:"id" db:"community_id"`
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description,omitempty" db:"description"`
    CreatedTime time.Time `json:"created_time" db:"create_time"`
}