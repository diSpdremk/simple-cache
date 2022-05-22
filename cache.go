package simple_cache

import (
	"github.com/diSpdremk/simple-map"
)

type CacheValue interface {
	Name() string
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

func (s *SCache) Get(k any, t CacheValue) (CacheValue, bool) {
	for _, m := range s.namedMaps {
		if m.name == t.Name() {
			findV, ok := m.sMap.Get(m.getKey(k))
			if !ok {
				return nil, false
			}
			return findV, true
		}
	}
	panic("type not register")
}

func (s *SCache) GetAllValues(t CacheValue) []CacheValue {
	for _, m := range s.namedMaps {
		if m.name == t.Name() {
			values := m.sMap.Values()
			if len(values) == 0 {
				return nil
			}
			return values
		}
	}
	panic("type not register")
}

func (s *SCache) Delete(k any, t CacheValue) {
	for _, m := range s.namedMaps {
		if m.name == t.Name() {
			rK := m.getKey(k)
			m.sMap.Delete(rK)
			return
		}
	}
}
