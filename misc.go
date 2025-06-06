package b24

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (b24 *API) getAgent(method, baseURL string, params *RequestParams) (*fiber.Agent, *fiber.Request) {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(method)

	req.Header.SetContentType(fiber.MIMEApplicationJSON)
	req.Header.SetCanonical([]byte("Authorization"), []byte("Bearer "+b24.Auth))

	a.MaxRedirectsCount(1)

	req.SetRequestURI(b24.buildURL(baseURL, params))
	return a, req
}

func (b24 *API) callMethod(options callMethodOptions) (err error) {
	url := b24.buildURL(options.BaseURL, options.Params)
	b24.log("Request URL:", url)

	a, req := b24.getAgent(options.Method, options.BaseURL, options.Params)

	if options.In != nil {
		b24.log("marshaling the data...")
		if err = marshal(options.In, req); err != nil {
			return
		}
	}

	b24.log("sending the data...")
	if err = a.Parse(); err != nil {
		log.Println(err)
		return
	}

	b24.log("getting the answer...")
	status, body, errs := a.Bytes()
	if errs != nil {
		log.Println("Errs: ", errs)
		err = errs[0]
		return
	}

	b24.log(string(body))

	err = b24.errorCheck(body, status, options)
	if err != nil && (strings.Contains(err.Error(), AccessDenied) || strings.Contains(err.Error(), UnableToGetApplicationByToken)) && b24.fallback() {
		return b24.callMethod(options)
	}

	if err != nil {
		return
	}
	b24.log("errorCheck passed")

	if err = json.Unmarshal(body, options.Out); err != nil {
		return fiber.NewError(400, string(body))
	}
	b24.log("unmarshal passed")

	err = statusChecker(status)
	return
}

func statusChecker(status int) error {
	switch status {
	case 400:
		return fiber.ErrBadRequest
	case 401:
		return fiber.ErrUnauthorized
	case 402:
		return fiber.ErrPaymentRequired
	case 403:
		return fiber.ErrForbidden
	case 404:
		return fiber.ErrNotFound
	case 201:
		return fiber.NewError(201, "Created")
	case 204:
		return fiber.NewError(204, "No content")
	case 200, 202, 302, 301:
		return nil
	default:
		return fiber.NewError(status, "unknown status")
	}
}

func marshal(data any, req *fiber.Request) error {
	m, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	req.SetBody(m)
	return nil
}

func (b24 *API) buildURL(method string, params *RequestParams) string {
	if b24.WebhookURL != "" {
		u, _ := url.Parse(b24.WebhookURL)
		u.Path = path.Join(u.Path, method+".json")
		b24.log("Webhook URL:", u.String())
		return u.String()
	}

	// Старый код для OAuth
	b24.fixDomain()

	u, err := url.Parse(b24.Domain)
	if err != nil {
		return "can't parse domain, due to " + err.Error()
	}

	u.Scheme = scheme
	u.Path = path.Join(u.Path, rest, method)
	query := u.Query()

	if params != nil && params.ModuleID != "" {
		query.Set("moduleId", params.ModuleID)
		query.Set("id", params.ID)
	}

	if params != nil && params.RefreshToken != "" {
		u.Path = oAuthToken
		query.Set("refresh_token", params.RefreshToken)
		query.Set("grant_type", "refresh_token")
		query.Set("client_id", b24.ClientID)
		query.Set("client_secret", b24.ClientSecret)
	}

	query.Set(Auth, b24.Auth)

	u.RawQuery = query.Encode()
	b24.log("Final URL:", u.String())
	return u.String()
}

func (b24 *API) log(message ...any) {
	if b24.Debug {
		log.Println(message...)
	}
}

func (b24 *API) errorCheck(body []byte, status int, options callMethodOptions) error {
	if len(body) == 0 && status == fiber.StatusCreated {
		return nil
	}

	e := ErrorResponse{}
	if err := json.Unmarshal(body, &e); err != nil {
		// if it cannot unmarshal, there is no errors in answer
		return nil
	}

	if e.Error == "" && e.ErrorDescription == "" {
		return nil
	}

	return fmt.Errorf("%s %s", e.Error, e.ErrorDescription)
}

func (b24 *API) fixDomain() {
	b24.Domain = strings.Trim(b24.Domain, " ")
	b24.Domain = strings.Trim(b24.Domain, "/")
	b24.Domain = strings.Trim(b24.Domain, ".")
	if strings.Contains(b24.Domain, "https://") || strings.Contains(b24.Domain, "http://") {
		uri, _ := url.Parse(b24.Domain)
		b24.Domain = uri.Hostname()
	}
	b24.Domain = strings.ReplaceAll(b24.Domain, "//", "/")
}

func (b24 *API) CustomMethod(method string, params interface{}) ([]byte, error) {
	var rawResponse json.RawMessage

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: method,
		In:      params,
		Out:     &rawResponse,
		Params:  nil,
	}

	if err := b24.callMethod(options); err != nil {
		return nil, err
	}

	return rawResponse, nil
}
