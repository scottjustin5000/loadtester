package loadtest 

import (
  "testing"
  "fmt"

  "github.com/scottjustin5000/loadtest/clients"
)

func TestConcurrencyHelper(t *testing.T) {
  val := getConcurrency(10, 9)
  if val != 9 {
    t.Errorf("expected remaining requests value (9), received %v", val)
  }
  val = getConcurrency(10, 11) 
  if val != 10 {
    t.Errorf("expected concurrency value (10), received %v", val)
  }

  val = getConcurrency(10, 10) 
  if val != 10 {
    t.Errorf("expected concurrency value (10), received %v", val)
  }
}

func TestGetFetchTimeout(t *testing.T) {
  client := clients.NewHttpClient(clients.ClientRequest{Url: "http://www.google.com", Timeout: 1, ReqType: clients.GET})
  res := fetch(client, "http://www.google.com")
  if res.Error == "" {
    t.Errorf("expected timeout error")
  }
}

func TestGetFetch(t *testing.T) {
  client := clients.NewHttpClient(clients.ClientRequest{Url: "http://www.google.com", Timeout: 3000, ReqType: clients.GET})
  res := fetch(client, "http://www.google.com")
  if res.Error != "" {
    t.Errorf("expected error to be undefined %q", res.Error)
  }
  if res.Timing == 0 {
    t.Errorf("expected timing to be greater than 0 %q", res.Timing)
  }
}

func TestMakeRequests(t *testing.T) {
  cln := NewLoadTest(LoadTestRequest{Url:"http://www.google.com", RequestTimeout: 30000, Concurrency:0, MaxRequests:0, Body:"", RequestsPerSecond: 3, Type: clients.GET})
  res := cln.makeRequests()
  if len(res) != 3 {
    t.Errorf("expected 3 results")
  }
  fmt.Println(len(res))
}

