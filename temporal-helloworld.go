package helloworld

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (result string, err error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	var resultOne string
	err = workflow.ExecuteActivity(ctx1, ActivityOne, name).Get(ctx1, &resultOne)
	if err != nil {
		logger.Error("Activity One failed.", "Error", err)
		return "", err
	}

	ctx2 := workflow.WithActivityOptions(ctx, ao)

	var resultTwo string
	err = workflow.ExecuteActivity(ctx2, ActivityTwo, name).Get(ctx2, &resultTwo)
	if err != nil {
		logger.Error("Activity Two failed.", "Error", err)
		return "", err
	}

	logger.Info("HelloWorld workflow completed.", "result", resultTwo)

	return resultTwo, nil
}

func ActivityOne(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + " 1!", nil
}

func ActivityTwo(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + " 2!", nil
}