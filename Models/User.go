package Models

type User struct {
	Email        string        `json:"email"`
	Organization string        `json:"organization"`
	Position     string        `json:"position"`
	Avatar       string        `json:"avatar"`
	Password     string        `json:"passWord"`
	FullName     string        `json:"fullName"`
	MainPlatform string        `json:"mainPlatform"`
	ProjectList  []ProjectList `json:"projectList"`
}
type ProjectList struct {
	ProjectId string `json:"projectId"`
	Role      string `json:"role"`
}

type UserRequest struct {
	Type string `json:"type"`
	User User   `json:"user"`
}
