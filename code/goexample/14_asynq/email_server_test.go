package asynq

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	handler := &EmailHandler{HasErr: errors.New("test error"), T: t}

	rdsConfig := asynq.RedisClientOpt{
		Addr:     "192.168.0.164:6379",
		Password: "",
		DB:       10,
	}
	config := asynq.Config{
		Concurrency: 1,
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return time.Second * time.Duration(n)
		},
		ErrorHandler: handler,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	}

	go func() {
		srv := asynq.NewServer(rdsConfig, config)

		mux := asynq.NewServeMux()
		mux.Handle(kTestAsyncQKey, handler)
		err := srv.Run(mux)
		assert.NoError(t, err)
	}()

	client := asynq.NewClient(rdsConfig)
	task := asynq.NewTask(kTestAsyncQKey, []byte("test"))
	_, err := client.Enqueue(task, asynq.Timeout(time.Second*5), asynq.MaxRetry(3))
	assert.NoError(t, err)

	time.Sleep(time.Second * 20)
	assert.Equal(t, 3+1, handler.ExecCount)
	assert.Equal(t, true, handler.Finish)
}

func TestRateLimit(t *testing.T) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: ":6379"},
		asynq.Config{
			Concurrency: 10,
			// If error is due to rate limit, don't count the error as a failure.
			IsFailure:      func(err error) bool { return !IsRateLimitError(err) },
			RetryDelayFunc: retryDelay,
		},
	)

	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

type RateLimitError struct {
	RetryIn time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited (retry in  %v)", e.RetryIn)
}

func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

func retryDelay(n int, err error, task *asynq.Task) time.Duration {
	var ratelimitErr *RateLimitError
	if errors.As(err, &ratelimitErr) {
		return ratelimitErr.RetryIn
	}
	return asynq.DefaultRetryDelayFunc(n, err, task)
}

// Rate is 10 events/sec and permits burst of at most 30 events.
var limiter = rate.NewLimiter(10, 30)

func handler(ctx context.Context, task *asynq.Task) error {
	if !limiter.Allow() {
		return &RateLimitError{
			RetryIn: time.Duration(rand.Intn(10)) * time.Second,
		}
	}
	log.Printf("[*] processing %s", task.Payload())
	return nil
}
