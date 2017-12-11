package loadtest
import(
"time"
"fmt"
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
  RequestsPerSecond int
  Type clients.RequestType
  RequestsPerSecondDuration float64
}

type LoadTest struct {
  Options LoadTestRequest
  client clients.Client
}

func NewLoadTest(options LoadTestRequest) *LoadTest {
  client, _ := clients.CreateClient(clients.ClientRequest{Type: clients.HTTP, Body: options.Body, Url: options.Url, ReqType: options.Type, Timeout: options.RequestTimeout})
  tester := &LoadTest{options, client}
  return tester
}

func getConcurrency(concurrency int, maxRequests int) int {
  temp := (maxRequests - concurrency)
  if temp > 0 && maxRequests > concurrency {
    return concurrency
  } else if temp <= 0 {
    return maxRequests
  }
  return 0
}

func fetch(client clients.Client, url string) RequestResult {
  
  timing, err:= client.MakeRequest()
  if err != nil {
    return RequestResult { url, timing, err.Error() }
  } 
  return RequestResult { url, timing, "" }
}

func (tester *LoadTest) makeRequests()[]RequestResult {
  var options = tester.Options
  results := make([]RequestResult, 0)
  workers := make([]int, options.RequestsPerSecond)
  wg := new(sync.WaitGroup)
  in := make(chan string, 2 * len(workers))

  for i := 0; i < len(workers); i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
        results = append(results, fetch(tester.client, options.Url))
    }()
  }

  for f, _ := range workers {
    in <- "req"+strconv.Itoa(f)
  } 

  close(in)
  wg.Wait()
  return results
}

func(tester *LoadTest) doEvery(d time.Duration, f func()[]RequestResult)[]RequestResult {
  timeStart := time.Now()
  results := make([]RequestResult, 0)
  for x := range time.Tick(d) {
    chunkedResult := f()
    results = append(results,chunkedResult...)
    d := time.Since(timeStart)
    timing := d.Seconds() * float64(time.Second/time.Millisecond)
    if timing > tester.Options.RequestsPerSecondDuration {
      return results
    }
   fmt.Println(x)
  }
  return results
}

func(tester *LoadTest) launchPerSecond() []RequestResult {
  wg := new(sync.WaitGroup)
  in := make(chan string, 2 * tester.Options.Concurrency)
  results := make([]RequestResult, 0)
  for i := 0; i < tester.Options.Concurrency; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for j := 0; j < len(in); j++ {
        rps := tester.doEvery(time.Second, tester.makeRequests)
        results = append(results, rps...)
      }
    }()
  }
  for j := 0; j < tester.Options.Concurrency; j++ {
    in <- "rps"+strconv.Itoa(j)
  }
  close(in)
  wg.Wait()

  return results
}

func(tester *LoadTest) blast()[]RequestResult {
  limit := tester.Options.MaxRequests

  results := make([]RequestResult, 0)

  for limit > 0 {
    workers := make([]int, getConcurrency(tester.Options.Concurrency, limit))
    wg := new(sync.WaitGroup)
    in := make(chan string, 2 * len(workers))

    for i := 0; i < len(workers); i++ {
      wg.Add(1)
      go func() {
        defer wg.Done()
        for url := range in {
          fmt.Println(url)
          results = append(results, fetch(tester.client, tester.Options.Url))
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
  return results
}

func(tester *LoadTest) Start() []RequestResult {
  if tester.Options.RequestsPerSecond > 0 {
    return tester.launchPerSecond()
  } else {
    return tester.blast()
  }
}

func (tester *LoadTest) Stop(){

}
