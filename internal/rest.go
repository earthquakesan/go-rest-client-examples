package internal

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

type Header struct {
	key   string
	value string
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// Unmarshal using json.Unmarshal standard golang method
func GetSingleCompany() Company {
	url := getEndpoint("/company")
	response := callApiWithRetry(url, "GET", nil)
	var company Company
	json.Unmarshal(response.Bytes(), &company)
	return company
}

func putSingleCompany(company Company) int {
	url := getEndpoint("/company")
	body, _ := json.Marshal(company)
	resp := callApiWithRetry(url, "PUT", body)
	return resp.StatusCode
}

// Response handling docs: https://req.cool/docs/tutorial/handle-response/
// Unmarshal with the req method (embedded)
func getSingleCompanyReq() Company {
	url := getEndpoint("/company")
	response := callApiWithRetry(url, "GET", nil)
	var company Company
	response.UnmarshalJson(&company)
	return company
}

func getCompanyList() []Company {
	url := getEndpoint("/companies")
	response := callApiWithRetry(url, "GET", nil)
	var companies []Company
	response.UnmarshalJson(&companies)
	return companies
}

func postCompanyList(companies []Company) int {
	url := getEndpoint("/companies")
	body, _ := json.Marshal(companies)
	resp := callApiWithRetry(url, "POST", body)
	return resp.StatusCode
}

func getEndpoint(path string) string {
	config := getConfig()
	return config.BaseUrl + path
}

func getJsonHeaders() []Header {
	return []Header{
		{
			key:   "Content-Type",
			value: "application/json",
		},
	}
}

func generateRequest(body []byte) *req.Request {
	client := req.C()
	request := client.R().
		SetRetryCount(-1).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		AddRetryHook(func(resp *req.Response, err error) {
			req := resp.Request.RawRequest
			if err != nil {
				log.Printf("Error: %s", err)
			}
			log.Printf("The response code from the server: %s", resp.Response.Status)
			log.Printf("Retry request: %s %s", req.Method, req.URL)
		}).
		AddRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode >= 500
		})

	headers := getJsonHeaders()
	for _, header := range headers {
		request.SetHeader(header.key, header.value)
	}

	if body != nil {
		request.SetBody(body)
	}

	return request
}

// see: https://req.cool/docs/tutorial/retry/
// this will retry until according to the retry condition
func callApiWithRetry(url, method string, body []byte) *req.Response {
	request := generateRequest(body)
	resp, err := request.Send(method, url)

	if err != nil {
		panic(err)
	}

	return resp
}
