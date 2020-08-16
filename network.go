package main

import (
	"io/ioutil"
	"net/http"
)

func proxyAndRespond(url string, w http.ResponseWriter, r *http.Request) ([]byte, error) {
	upstreamReq, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(upstreamReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(data)
	return data, err
}
