package cache

import "fmt"

type cache[K comparable, V any] struct {
	sets []*set[K, V]
	hash func(key K) int
}

func NewCache[K comparable, V any](setCount int, hash func(key K) int) (*cache[K, V], error) {
	if setCount < 1 {
		return nil, fmt.Errorf("NewCache - setCount cannot be less than 1")
	}

	c := cache[K, V]{
		sets: []*set[K, V]{},
		hash: hash,
	}

	for i := 0; i < setCount; i++ {
		c.sets = append(c.sets, newSet[K, V]())
	}

	return &c, nil
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	idx := c.hash(key) % len(c.sets)

	val, ok := c.sets[idx].get(key)
	return val, ok
}

func (c *cache[K, V]) Set(key K, val V) {
	idx := c.hash(key) % len(c.sets)

	c.sets[idx].set(key, val)
}

func (c *cache[K, V]) Delete(key K) (V, bool) {
	idx := c.hash(key) % len(c.sets)

	val, ok := c.sets[idx].delete(key)
	return val, ok
}
