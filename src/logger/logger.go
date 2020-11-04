package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Log represent a log
type Log struct {
	Date time.Time
	URL  string
	Data interface{}
}

// LogRequest is a logging middleware to activate with martini. It gets
// request body, date and url
func LogRequest(r *http.Request) {

	l := Log{
		Date: time.Now(),
		URL:  r.URL.String(),
	}

	// keep the body content in a []byte
	b, _ := ioutil.ReadAll(r.Body)
	// rewind the body, so that json.Decoder will be able to read
	// then entire content, and we will be able to reset
	// the body with previously saved content...
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	// read the content to decode
	json.NewDecoder(r.Body).Decode(&l.Data)

	// we can now reset body for later use
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	fmt.Println(l)

}
