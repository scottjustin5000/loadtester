package loadtest 

import (
  "testing"
  "fmt"
)

func TestMakeRequests(t *testing.T) {
  res := makeRequests(3, "http://www.google.com")
  if len(res) < 3 {
    t.Errorf("expected 3 results")
  }
}

func TestGetFetchTimeout(t *testing.T) {
  res := fetch("http://www.google.com", 1)
  if res.Error == "" {
    t.Errorf("expected timeout error")
  }
  fmt.Println(res.Error)
}

func TestGetFetch(t *testing.T) {
  res := fetch("http://www.google.com", 3000)
  if res.Error != "" {
    t.Errorf("expected error to be undefined %q", res.Error)
  }
  if res.Timing == 0 {
    t.Errorf("expected timing to be greater than 0 %q", res.Timing)
  }
}