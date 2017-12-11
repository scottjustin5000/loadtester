package clients

import (
    "fmt"
    "context"
    "net/http"
    "time"
    "io/ioutil"
    "bytes"
)

type RequestType int64

const (
  GET RequestType = iota
  POST
  PUT
)

type HttpClient struct {
  request ClientRequest
  client *http.Client
  cx context.Context
  cancel context.CancelFunc

}

func NewHttpClient(request ClientRequest) *HttpClient {
  cx, cancel := context.WithCancel(context.Background())
  client := &http.Client{
    Transport: &http.Transport{
      MaxIdleConnsPerHost: 50,
    },
    Timeout: time.Duration(request.Timeout) * time.Millisecond,
  }
  self := &HttpClient{request: request, client: client, cx:cx, cancel:cancel}
  return self
}

func(c *HttpClient) getRequestObject() (*http.Request, error) {
  var req *http.Request
  var err error
  if c.request.ReqType == GET {
    req, err = http.NewRequest("GET", c.request.Url, nil)
  } else if c.request.ReqType == POST {
    req, err := http.NewRequest("POST", c.request.Url, bytes.NewBuffer([]byte(c.request.Body)))
    if err == nil {
      req.Header.Set("Content-Type", "application/json")
    }
  } else if c.request.ReqType == PUT {
    req, err := http.NewRequest("PUT", c.request.Url, bytes.NewBuffer([]byte(c.request.Body)))
    if err == nil {
      req.Header.Set("Content-Type", "application/json")
    }
  }
  if err == nil {
    //defer c.cancel()
    req = req.WithContext(c.cx)
  }
  return req, err
}

func (c *HttpClient) MakeRequest() (float64, error) {
  req, err := c.getRequestObject()
  if err != nil {
    fmt.Println("Error Occured. %+v", err)
    return 0, err
  }
  timeStart := time.Now()
  response, err := c.client.Do(req)
  if err != nil {
    fmt.Println("Error sending request to API endpoint. %+v", err)
    return 0, err
  }
  defer response.Body.Close()
  _, err = ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Couldn't parse response body. %+v", err)
  }
  fmt.Println(time.Since(timeStart), c.request.Url)
  d := time.Since(timeStart)
  return d.Seconds() * float64(time.Second/time.Millisecond), nil
}

func (c *HttpClient)  Cancel() {
  if c.cancel != nil {
     c.cancel()
  }
 
}
