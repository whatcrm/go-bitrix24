package b24

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Create) Robots(in any) (out *MainResult, err error) {
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

func (c *Delete) Robots(code string) (out *MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizprocRobotDel,
		In: &RequestParams{
			Code: code,
		},
		Out:    &out,
		Params: nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (b24 *API) RobotsEventSend(in *RequestParams) (out *MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizProcEventSend,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = b24.callMethod(options)
	return
}
