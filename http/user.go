package main

type User struct {
	// username string `form:"username" json:"username" xml:"username"  binding:"required"`
	// age      int64  `form:"age" json:"age" xml:"age" binding:"required"`
	UserName string `json:"username"`
	Age      int    `json:"age"`
}
