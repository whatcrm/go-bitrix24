package b24

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Create) Robots(in any) (*MainResult, error) {
	out := &MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizprocRobotAdd,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Delete) Robots(code string) (*MainResult, error) {
	out := &MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizprocRobotDel,
		In: &RequestParams{
			Code: code,
		},
		Out:    &out,
		Params: nil,
	}

	return out, c.b24.callMethod(options)
}

func (b24 *API) RobotsEventSend(in *RequestParams) (*MainResult, error) {
	out := &MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: BizProcEventSend,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	return out, b24.callMethod(options)
}
