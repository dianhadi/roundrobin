package rr

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (h Handler) Request(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var resp *http.Response
	var err error
	// Choose the next application API to use using round robin algorithm
	nextResource := h.rrModule.GetNextResource()
	i := 0

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	for {
		// logging to show the destination, for demo purpose
		log.Printf("The request is going to %s", nextResource)

		reqURL, err := url.Parse(nextResource)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		reqURL.Path = "/"

		req, err := http.NewRequest(http.MethodPost, reqURL.String(), bytes.NewReader(reqBody))
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err = h.httpModule.Do(req)

		// Exit the loop if the request is successful
		if err == nil {
			break
		}

		log.Printf("Error sending request to %s: %v", nextResource, err)
		// Try the next application API
		nextResource = h.rrModule.GetNextResource()
		// Exit the loop if the request is same with the first one
		if i > h.maxTry {
			break
		}
		i++
	}
	if resp == nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Copy the response body to the client
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
