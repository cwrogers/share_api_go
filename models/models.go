package models

type Post struct {
	Id        int    `json:"id"`
	Data      string `json:"content"`
	Author    int    `json:"author"`
	CreatedAt int64  `json:"created_at"`
	Type      string `json:"type"`
}

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ShareRequest struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
	Time int64  `json:"time"`
}
