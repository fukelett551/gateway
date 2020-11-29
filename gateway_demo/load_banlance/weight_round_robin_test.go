package load_banlance

import (
	"fmt"
	"testing"
)

func TestWeightRoundRobinBalance(t *testing.T) {
	wrrb := &WeightRoundRobinBalance{}
	wrrb.Add("127.0.0.1:8080", "2")
	wrrb.Add("127.0.0.1:8081", "3")
	wrrb.Add("127.0.0.1:8082", "5")
	fmt.Println("weight_round_robin")
	for i := 0; i < 10; i++ {
		ip, _ := wrrb.Get()
		fmt.Println(ip)
	}
}
