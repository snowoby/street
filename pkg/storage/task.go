package storage

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"street/ent"
	"street/pkg/d"
	"time"
)

type taskService struct {
	client *asynq.Client
}

//func NewDefaultConfig() *asynq.Client {
//	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1})
//}

func newTasker(client *asynq.Client) *taskService {
	return &taskService{
		client: client,
	}
}

func (t *taskService) imageCompress(file *ent.File) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}
	task, err := asynq.NewTask(d.StringTaskImageCompress, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
	if err != nil {
		return nil, err
	}
	return t.client.Enqueue(task)

}

func (t *taskService) avatarCompress(file *ent.File) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}
	task, err := asynq.NewTask(d.StringTaskAvatar, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
	if err != nil {
		return nil, err
	}
	return t.client.Enqueue(task)

}
