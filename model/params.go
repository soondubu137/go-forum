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