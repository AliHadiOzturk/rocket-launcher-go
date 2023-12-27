package utils

import (
	"io"
	"log"
	"net/http"
)

type Http struct {
	retryCount int
}

func NewHttpUtil() *Http {
	return &Http{
		retryCount: 0,
	}
}

func (h Http) Get(url string, headers map[string]string) []byte {
	h.retryCount++
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating request")
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic("Error while sending request")
	}

	if resp.StatusCode != 200 {
		if h.retryCount == 3 {
			return nil
		}
		return h.Get(url, headers)
	}

	body, err := io.ReadAll(io.Reader(resp.Body))

	if err != nil {
		log.Fatalln("Error while reading body")
		return nil
	}

	return body
}
