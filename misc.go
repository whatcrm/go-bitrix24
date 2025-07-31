package b24

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func (b24 *API) callMethod(options callMethodOptions) error {
	client := &http.Client{}

	if b24.Proxy != "" {
		b24.log("setting the proxy...")
		proxyURL, err := url.Parse(b24.Proxy)
		if err != nil {
			return err
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	reader := &bytes.Reader{}
  
	//добавлено для поиска по username
	if options.Params != nil && (options.Params.Filter != nil || options.Params.Select != nil) {
		b24.log("marshaling params with filter/select...")
		if err = marshal(options.Params, req); err != nil {
			return err
		}
    
    reader = r
	} else if options.In != nil {
    b24.log("marshaling the data...")
		r, err := marshal(options.In)
		if err != nil {
			return err
		}

		reader = r
	}

	req, err := http.NewRequest(options.Method, b24.buildURL(options.BaseURL, options.Params), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+b24.Auth)

	b24.log("sending the data...")
	resp, err := client.Do(req)
	if err != nil {
		b24.log("Request failed: ", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	b24.log(string(body))

	err = b24.errorCheck(body, resp.StatusCode, options)
	if err != nil && (strings.Contains(err.Error(), AccessDenied) || strings.Contains(err.Error(), UnableToGetApplicationByToken)) && b24.fallback() {
		return b24.callMethod(options)
	}

	if err != nil {
		return err
	}
	b24.log("errorCheck passed")

	if err = json.Unmarshal(body, options.Out); err != nil {
		return fmt.Errorf("400 Bad Request: %q", body)
	}

	b24.log("unmarshal passed")
	return nil
}

func statusChecker(status int) error {
	switch status {
	case http.StatusBadRequest:
		return errors.New("400 Bad Request")
	case http.StatusUnauthorized:
		return errors.New("401 Unauthorized")
	case http.StatusPaymentRequired:
		return errors.New("402 Payment Required")
	case http.StatusForbidden:
		return errors.New("403 Forbidden")
	case http.StatusNotFound:
		return errors.New("404 Not Found")
	case http.StatusCreated:
		return errors.New("201 Created")
	case http.StatusNoContent:
		return errors.New("204 No Content")
	case http.StatusOK, http.StatusAccepted, http.StatusFound, http.StatusMovedPermanently:
		return nil
	default:
		return fmt.Errorf("%d Unknown Status", status)
	}
}

func marshal(data any) (*bytes.Reader, error) {
	m, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(m), nil
}

func (b24 *API) buildURL(method string, params *RequestParams) string {
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
	return u.String()
}

func (b24 *API) log(message ...any) {
	if b24.Debug {
		log.Println(message...)
	}
}

func (b24 *API) errorCheck(body []byte, status int, options callMethodOptions) error {
	if len(body) == 0 && status == http.StatusCreated {
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
