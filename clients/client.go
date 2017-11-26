package clients 
type Client interface {
  MakeRequest()
  Cancel()
}