package load_banlance

import (
	"errors"
	"fmt"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	ipList   []*WeightNode
}

type WeightNode struct {
	originWeight int
	curWeight    int
	ipAddr       string
}

func (wrrb *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		fmt.Println("params must len 2")
		return errors.New("params len error")
	}
	weight, _ := strconv.ParseInt(params[1], 10, 64)
	w := &WeightNode{
		originWeight: int(weight),
		ipAddr:       params[0],
		curWeight:    0,
	}
	wrrb.ipList = append(wrrb.ipList, w)
	return nil
}

func (wrrb *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(wrrb.ipList); i++ {
		total += wrrb.ipList[i].originWeight
		wrrb.ipList[i].curWeight += wrrb.ipList[i].originWeight
	}
	for i := 0; i < len(wrrb.ipList); i++ {
		if best == nil || best.curWeight < wrrb.ipList[i].curWeight {
			best = wrrb.ipList[i]
		}
	}
	best.curWeight -= total
	return best.ipAddr
}

func (wrrb *WeightRoundRobinBalance) Get() (string, error) {
	return wrrb.Next(), nil
}
