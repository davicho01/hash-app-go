package handlers

import (
	"app/handlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestHash_GenerateHash(t *testing.T) {

	handler := handlers.NewHash()

	req, _ := http.NewRequest("POST", "/hash", nil)
	writer := httptest.NewRecorder()
	handler.GenerateHash(writer, req)

	if writer.Body.String() == "Password Required!\n" {
		t.Error("No Password Required message")
	}

	req, _ = http.NewRequest("POST", "/hash", nil)
	req.PostForm = url.Values{}
	req.PostForm.Add("password", "angryMonkey")
	writer = httptest.NewRecorder()
	handler.GenerateHash(writer, req)

	if writer.Body.String() != "1" {
		t.Error("Return value should be 1")
	}

	req, _ = http.NewRequest("POST", "/hash", nil)
	req.PostForm = url.Values{}
	req.PostForm.Add("password", "angryMonkey")
	writer = httptest.NewRecorder()
	handler.GenerateHash(writer, req)

	if writer.Body.String() != "2" {
		t.Error("Return value should be 2")
	}

}

func TestHash_GetHashByIdentifier(t *testing.T) {

	handler := handlers.NewHash()

	req, _ := http.NewRequest("POST", "/hash", nil)
	req.PostForm = url.Values{}
	req.PostForm.Add("password", "angryMonkey")
	writer := httptest.NewRecorder()
	handler.GenerateHash(writer, req)

	if writer.Body.String() != "1" {
		t.Error("Return value should be 1")
		return
	}

	req, _ = http.NewRequest("GET", "/hash/1", nil)
	writer = httptest.NewRecorder()
	handler.GetHashByIdentifier(writer, req)

	if writer.Code != 425 || writer.Body.String() != "Password Hash Not Yet Been Processed!\n" {
		t.Error("Return value should be Password Hash Not Yet processed!")
		return
	}

	time.Sleep(6 * time.Second)

	req, _ = http.NewRequest("GET", "/hash/1", nil)
	writer = httptest.NewRecorder()
	handler.GetHashByIdentifier(writer, req)

	if writer.Code != 200 || writer.Body.String() != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Error("Return value should be ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==")
		return
	}

}

func TestHash_GetHashStats(t *testing.T) {

	handler := handlers.NewHash()

	req, _ := http.NewRequest("GET", "/stats", nil)
	writer := httptest.NewRecorder()
	handler.GetHashStats(writer, req)

	if writer.Body.String() != "{\"total\":0,\"average\":0}\n" {
		t.Error("Return value should be {\"total\":0,\"average\":0}")
		return
	}
}
