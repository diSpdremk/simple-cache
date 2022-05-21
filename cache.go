package simple_cache

import (
	"github.com/diSpdremk/simple-map"
)

type CacheValue interface {
	Name() string
	Set(v CacheValue)
}

type SCache struct {
	namedMaps []*sNamedMap
}

func NewSCache() *SCache {
	return &SCache{}
}

type sNamedMap struct {
	name   string
	sMap   *simple_map.SMap[string, CacheValue]
	getKey func(k any) string
}

func (s *SCache) Register(name string, f func(k any) string) *SCache {
	s.namedMaps = append(s.namedMaps, &sNamedMap{
		name:   name,
		sMap:   simple_map.NewSimpleMap[string, CacheValue](),
		getKey: f,
	})
	return s
}

func (s *SCache) Put(k any, v CacheValue) {
	for _, m := range s.namedMaps {
		if m.name == v.Name() {
			m.sMap.Set(m.getKey(k), v)
			return
		}
	}
	panic("type not register")
}

func (s *SCache) Get(k any, v CacheValue) bool {
	for _, m := range s.namedMaps {
		if m.name == v.Name() {
			findV, ok := m.sMap.Get(m.getKey(k))
			if !ok {
				return false
			}
			v.Set(findV)
			return true
		}
	}
	panic("type not register")
}
