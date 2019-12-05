package goverisure

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	username string
	password string
	apiURL   string
	cookie   *http.Cookie
}

var (
	apiURL = "https://e-api02.verisure.com"
)

const (
	urlLogin            = "%s/xbn/2/cookie"
	urlGetInstallations = "%s/xbn/2/installation/search?email=%s"
	urlGetOverview      = "%s/xbn/2/installation/%s/overview"
)

func New() *Handler {
	username, ok := os.LookupEnv("VERISURE_LOGIN")
	if !ok {
		panic("missing environment key: VERISURE_LOGIN")
	}

	password, ok := os.LookupEnv("VERISURE_PASSWORD")
	if !ok {
		panic("missing environment key: VERISURE_PASSWORD")
	}

	return &Handler{
		username: username,
		password: password,
		apiURL:   apiURL,
	}
}

func (h *Handler) Get(method, url string, args ...interface{}) ([]byte, error) {

	req, err := http.NewRequest(method, fmt.Sprintf(url, args...), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")

	if url == urlLogin {
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CPE/%s:%s", h.username, h.password)))))
	}

	if h.cookie != nil {
		req.AddCookie(h.cookie)
	}

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return nil
	}

	log.Printf("Request: %+v", req)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	for _, cookie := range res.Cookies() {
		if cookie.Name == "vid" {
			h.cookie = cookie
		}
	}
	body, _ := ioutil.ReadAll(res.Body)
	log.Printf("body: %s", body)

	return body, nil
}

func (h *Handler) Login() error {
	// get a cookie if we have none
	if h.cookie == nil {
		if _, err := h.Get("POST", urlLogin, h.apiURL); err != nil {
			return err
		}
	}
	// get new cookie if it expired
	if h.cookie.Expires.Before(time.Now()) {
		if _, err := h.Get("POST", urlLogin, h.apiURL); err != nil {
			return err
		}
	}
	// we already have a non-expired cookie
	return nil
}
