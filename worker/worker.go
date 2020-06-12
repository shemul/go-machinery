/**
 * @author Kief H. Shemul
 * @email theshemul@gmail.com
 * @create date 2020-06-12 23:40:57
 * @modify date 2020-06-12 23:40:57
 * @desc go-machinery init
 */

package worker

import (
	"github.com/RichardKnop/machinery/v1"
)

func StartWorker(taskserver *machinery.Server) error {

	worker := taskserver.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil

}
