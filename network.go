package main

import (
	"io/ioutil"
	"net/http"
)

func addHeaders(w http.ResponseWriter, headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
}

func proxyAndRespond(url string, w http.ResponseWriter, r *http.Request) ([]byte, http.Header, error) {
	upstreamReq, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, nil, err
	}

	res, err := http.DefaultClient.Do(upstreamReq)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	addHeaders(w, res.Header)

	_, err = w.Write(data)
	return data, res.Header, err
}
