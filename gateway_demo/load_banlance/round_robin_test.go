package load_banlance

import (
	"fmt"
	"testing"
)

func TestRoundRobinBalance(t *testing.T) {
	rrb := &RoundRobinBalance{}
	rrb.Add("127.0.0.1:8080")
	rrb.Add("127.0.0.1:8081")
	rrb.Add("127.0.0.1:8082")
	rrb.Add("127.0.0.1:8083")
	rrb.Add("127.0.0.1:8084")
	rrb.Add("127.0.0.1:8085")
	fmt.Println("roundrobin")
	for i := 0; i < 10; i++ {
		ip, _ := rrb.Get()
		fmt.Println(ip)
	}
}
