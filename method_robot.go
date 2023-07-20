package b24

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Create) RobotAdd(in *interface{}) (out *MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizprocRobotAdd,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}
