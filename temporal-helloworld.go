package helloworld

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

)

type ActivityResult struct {
  ActivityResultOne string
  ActivityResultTwo string
}

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (ActivityResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	var result1 string
	err1 := workflow.ExecuteActivity(ctx, Activity1, name).Get(ctx, &result1)
	if err1 != nil {
		logger.Error("Activity failed.", "Error 1", err1)

		return ActivityResult {
			ActivityResultOne: "",
			ActivityResultTwo: "",
		} , err1
	}

	var result2 string
	err2 := workflow.ExecuteActivity(ctx, Activity2, name).Get(ctx, &result2)
	if err2 != nil {
		logger.Error("Activity failed.", "Error 2", err2)
		return ActivityResult {
			ActivityResultOne: "",
			ActivityResultTwo: "",
		} , err1
	}

	logger.Info("HelloWorld workflow completed.", "result", result1)

	result := ActivityResult {
    ActivityResultOne: result1,
    ActivityResultTwo: result2,
  }

	return result, nil
}

func Activity1(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + " 1!", nil
}

func Activity2(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + " 2!", nil
}
