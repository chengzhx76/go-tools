package util

import (
	"compress/flate"
	"compress/gzip"
	"errors"
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	HEADER_CONTENT_TYPE string = "Content-Type"

	CONTENT_TYPE_JSON string = "application/json;charset=UTF-8"
	CONTENT_TYPE_FORM string = "application/x-www-form-urlencoded"
)

func GetRequest(url string, headers map[string]string) (string, error) {
	result, err := GetRequestByte(url, headers)
	if err != nil {
		log.Println("get request err", err)
		return SYMBOL_EMPTY, err
	}
	return string(result), nil
}

func PostRequest(url string, headers map[string]string, params map[string]any) (string, error) {
	result, err := PostRequestByte(url, headers, params)
	if err != nil {
		log.Println("get request err", err)
		return SYMBOL_EMPTY, err
	}
	return string(result), nil
}

func GetRequestByte(url string, headers map[string]string) ([]byte, error) {
	resp, err := GetRequestToResponse(url, headers)
	if err != nil {
		log.Println("get request error", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := switchContentEncoding(resp)
	if err != nil {
		log.Println("get encoding request err", err)
		return nil, err
	}
	result, err := io.ReadAll(body)
	if err != nil {
		log.Println("get request error", err)
		return nil, err
	}

	return result, nil
}

func switchContentEncoding(resp *http.Response) (bodyReader io.Reader, err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(resp.Body)
	case "deflate":
		bodyReader = flate.NewReader(resp.Body)
	default:
		bodyReader = resp.Body
	}
	return
}

func PostRequestByte(link string, headers map[string]string, params map[string]any) ([]byte, error) {
	resp, err := PostRequestToResponse(link, headers, params)
	if err != nil {
		log.Println("post request error", err)
		return nil, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("post request error", err)
		return nil, err
	}

	return result, nil
}

func PostRequestToResponse(link string, headers map[string]string, params map[string]any) (*http.Response, error) {
	client := &http.Client{}
	val, _ := headers[HEADER_CONTENT_TYPE]

	reqData := SYMBOL_EMPTY
	if val == CONTENT_TYPE_FORM {
		form := url.Values{}
		for key, val := range params {
			form.Add(key, val.(string))
		}
		reqData = form.Encode()
	} else {
		bte, err := JSONMarshal(params, true)
		if err != nil {
			log.Println("marshal params error", err)
			return nil, err
		}
		reqData = string(bte)
	}

	req, err := http.NewRequest(http.MethodPost, link, strings.NewReader(reqData))
	if err != nil {
		log.Println("post request error", err)
		return nil, err
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("post request error", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("post request status code<%d>", resp.StatusCode)
	}
	return resp, err
}

func GetRequestToResponse(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("get request err", err)
		return nil, err
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("get request err", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println(fmt.Sprintf("get request status code<%v>", resp.StatusCode))
		return nil, errors.New("request status code not is 200")
	}

	return resp, err
}
