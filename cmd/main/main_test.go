package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	go main()

	time.Sleep(1 * time.Second)

	res, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Status code was %d; want 200", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	if string(body) != "Hello World!" {
		t.Errorf("Body was %q; want \"Hello World!\"", string(body))
	}
}
