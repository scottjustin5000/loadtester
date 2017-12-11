package clients 

import "errors"

type Client interface {
  MakeRequest() (float64, error)
  Cancel()
}



type ClientType int64
 const (
  HTTP ClientType = iota
  WS
  RPC
)

type ClientRequest struct {
  Type ClientType
  Body string
  Url string
  ReqType RequestType
  Timeout int
  Origin string
}

 func CreateClient(request ClientRequest)(Client, error) {
  switch request.Type{
  case HTTP:
    return NewHttpClient(request),nil
  default:
    //if type is invalid, return an error
    return nil, errors.New("Invalid Client Type")
  }
}