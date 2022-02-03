package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"os"
	"street/ent"
	"street/pkg/data/value"
	"time"
)

type Task struct {
	client *asynq.Client
}

func NewDefaultConfig() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1})
}

func New(client *asynq.Client) *Task {
	return &Task{
		client: client,
	}
}

func (t *Task) ImageCompress(file *ent.File) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}
	task, err := asynq.NewTask(value.StringImageCompress, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
	if err != nil {
		return nil, err
	}
	return t.client.Enqueue(task)

}
