package b24

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
	"path"
	"regexp"
	"strings"
)

func (b24 *API) getAgent(method, baseURL string, params *RequestParams) (*fiber.Agent, *fiber.Request) {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(method)

	req.Header.SetContentType(fiber.MIMEApplicationJSON)
	req.Header.SetCanonical([]byte("Authorization"), []byte("Bearer "+b24.Auth))

	a.MaxRedirectsCount(1)

	req.SetRequestURI(b24.buildURL(baseURL, params))
	//req.SetRequestURI("https://eokh00hewgqxm3a.m.pipedream.net/")
	return a, req
}

func (b24 *API) callMethod(options callMethodOptions) (err error) {

	a, req := b24.getAgent(options.Method, options.BaseURL, options.Params)

	if options.In != nil {
		b24.log("marshaling the data...")
		req, err = marshal(options.In, req)
		if err != nil {
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

	if err = errorCheck(body, status); err != nil {
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

func marshal(data interface{}, req *fiber.Request) (*fiber.Request, error) {
	m, err := json.Marshal(&data)
	if err != nil {
		return req, err
	}

	req.SetBody(m)
	return req, nil
}

func (b24 *API) buildURL(method string, params *RequestParams) string {
	// http://portal.bitrix24.com/rest/placement.bind/?access_token=sode3flffcmv500fuagrprhllx3soi72qq
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
	log.Println(u.String())
	return u.String()
}

func (b24 *API) log(message ...interface{}) {
	if b24.Debug {
		log.Println(message...)
	}
}

func isRegex(text string) bool {
	re := regexp.MustCompile("^[-a-zA-z0-9]{1,}\\.(bitrix24)\\.(ru|com|kz|kg)")
	return re.MatchString(text)
}

func errorCheck(body []byte, status int) error {
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

	log.Println("ERROR: ", e)
	return fmt.Errorf(e.Error + " " + e.ErrorDescription)
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
