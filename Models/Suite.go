package Models

import (
	"sort"
	"time"
)

type Suite struct {
	Id             string    `json:"id" bson:"id"`
	Environment    string    `json:"environment" bson:"environment"`
	StartedTestsAt time.Time `json:"startedTestsAt" bson:"startedTestsAt"`
	EndedTestsAt   time.Time `json:"endedTestsAt" bson:"endedTestsAt"`
	TotalDuration  int       `json:"totalDuration" bson:"totalDuration"`
	TotalSuites    int       `json:"totalSuites" bson:"totalSuites"`
	TotalTests     int       `json:"totalTests" bson:"totalTests"`
	TotalFailed    int       `json:"totalFailed" bson:"totalFailed"`
	TotalPassed    int       `json:"totalPassed" bson:"totalPassed"`
	TotalPending   int       `json:"totalPending" bson:"totalPending"`
	TotalSkipped   int       `json:"totalSkipped" bson:"totalSkipped"`
	Runs           string    `json:"runs" bson:"runs"`
	BranchInfo     string    `json:"branchInfo" bson:"branchInfo"`
	BrowserPath    string    `json:"browserPath" bson:"browserPath"`
	BrowserName    string    `json:"browserName" bson:"browserName"`
	BrowserVersion string    `json:"browserVersion" bson:"browserVersion"`
	OsName         string    `json:"osName" bson:"osName"`
	OsVersion      string    `json:"osVersion" bson:"osVersion"`
	CypressVersion string    `json:"cypressVersion" bson:"cypressVersion"`
	Status         string    `json:"status" bson:"status"`
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

func SortSuiteByStartDate(listSuite []Suite) []Suite {
	sort.Slice(listSuite, func(i, j int) bool {
		return listSuite[i].StartedTestsAt.After(listSuite[j].StartedTestsAt)
	})
	return listSuite
}
