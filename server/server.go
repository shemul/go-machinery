/**
 * @author Kief H. Shemul
 * @email theshemul@gmail.com
 * @create date 2020-06-12 23:40:57
 * @modify date 2020-06-12 23:40:57
 * @desc go-machinery init
 */

package server

import (
	base64 "encoding/base64"
	"encoding/json"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gofiber/fiber"
	task "github.com/shemul/go-machinery/tasks"
	"github.com/shemul/go-machinery/utils"
)

func StartServer(taskserver *machinery.Server) {

	app := fiber.New()

	app.Post("/send_email", func(ctx *fiber.Ctx) {
		p := new(task.Payload)
		if err := ctx.BodyParser(p); err != nil {
			utils.Logger.Fatal(err)
		}

		reqJSON, err := json.Marshal(p)
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		b64EncodedReq := base64.StdEncoding.EncodeToString([]byte(reqJSON))
		task := tasks.Signature{
			Name: "send_email",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: b64EncodedReq,
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
