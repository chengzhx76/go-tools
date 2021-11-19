package tool

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"log"
)

const (
	HEADER_CONTENT_TYPE string = "Content-Type"

	CONTENT_TYPE_JSON string = "application/json;charset=UTF-8"
	CONTENT_TYPE_FORM string = "application/x-www-form-urlencoded"
)

func GetRequest(url string, headers map[string]string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("get request error", err)
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("get request error", err)
	}

	//resp, err := http.Get(url)

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("get request error", err)
	}

	return string(result)
}

func GetRequestByte(url string, headers map[string]string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("get request error", err)
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("get request error", err)
	}

	//resp, err := http.Get(url)

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("get request error", err)
	}

	return result
}

func PostRequest(link string, headers map[string]string, params map[string]string) []byte {
	client := &http.Client{}
	val, _ := headers[HEADER_CONTENT_TYPE]

	reqData := ""
	if val == CONTENT_TYPE_FORM {
		form := url.Values{}
		for key, val := range params {
			form.Add(key, val)
		}
		reqData = form.Encode()
	} else {
		bte, err := JSONMarshal(params, true)
		if err != nil {
			log.Fatal("marshal params error", err)
			return nil
		}
		reqData = string(bte)
	}

	req, err := http.NewRequest("POST", link, strings.NewReader(reqData))
	if err != nil {
		log.Fatal("post request error", err)
		return nil
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("post request error", err)
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("post request error", err)
	}

	return result
}
