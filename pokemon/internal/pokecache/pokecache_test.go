package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val string
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: "some test data",
		},
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			val: "some more test data",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, []byte(c.val))

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expeced to find key %q in cache", c.key)
				return
			}
			if string(val) != c.val {
				t.Errorf("expected value %q, got %q", c.val, string(val))
			}
		})
	}
}
