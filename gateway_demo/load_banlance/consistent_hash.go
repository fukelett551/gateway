package load_banlance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type Uint32Slice []uint32

func (u Uint32Slice) Len() int {
	return len(u)
}
func (u Uint32Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

func (u Uint32Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type ConsistentHashBalance struct {
	mux      sync.RWMutex
	hash     Hash
	replicas int         //虚拟节点的个数
	keys     Uint32Slice //已排序的key集合
	hashMap  map[uint32]string
}

func NewConsistentHashBalance(replicas int, hash Hash) *ConsistentHashBalance {
	m := &ConsistentHashBalance{
		replicas: replicas,
		hash:     hash,
		hashMap:  make(map[uint32]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (ch *ConsistentHashBalance) isEmpty() bool {
	return len(ch.keys) == 0
}

func (ch *ConsistentHashBalance) Add(params ...string) error {
	if len(params) != 1 {
		fmt.Println("params len must 1")
		return errors.New("params len must 1")
	}

	ch.mux.Lock()
	defer ch.mux.Unlock()
	for i := 0; i < ch.replicas; i++ {
		hashValue := ch.hash([]byte(strconv.Itoa(i) + params[0]))
		ch.keys = append(ch.keys, hashValue)
		ch.hashMap[hashValue] = params[0]
	}
	sort.Sort(ch.keys)
	return nil
}

func (ch *ConsistentHashBalance) Next(key string) string {
	if ch.isEmpty() {
		return ""
	}
	hashValue := ch.hash([]byte(key))

	idx := sort.Search(len(ch.keys), func(i int) bool { return ch.keys[i] >= hashValue })

	if idx == len(ch.keys) {
		idx = 0
	}
	ch.mux.RLock()
	defer ch.mux.RUnlock()
	return ch.hashMap[ch.keys[idx]]
}

func (ch *ConsistentHashBalance) Get(key string) (string, error) {
	return ch.Next(key), nil
}
