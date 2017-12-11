package clients

import (
  "testing"
  "fmt"
)

func TestGetTimeout(t *testing.T) {
  client := NewHttpClient(ClientRequest{ReqType: GET, Url: "http://google.com", Timeout:50})
  _, err := client.MakeRequest()
  if err == nil {
    t.Errorf("expected timeout error")
  }
}

func TestGet(t *testing.T) {
  client := NewHttpClient(ClientRequest{ReqType:GET, Url: "http://google.com", Timeout:60000})
  val, err := client.MakeRequest()
  fmt.Println(val)
  if err != nil {
    t.Errorf("expected error to be nil %q", err)
  }
  if val == 0 {
    t.Errorf("expected duration to be greater than 0 %q", val)
  }
}

