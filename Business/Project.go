package Business

import (
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
)

func CreateProject(project Models.Project) (bool, General.ErrorCode) {
	return Database.CreateProject(project)
}

func GetProjects(email string, Id string) ([]Models.Project, error) {
	return Database.GetProjects(email, Id)
}