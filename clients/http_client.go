package clients

import (
    "fmt"
    "context"
    "net/http"
    "time"
)

var cx context.Context
var cancel context.CancelFunc

type RequestType int64
 const (
  GET RequestType = iota
  POST
  PUT
)

type HttpClient struct {
  requestType RequestType
  url string 
  body string 
  contentType string
}


func NewHttpClient(requestType RequestType, url string, body string, contentType string) *HttpClient {
  tester := &HttpClient{requestType, url, body, contentType}

  return tester
}

func (c *HttpClient) MakeRequest() (float64, error) {
  cx, cancel = context.WithCancel(context.Background())
  timeStart := time.Now()
  req, _ := http.NewRequest("GET", c.url, nil)
  req = req.WithContext(cx)

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()
  fmt.Println(time.Since(timeStart), c.url)
  d := time.Since(timeStart)
  return d.Seconds() * float64(time.Second/time.Millisecond), nil
}

func Cancel() {
  if cancel != nil {
     cancel()
  }
 
}
