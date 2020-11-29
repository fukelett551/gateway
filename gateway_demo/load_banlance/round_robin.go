package load_banlance

import (
	"errors"
	"fmt"
)

type RoundRobinBalance struct {
	curIndex int
	ipList   []string
}

func (rr *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		fmt.Println("params must g 0")
		return errors.New("params len error")
	}
	rr.ipList = append(rr.ipList, params[0])
	return nil
}

func (rr *RoundRobinBalance) Next() string {
	if len(rr.ipList) == 0 {
		return ""
	}
	if rr.curIndex >= len(rr.ipList) {
		rr.curIndex = 0
	}
	ip := rr.ipList[rr.curIndex]
	rr.curIndex = (rr.curIndex + 1) % len(rr.ipList)
	return ip
}

func (rr *RoundRobinBalance) Get() (string, error) {
	return rr.Next(), nil
}
