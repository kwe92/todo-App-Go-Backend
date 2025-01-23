package utilities

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetId(t *testing.T) {
	idString := GetId()

	received, err := strconv.Atoi(idString)

	CheckError(err)

	fmt.Println("Received ID: ", received)

	if received > 1000 {
		t.Fatal("expected: value less than 1000 received: ", received)
	}
}

func TestGetDate(t *testing.T) {

	expected := time.Now().Format("01-02-2006")
	received := GetDate()

	fmt.Println("expected date: ", expected)
	fmt.Println("received date: ", received)

	CompareTestResults(t, expected, received)

}

func TestParseURL(t *testing.T) {

	host := "http://www.example.com"

	path := "/software/index.html"

	req := &http.Request{
		Host: host,
		URL: &url.URL{
			Host: host,
			Path: path,
		},
	}
	expected := host + path

	received := ParseURL(req)

	fmt.Println("expected url: ", expected)
	fmt.Println("received url: ", received)

	CompareTestResults(t, expected, received)

}

func TestSetHeader(t *testing.T) {
	resp := httptest.NewRecorder()

	fmt.Println(resp.Header())

	SetHeader(resp)

	expected := http.Header{"Content-Type": []string{"application/json"}}

	actual := resp.Header()

	fmt.Println("expected Header:", expected)
	fmt.Println("actual Header:", actual)

	assert.Equal(t, expected, actual)

}