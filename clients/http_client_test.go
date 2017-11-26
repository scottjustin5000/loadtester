package clients

import (
  "testing"
  "fmt"
)

func TestGet(t *testing.T) {
  client := NewHttpClient(GET, "http://google.com", "", "")
  val, err := client.MakeRequest()
  fmt.Println(val)
  if err != nil {
    t.Errorf("expected error to be nil %q", err)
  }
  if val == 0 {
    t.Errorf("expected duration to be greater than 0 %q", val)
  }
}