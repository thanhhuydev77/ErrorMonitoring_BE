package Models

type User struct {
	Email        string        `json:"email" bson:"email"`
	Organization string        `json:"organization" bson:"organization"`
	Position     string        `json:"position" bson:"position"`
	Avatar       string        `json:"avatar" bson:"avatar"`
	Password     string        `json:"password" bson:"password"`
	FullName     string        `json:"fullName" bson:"fullName"`
	MainPlatform string        `json:"mainPlatform" bson:"mainPlatform"`
	ProjectList  []ProjectList `json:"projectList" bson:"projectList"`
}
type ProjectList struct {
	ProjectId string `json:"projectId" bson:"projectId"`
	Role      string `json:"role" bson:"role"`
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
