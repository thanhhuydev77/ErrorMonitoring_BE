package Models

import (
	"time"
)

type Project struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Platform   string     `json:"platform"`
	UserList   []UserRole `json:"userList"`
	Issues     []Issue    `json:"issues"`
	EnvList    []string   `json:"envList"`
	CreateTime time.Time  `json:"createTime"`
	CreateUser string     `json:"createUser"`
	Active     bool       `json:"active"`
}

type UserRole struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ProjectRequest struct {
	Type    string  `json:"type"`
	Project Project `json:"project"`
}
