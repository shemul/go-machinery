package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	tasks "github.com/shemul/go-machinery/tasks"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}
func GetMachineryServer() *machinery.Server {
	Logger.Info("initing task server")

	taskserver, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})
	if err != nil {
		Logger.Fatal(err.Error())
	}

	taskserver.RegisterTasks(map[string]interface{}{
		"add": tasks.Add,
	})

	return taskserver
}
