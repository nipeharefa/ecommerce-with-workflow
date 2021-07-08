package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/workflowexecutions/v1"
)

func Home(c echo.Context) error {

	wex := &workflowexecutions.Execution{
		Argument: `{"firstName":"Sherlock","lastName":"Holmes","url":"https://jsonplaceholder.typicode.com/comments?postId=1"}`,
	}

	cc := workflowService.Projects.Locations.Workflows.Executions.Create("projects/gowirio/locations/asia-southeast1/workflows/workflow-test", wex)

	wc, err := cc.Do()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	asd := workflowService.Projects.Locations.Workflows.Executions.Get(wc.Name)
	var a *workflowexecutions.Execution
	for {
		a, err = asd.Do()
		fmt.Println(a.State, err, a.Result)
		if a.State == "SUCCEEDED" {
			break
		}
	}
	return c.JSON(http.StatusOK, a)
}
