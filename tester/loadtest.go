package loadtest
import(
"time"
"fmt"
"context"
"net/http"
"sync"
"strconv"

"github.com/scottjustin5000/loadtest/clients"
)

type LoadTestResponse struct {
  StatusCode int
  Milliseconds int
}

 type LoadTestRequest struct {
  Url string
  Concurrency  int
  MaxRequests int
  RequestTimeout int
  Body string
  ContentType string
  RequestsPerSecond int
  Type clients.RequestType
  RequestsPerSecondDuration float64
}

type LoadTest struct {
  Options LoadTestRequest
}

func NewLoadTest(options LoadTestRequest) *LoadTest {
  tester := &LoadTest{options}

  return tester
}

func makeRequests(requests int, url string)[]RequestResult {
  //client := clients.NewHttpClient(options.Type, options.Url, options.Body, options.ContentType)
 // x := make([]float64, 0)
  results := make([]RequestResult, 0)
  workers := make([]int, requests)
  wg := new(sync.WaitGroup)
  in := make(chan string, 2 * len(workers))

  for i := 0; i < len(workers); i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for i := 0; i < len(in); i++ {

        results = append(results, fetch(url, 3000))
      }
    }()
  }

  for f, _ := range workers {
    in <- "req"+strconv.Itoa(f)
  } 

  close(in)
  wg.Wait()
  return results
}

func fetch(url string, timeout int) RequestResult {
  r, _ := http.NewRequest("GET", url, nil)
  timeStart := time.Now()
  timeoutRequest, cancelFunc := context.WithTimeout(r.Context(),time.Duration(timeout) *time.Millisecond)
  defer cancelFunc()
   
  r = r.WithContext(timeoutRequest)
   
  _, err := http.DefaultClient.Do(r)

  d := time.Since(timeStart)
  timing:= d.Seconds() * float64(time.Second/time.Millisecond)
  if err != nil {
    fmt.Println("Error:", err)
    return RequestResult { url, timing, err.Error() }
  } 
  return RequestResult { url, timing, "" }
}

func doEvery(d time.Duration, options LoadTestRequest, f func(int, string)[]RequestResult)[]RequestResult {
  timeStart := time.Now()
  xx := make([]RequestResult, 0)
  for x := range time.Tick(d) {
    vv := f(options.RequestsPerSecond, options.Url)
    xx = append(xx,vv...)
    d := time.Since(timeStart)
    timing := d.Seconds() * float64(time.Second/time.Millisecond)
    if timing > options.RequestsPerSecondDuration {
      return xx
    }
   fmt.Println(x)
  }
  return xx
}

func launchPerSecond(options LoadTestRequest) []RequestResult {
  wg := new(sync.WaitGroup)
  in := make(chan string, 2 * options.Concurrency)
  results := make([]RequestResult, 0)
  for i := 0; i < options.Concurrency; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for j := 0; j < len(in); j++ {
        rps := doEvery(time.Second, options, makeRequests)
        results = append(results, rps...)
      }
    }()
  }
  for j := 0; j < options.Concurrency; j++ {
    in <- "rps"+strconv.Itoa(j)
  }
  close(in)
  wg.Wait()

  return results
}

func getConcurrency(concurrency int, maxRequests int) int {
  temp := (maxRequests - concurrency)
  if temp > 0 && maxRequests > concurrency {
    return concurrency
  } else if temp < 0 {
    return maxRequests
  }
  return 0
}

func blast(options LoadTestRequest)[]RequestResult {
  limit := options.MaxRequests
  //client := clients.NewHttpClient(options.Type, options.Url, options.Body, options.ContentType)
 // x := make([]float64, 0)
  xx := make([]RequestResult, 0)
  //y := make([]int, options.MaxRequests)
  for limit > 0 {
    //make sure we do not send more than the max request
    //var workers = 3 //options.Concurrency //getConcurrency(Options.concurrency, limit)
    workers := make([]int, getConcurrency(options.Concurrency, limit))
    wg := new(sync.WaitGroup)
    in := make(chan string, 2 * len(workers))

    for i := 0; i < len(workers); i++ {
      wg.Add(1)
      go func() {
        defer wg.Done()
        for url := range in {
          fmt.Println(url)
          //3000 should be replaced with options.RequestTimeout
          xx = append(xx, fetch("http://www.google.com", 3000))
        }
      }()
    }

    for f, _ := range workers {
      in <- strconv.Itoa(limit)+"req"+strconv.Itoa(f)
    } 

    close(in)
    wg.Wait()
    limit--
  }
  return xx

}

func(tester *LoadTest) Start() []RequestResult {
  if tester.Options.RequestsPerSecond > 0 {
    return launchPerSecond(tester.Options)
  } else {
    return blast(tester.Options)
  }
}

func (tester *LoadTest) Stop(){

}
