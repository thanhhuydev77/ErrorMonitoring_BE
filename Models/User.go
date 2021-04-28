package Models

type User struct {
	Email        string        `json:"email"`
	Organization string        `json:"organization"`
	Position     string        `json:"position"`
	Avatar       string        `json:"avatar"`
	Password     string        `json:"password"`
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

func RemoveProjectRole(slice []ProjectList, value ProjectList) []ProjectList {
	index := findIndexUser(slice, value)
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func ModifyProjectRole(slice []ProjectList, userRole ProjectList) []ProjectList {
	for i := range slice {
		if slice[i].ProjectId == userRole.ProjectId {
			slice[i].Role = userRole.Role
		}
	}
	return slice
}
func findIndexUser(slice []ProjectList, role ProjectList) int {
	for i, userRole2 := range slice {
		if userRole2.ProjectId == role.ProjectId {
			return i
		}
	}
	return -1
}
