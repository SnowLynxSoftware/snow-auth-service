package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	Created      int    `json:"created"`
	LastModified int    `json:"lastModified"`
}
