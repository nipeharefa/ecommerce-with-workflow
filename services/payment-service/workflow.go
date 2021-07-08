package main

import (
	"context"
	"log"

	"google.golang.org/api/workflowexecutions/v1"
)

var workflowService *workflowexecutions.Service

func doWorkflow() {
	var err error
	ctx := context.Background()
	workflowService, err = workflowexecutions.NewService(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
