package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/imroc/req/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// Unmarshal using json.Unmarshal standard golang method
func getSingleCompany() Company {
	url := getEndpoint("/company")
	response := callApiWithRetry(url)
	var company Company
	json.Unmarshal(response.Bytes(), &company)
	return company
}

// Response handling docs: https://req.cool/docs/tutorial/handle-response/
// Unmarshal with the req method (embedded)
func getSingleCompanyReq() Company {
	url := getEndpoint("/company")
	response := callApiWithRetry(url)
	var company Company
	response.UnmarshalJson(&company)
	return company
}

func getCompanyList() []Company {
	url := getEndpoint("/companies")
	response := callApiWithRetry(url)
	var companies []Company
	response.UnmarshalJson(&companies)
	return companies
}

func getEndpoint(path string) string {
	config := getConfig()
	return config.BaseUrl + path
}

// see: https://req.cool/docs/tutorial/retry/
// this will retry until according to the retry condition
func callApiWithRetry(url string) *req.Response {
	client := req.C()
	resp, _ := client.R().
		SetRetryCount(-1).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		AddRetryHook(func(resp *req.Response, err error) {
			req := resp.Request.RawRequest
			log.Printf("Retry request: %s %s", req.Method, req.URL)
		}).
		AddRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode >= 500
		}).Get(url)
	return resp
}
