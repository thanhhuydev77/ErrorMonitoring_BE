package Models

import (
	"time"
)

type Project struct {
	Id           string     `json:"id" bson:"id"`
	Name         string     `json:"name" bson:"name"`
	Platform     string     `json:"platform" bson:"platform"`
	UserList     []UserRole `json:"userList" bson:"user_list"`
	Issues       []Issue    `json:"issues" bson:"issues"`
	EnvList      []string   `json:"envList" bson:"envList"`
	CreateTime   time.Time  `json:"createTime" bson:"createTime"`
	CreateUser   string     `json:"createUser" bson:"createUser"`
	Active       bool       `json:"active" bson:"active"`
	Suites       []Suite    `json:"suites" bson:"suites"`
	EnableTrello bool       `json:"enableTrello" bson:"enableTrello"`
	TrelloInfo   TrelloInfo `json:"trelloInfo" bson:"trelloInfo"`
	EnableSlack  bool       `json:"enableSlack" bson:"enableSlack"`
	SlackInfo    SlackInfo  `json:"slackInfo" bson:"slackInfo"`
}

type UserRole struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ProjectRequest struct {
	Type    string  `json:"type"`
	Project Project `json:"project"`
}

func Contain(listUserRole []UserRole, userRole UserRole) bool {
	for _, userRole2 := range listUserRole {
		if userRole2.Email == userRole.Email {
			return true
		}
	}
	return false
}
func RemoveUserRole(slice []UserRole, value UserRole) []UserRole {
	index := findIndex(slice, value)
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func ModifyUserRole(slice []UserRole, userRole UserRole) []UserRole {
	for _, userRole2 := range slice {
		if userRole2.Email == userRole.Email {
			userRole2.Role = userRole.Role
		}
	}
	return slice
}
func findIndex(slice []UserRole, role UserRole) int {
	for i, userRole2 := range slice {
		if userRole2.Email == role.Email {
			return i
		}
	}
	return -1
}
