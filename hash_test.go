package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

type args struct {
	w http.ResponseWriter
	r *http.Request
}

func (a args) clone() args {
	v := args{
		r: a.r.Clone(context.Background()),
		w: httptest.NewRecorder(),
	}
	return v
}

func (a args) new(r *http.Request, p string) args {
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Form = make(url.Values)
	r.Form["password"] = []string{p}
	v := args{
		r: r,
		w: httptest.NewRecorder(),
	}
	return v
}

func TestHashCreationHandler1(t *testing.T) {

	hpReq, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/hash", bytes.NewReader([]byte("password=angryMonkey")))
	hpArgs := args{}.new(hpReq, "someValue")

	tests := []struct {
		name string
		args args
	}{
		{
			name: "should return a simple number",
			args: hpArgs,
		},
		{
			name: "should return the next integer",
			args: hpArgs.clone(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			HashCreationHandler(tt.args.w, tt.args.r)
			recorder, _ := tt.args.w.(*httptest.ResponseRecorder)
			resp := recorder.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			val, err := strconv.Atoi(string(body))
			if err != nil {
				t.Error("did not return a number")
			}
			if val != i+1 {
				t.Error("returned unexpected number")
			}
			if resp.StatusCode != http.StatusOK {
				t.Error("unexpected response status")
			}

		})
	}

}

func TestHashRetriever(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HashRetriever(tt.args.w, tt.args.r)
		})
	}
}

func TestHashEncode(t *testing.T) {
	v := "angryMonkey"
	actual := HashEncode(v)
	expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	if actual != expected {
		t.Error("password was not hashed and encoded correctly")
	}
}
