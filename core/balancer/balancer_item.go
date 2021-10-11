package balancer

import (
	"errors"
	"sync"
)

type balancerItem struct {
	backends []string
	pointer  int
	sync.Mutex
}

func (b *balancerItem) Get(key string) (string, error) {
	b.Lock()
	defer b.Unlock()
	if b.backends == nil || len(b.backends) == 0 {
		return "", errors.New("missing rounting configuration " + key)
	}
	if b.pointer >= len(b.backends) || b.pointer < 0 {
		b.pointer = 0
	}
	addr := b.backends[b.pointer]
	b.pointer = b.pointer + 1
	return addr, nil
}
