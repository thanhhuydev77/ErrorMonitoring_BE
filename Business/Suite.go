package Business

import (
	"github.com/pkg/errors"
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
	"time"
)

func CreateSuite(ProjectId string, suite Models.Suite) bool {
	suite.Id = General.CreateUUID()
	result := false
	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return false
	}
	project[0].Suites = append(project[0].Suites, suite)
	result = Database.UpdateSuiteList(project[0])
	return result
}
func GetSuite(ProjectId string, Id string) (Models.Suite, bool) {

	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return Models.Suite{}, false
	}
	for _, iss := range project[0].Suites {
		if iss.Id == Id {
			return iss, true
		}
	}
	return Models.Suite{}, false
}
func FilterSuite(filter Models.SuiteFilter) ([]Models.Suite, error) {
	var listSuite []Models.Suite
	if filter.ProjectId == "" {
		return nil, errors.Errorf("ProjectID empty")
	}
	projectList, err := Database.GetProjectWithIssue(filter.ProjectId)
	if err != nil || len(projectList) == 0 {
		return nil, errors.Errorf("ProjectID invalid")
	}
	listSuite = projectList[0].Suites

	if filter.Status != "" {
		listSuite = Models.FilterSuite(listSuite, "Status", filter.Status)
	}
	if filter.Environment != "" {
		listSuite = Models.FilterSuite(listSuite, "Environment", filter.Environment)
	}
	defaultTime := time.Time{}
	if filter.FromDate != defaultTime {
		listSuite = Models.FilterSuite(listSuite, "FromDate", filter.FromDate.Format(time.RFC3339))
	}
	if filter.ToDate != defaultTime {
		listSuite = Models.FilterSuite(listSuite, "ToDate", filter.ToDate.Format(time.RFC3339))
	}

	return listSuite, nil
}
