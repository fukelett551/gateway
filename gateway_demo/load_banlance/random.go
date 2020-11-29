package load_banlance

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RandomBalance struct {
	curIndex int
	ipList   []string
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		fmt.Println("params must g 0")
		return errors.New("params len error")
	}
	r.ipList = append(r.ipList, params[0])
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.ipList) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	r.curIndex = rand.Intn(len(r.ipList))
	return r.ipList[r.curIndex]
}

func (r *RandomBalance) Get() (string, error) {
	return r.Next(), nil
}
