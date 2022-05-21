package simple_cache

import (
	"testing"
)

const (
	testParam = 100
	testKey   = "1"
	cacheName = "test"
)

type TestClass struct {
	TestParam int
}

func (t *TestClass) Name() string {
	return cacheName
}

func (t *TestClass) Set(value CacheValue) {
	*t = *(value.(*TestClass))
}

func TestCache(t *testing.T) {
	c := NewSCache().Register(cacheName, func(k any) string {
		return k.(string)
	})
	c.Put(testKey, &TestClass{testParam})
	var testObject TestClass
	c.Get(testKey, &testObject)
	if testObject.TestParam != testParam {
		t.Failed()
	}
}
