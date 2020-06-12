/**
 * @author Kief H. Shemul
 * @email theshemul@gmail.com
 * @create date 2020-06-12 23:40:57
 * @modify date 2020-06-12 23:40:57
 * @desc go-machinery init
 */

package main

import (
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/shemul/go-machinery/server"
	"github.com/shemul/go-machinery/utils"
	"github.com/shemul/go-machinery/worker"
	"github.com/urfave/cli"
)

var (
	app        *cli.App
	taskserver *machinery.Server
)

func init() {
	app = cli.NewApp()
	app.Name = "go-machinery"
	app.Usage = "machinery worker and send example tasks with machinery send"
	app.Author = "Kief H. Shemul"
	app.Email = "theshemul@gmail.com"

	taskserver = utils.GetMachineryServer()

}

func main() {

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the server that takes task input",
			Action: func(c *cli.Context) {
				utils.Logger.Info("server")
				server.StartServer(taskserver)
			},
		},
		{
			Name:  "worker",
			Usage: "Run the worker that will consume tasks",
			Action: func(c *cli.Context) {
				utils.Logger.Info("worker")
				worker.StartWorker(taskserver)
			},
		},
	}

	app.Run(os.Args)

}
