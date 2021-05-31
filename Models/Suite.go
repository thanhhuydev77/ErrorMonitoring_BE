package Models

import "time"

type Suite struct {
	Id             string    `json:"id"`
	Environment    string    `json:"environment"`
	StartedTestsAt time.Time `json:"startedTestsAt"`
	EndedTestsAtn  time.Time `json:"endedTestsAtn"`
	TotalDuration  int       `json:"totalDuration"`
	TotalSuites    int       `json:"totalSuites"`
	TotalTests     int       `json:"totalTests"`
	TotalFailed    int       `json:"totalFailed"`
	TotalPassed    int       `json:"totalPassed"`
	TotalPending   int       `json:"totalPending"`
	TotalSkipped   int       `json:"totalSkipped"`
	Runs           string    `json:"runs"`
	BranchInfo     string    `json:"branchInfo"`
	BrowserPath    string    `json:"browserPath"`
	BrowserName    string    `json:"browserName"`
	BrowserVersion string    `json:"browserVersion"`
	OsName         string    `json:"osName"`
	OsVersion      string    `json:"osVersion"`
	CypressVersion string    `json:"cypressVersion"`
	Status         string    `json:"status"`
}

type SuiteRequest struct {
	Type  string `json:"type"`
	Suite Suite  `json:"suite"`
}

type SuiteFilter struct {
	ProjectId   string    `json:"projectId"`
	Environment string    `json:"environment"`
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`
	Status      string    `json:"status"`
}

func FilterSuite(listSuite []Suite, field string, value string) []Suite {
	var result []Suite
	switch field {
	case "Status":
		for _, suite := range listSuite {
			if suite.Status == value {
				result = append(result, suite)
			}
		}
		break
	case "Environment":
		for _, suite := range listSuite {
			if suite.Environment == value {
				result = append(result, suite)
			}
		}
		break
	case "FromDate":
		a, err := time.Parse(time.RFC3339, value)
		if err == nil {
			for _, suite := range listSuite {
				if suite.StartedTestsAt.After(a) {
					result = append(result, suite)
				}
			}
		}
		break
	case "ToDate":
		a, err := time.Parse(time.RFC3339, value)
		if err == nil {
			for _, suite := range listSuite {
				if suite.StartedTestsAt.Before(a) {
					result = append(result, suite)
				}
			}
		}
		break
	}
	return result
}
