package main 
import (
"fmt"
"flag"
"strings"

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
  RequestsPerSecond int
  Type clients.RequestType
  RequestsPerSecondDuration float64
}
***/

func determineRequestType(rt string) clients.RequestType {
  if strings.EqualFold(rt, "GET") {
    return clients.GET
  } else if strings.EqualFold(rt, "POST") {
    return clients.POST
  } else if strings.EqualFold(rt, "PUT") {
    return clients.PUT
  }
  return clients.GET
}

func main() {
 urlPtr := flag.String("u", "", "url to be tested")
 concurrencyPtr := flag.Int("c", 0, "number of concurrent clients")
 maxRequestPtr := flag.Int("n", 0, "number of requests")
 requestTimeoutPtr := flag.Int("t", 60000, "timeout per requests (ms)")
 requestTypePtr := flag.String("rt", "GET", "request type")
 requestsPerSecondPtr := flag.Int("rps", 0, "requests per second")
 requestsPerSecondDurationPtr := flag.Float64("d", 60000, "request per second duration")
 bodyPtr := flag.String("b", "", "json string representation of request payload")
 
 flag.Parse()
 requestType := determineRequestType(*requestTypePtr)

 client := loadtest.NewLoadTest(loadtest.LoadTestRequest{Url:*urlPtr, RequestTimeout: *requestTimeoutPtr, Concurrency:*concurrencyPtr, MaxRequests:*maxRequestPtr, Body:*bodyPtr, RequestsPerSecond: *requestsPerSecondPtr, Type: requestType, RequestsPerSecondDuration: *requestsPerSecondDurationPtr})
 x := client.Start()
 fmt.Println(x)

}
