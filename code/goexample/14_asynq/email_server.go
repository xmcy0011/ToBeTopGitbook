package asynq

import (
	"context"
	"github.com/hibiken/asynq"
	"testing"
	"time"
)

var (
	kTestAsyncQKey = "email:send"
)

type EmailHandler struct {
	HasErr    error
	T         *testing.T
	ExecCount int
	Finish    bool
}

func (e *EmailHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	retried, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)

	if retried >= maxRetry {
		e.T.Log("task failed with max retry.")
		e.Finish = true
	} else {
		e.T.Log(time.Now().Format("15:04:05"), "task error, retry", ",retried:", retried, ",maxRetry:", maxRetry)
	}
}

func (e *EmailHandler) ProcessTask(context.Context, *asynq.Task) error {
	e.ExecCount++
	e.T.Log("ProcessTask", time.Now().Format("15:04:05"))

	return e.HasErr
}
