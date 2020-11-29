package load_banlance

import (
	"fmt"
	"testing"
)

func TestRanomBanlance(t *testing.T) {
	rb := &RandomBalance{}
	rb.Add("127.0.0.1:8080")
	rb.Add("127.0.0.1:8081")
	rb.Add("127.0.0.1:8082")
	rb.Add("127.0.0.1:8083")
	rb.Add("127.0.0.1:8084")
	rb.Add("127.0.0.1:8085")
	fmt.Println("random")
	for i := 0; i < 10; i++ {
		ip, _ := rb.Get()
		fmt.Println(ip)
	}
}
