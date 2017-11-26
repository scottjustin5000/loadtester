package main 
import (
"fmt"

"github.com/scottjustin5000/loadtest/tester"
"github.com/scottjustin5000/loadtest/clients"
) 

/**
 * Options:
 *  - Url: web address to test
 *  - Concurrency: how many concurrent clients.
 *  - MaxRequests: total requests to send
 *  - RequestTimeout: how long to allow individual requests to execute
 *  - RequestsPerSecondDuration: how long to run the tests.
 *  - Type: request-type to use: GET, POST, PUT.
 *  - Body: the contents to send along a POST or PUT request.
 *  - ContentType: the MIME type to use for the body.
 *  - RequestsPerSecond: number of requests per second to send.
 */

 /* type LoadTestRequest struct {
  Url string
  Concurrency  int
  MaxRequests int
  RequestTimeout int
  Body string
  ContentType string
  RequestsPerSecond int
  Type clients.RequestType
  RequestsPerSecondDuration float64
}*/

func main() {
 client := loadtest.NewLoadTest(loadtest.LoadTestRequest{Url:"https://google.com", Concurrency:2, MaxRequests:0, Body:"", ContentType: "", RequestsPerSecond:2, Type: clients.GET})
 x := client.Start()
 fmt.Println(x)

}
