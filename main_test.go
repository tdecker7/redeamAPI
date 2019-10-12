package main

import (
	"testing"
	"fmt"
	"net/http"
	"bytes"
	"io"
)

var (
	baseRoute    = "http://localhost:9000"
	readRoute = fmt.Sprintf("%s/books/", baseRoute)
	createRoute  = fmt.Sprintf("%s/create-book", baseRoute)
	updateRoute  = fmt.Sprintf("%s/update-book/", baseRoute)
	deleteRoute  = fmt.Sprintf("%s/delete-book/", baseRoute)
)

func sendRequest(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	var err error
	var resp *http.Response

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("Error trying to send request: ", err)
		return resp, err
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error trying to send request: ", err)
		return resp, err
	}
	defer resp.Body.Close()
	return resp, err
}


func TestReadAllRoute200(t *testing.T) {
	// route := fmt.Sprintf("%s%s", baseRoute, readAllRoute)
	resp, err := http.Get(readRoute)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d , got status code %d\n", http.StatusOK, resp.StatusCode)
	}
}

func TestReadAllRoute405(t *testing.T) {
	var jsonStr []byte
	resp, err := http.Post(readRoute, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error("Received error in TestReadAllRoute400: ", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, found %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestReadOneRoute200(t *testing.T) {
	readRoute = fmt.Sprintf("%s%d", readRoute, 1)
	resp, err := http.Get(readRoute)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d , got status code %d\n", http.StatusOK, resp.StatusCode)
	}
}

func TestReadOneRoute405(t *testing.T) {
	var jsonStr []byte
	readRoute = fmt.Sprintf("%s%d", readRoute, 1)
	resp, err := http.Post(readRoute, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error("Received error in TestReadAllRoute400: ", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, found %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestCreateRoute200(t *testing.T) {
	jsonStr := []byte(`{"title": "test_title", "author": "test_author", "publisher": "test_publisher", "publish_date": "2019-10-12"`)
	resp, err := http.Post(createRoute, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error("Received error in TestCreateRoute200: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, found %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestCreateRoute405(t *testing.T) {
	// var jsonStr []byte
	resp, err := http.Get(createRoute)
	if err != nil {
		t.Error("Received error in TestCreateRoute400: ", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, found %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestUpdateRoute200(t *testing.T) {
	jsonStr := []byte(`{"title": "test_title", "author": "test_author", "publisher": "test_publisher", "publish_date": "2019-10-12"`)
	updateRoute = fmt.Sprintf("%s%d", updateRoute, 1)
	resp, err := http.Post(updateRoute, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error("Error found in TestUpdateRoute200: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, found %d", http.StatusOK, resp.StatusCode)
	}

}

func TestUpdateRoute405(t *testing.T) {
	resp, err := http.Get(updateRoute) 
	if err != nil {
		t.Error("Error found in TestUpdateRoute405: ", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, found %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestDeleteRoute200(t *testing.T) {
	deleteRoute = fmt.Sprintf("%s%d", deleteRoute, 1)
	resp, err := sendRequest("DELETE", deleteRoute, nil)
	if err != nil {
		t.Error("Error in TestDeleteRoute200: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, found %d", http.StatusOK, resp.StatusCode)
	}
}

func TestDeleteRoute405(t *testing.T) {
	resp, err := sendRequest("POST", deleteRoute, nil)
	if err != nil {
		t.Error("Error in TestDeleteRoute405: ", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, found %d", http.StatusOK, resp.StatusCode)
	}
}