package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type args struct {
	w http.ResponseWriter
	r *http.Request
}

func (a args) Clone() args {
	v := args{
		r: a.r.Clone(context.Background()),
		w: httptest.NewRecorder(),
	}
	return v
}

func TestHashCreator(t *testing.T) {

	hpReq, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/hash", bytes.NewReader([]byte("password=angryMonkey")))
	hpArgs := args{
		w: httptest.NewRecorder(),
		r: hpReq,
	}
	hpArgs.r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
			args: hpArgs.Clone(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			HashCreator(tt.args.w, tt.args.r)
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
