package load_banlance

import (
	"fmt"
	"testing"
)

func TestConsistentHashBalance(t *testing.T) {
	chb := NewConsistentHashBalance(10, nil)
	chb.Add("127.0.0.1:8080")
	chb.Add("127.0.0.1:8081")
	chb.Add("127.0.0.1:8082")
	chb.Add("127.0.0.1:8083")
	chb.Add("127.0.0.1:8084")
	chb.Add("127.0.0.1:8085")

	fmt.Println("consistent_hash")
	fmt.Println(chb.Get("127.0.0.1"))
	fmt.Println(chb.Get("192.168.0.1"))
	fmt.Println(chb.Get("127.0.0.1"))

}
