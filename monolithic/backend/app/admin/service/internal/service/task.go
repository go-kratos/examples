package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"

	"kratos-monolithic-demo/app/admin/service/internal/data"
)

type TaskService struct {
	log *log.Helper

	Server *asynq.Server

	userRepo *data.UserRepo
}

func NewTaskService(
	logger log.Logger,
	userRepo *data.UserRepo,
) *TaskService {
	l := log.NewHelper(log.With(logger, "module", "task/service/admin-service"))
	return &TaskService{
		log:      l,
		userRepo: userRepo,
	}
}

// StartAllPeriodicTask 启动所有的定时任务
func (s *TaskService) StartAllPeriodicTask() {
}

// StartAllDelayTask 启动所有的延迟任务
func (s *TaskService) StartAllDelayTask() {

}
