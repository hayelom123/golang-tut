package main


import "fmt"

type Client struct{
     Timeout int
}

// Define a function type that modifies the client
type Option func(*Client) 

// Factory function accepting variadic options
func NewClient(opts ...Option) *Client{
	 client:=&Client{Timeout:30}// 1. Set the hardcoded defaults first

	for _,opt := range opts{
			opt(client)        // 2. Override defaults with user options
	}
	return client
}

// Exported option modifier

func WithTimeout(t int) Option{
	return func(c *Client){
		  c.Timeout = t
	}
}



func main(){
 c1 := NewClient()                 // Uses default (Timeout: 30)
 c2 := NewClient(WithTimeout(60))  // Overrides default (Timeout: 60)

 fmt.Println("Default:",c1.Timeout)
 fmt.Println("Overrides default:",c2.Timeout)
}