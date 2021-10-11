package balancer

import (
	"errors"
	"regexp"
	"strings"
	"sync"
)

type LoadBalancer struct {
	routes map[string]*balancerItem
	sync.Mutex
}

func NewLoadBalancer() *LoadBalancer {
	balancer := &LoadBalancer{}
	return balancer
}

func (b *LoadBalancer) Get(key string) (string, error) {
	if b.routes == nil {
		return "", errors.New("missing rounting configuration")
	}
	b.Lock()
	defer b.Unlock()
	for pattern, r := range b.routes {
		pattern = strings.TrimLeft(pattern, "/")
		pattern = strings.TrimRight(pattern, "/")
		match, err := regexp.MatchString(pattern, key)
		if match && err == nil {
			return r.Get(key)
		}
	}
	return "", errors.New("Not Found")
}

func (b *LoadBalancer) Add(key string, data []string) {
	b.Lock()
	defer b.Unlock()
	if b.routes == nil {
		b.routes = make(map[string]*balancerItem)
	}
	ba := balancerItem{}
	ba.backends = make([]string, 0)
	ba.pointer = 0
	b.routes[key] = &ba
	for _, r := range data {
		b.AddToBackend(key, r)
	}
}

func (b *LoadBalancer) AddToBackend(route string, address string) {
	if b.routes == nil {
		b.routes = make(map[string]*balancerItem)
	}
	ba, ok := b.routes[route]
	if !ok {
		b.routes = map[string]*balancerItem{}
	}
	if ba == nil {
		ba = &balancerItem{}
		b.routes[route] = ba
	}
	if ba.backends == nil {
		ba.backends = make([]string, 0)
	}
	if !isFound(route, ba.backends) {
		ba.backends = append(ba.backends, address)
	}
}

func isFound(k string, data []string) bool {
	for _, v := range data {
		if v == k {
			return true
		}
	}
	return false
}
