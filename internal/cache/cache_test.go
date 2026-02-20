package cache

import (
	"fmt"
	"testing"
	"time"
)
func TestCacheAddGet(t *testing.T) {
	const cacheResetInterval = 5 * time.Second
	cases := []struct{
		key string
		value []byte
	}{
		{
			key: "https://google.com",
			value: []byte("the OG browser"),
		},
		{
			key: "https://youtube.com",
			value: []byte("the worlds best video platform!"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cacher := NewCache(cacheResetInterval)
			cacher.Add(c.key, c.value)
			val, ok := cacher.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

	cacher := NewCache(cacheResetInterval)
	if _, ok := cacher.Get("non existent"); ok {
		t.Errorf("expected ok to be false for non existing key")
	}

}
func TestReapLoop(t *testing.T) {
	const cacheResetInterval = 5 * time.Second
	const waitTime = cacheResetInterval + 5 * time.Millisecond
	cacher := NewCache(cacheResetInterval)

	//just testing the Add() and Get() BEFORE the cacheResetInterval erases all cached data
	testKey := "https://google.com"
	cacher.Add(testKey, []byte("the OG browser"))
	_, ok := cacher.Get(testKey)
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	
	//Now testing the Add() and Get() AFTER the cacheResetInterval erases all cached data
	time.Sleep(waitTime)
	_, ok = cacher.Get(testKey)
	if ok {
		t.Errorf("expected Reaploop() to erase %s", testKey)
	}
}
	