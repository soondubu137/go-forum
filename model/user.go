package model

type User struct {
    UserID   int64  `db:"user_id"`
    Username string `db:"username"`
    Salt     string `db:"salt"`
    Password string `db:"password"`
}