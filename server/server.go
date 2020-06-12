/**
 * @author Kief H. Shemul
 * @email theshemul@gmail.com
 * @create date 2020-06-12 23:40:57
 * @modify date 2020-06-12 23:40:57
 * @desc go-machinery init
 */

package server

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gofiber/fiber"
	"github.com/shemul/go-machinery/utils"
)

type payload struct {
	Value1 int64 `json:"value_1"`
	Value2 int64 `json:"value_2"`
}

func StartServer(taskserver *machinery.Server) {

	app := fiber.New()

	app.Post("/send_task", func(ctx *fiber.Ctx) {
		p := new(payload)
		err := ctx.BodyParser(p)
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		task := tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: p.Value1,
				},
				{
					Type:  "int64",
					Value: p.Value2,
				},
			},
		}

		res, err := taskserver.SendTask(&task)
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		ctx.JSON(&fiber.Map{
			"task_uuid": res.GetState().TaskUUID,
		})

	})

	app.Listen("localhost:5000")

}
