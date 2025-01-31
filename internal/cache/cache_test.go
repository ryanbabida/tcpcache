package cache

import "testing"

func TestNewCache(t *testing.T) {
	c, err := NewCache[string, int](10, func(key string) int { return 0 })
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(c.sets) != 10 {
		t.Fatalf("expected: %d, actual: %d", 10, len(c.sets))
	}
}

func TestNewCache_InvalidSetCount(t *testing.T) {
	c, err := NewCache[string, int](0, func(key string) int { return 0 })
	if err == nil {
		t.Errorf("should be error due to invalid set count")
	}

	if c != nil {
		t.Errorf("expected: %v, actual: %v", nil, c)
	}
}

func TestCacheSet(t *testing.T) {
	c, err := NewCache[string, int](10, func(key string) int { return 0 })
	if err != nil {
		t.Errorf("%v", err)
	}

	c.Set("key", 1)
	val, ok := c.Get("key")
	if !ok {
		t.Errorf("%v", err)
	}

	if val != 1 {
		t.Fatalf("expected: %d, actual: %d", 1, val)
	}
}

func TestCacheDelete(t *testing.T) {
	c, err := NewCache[string, int](10, func(key string) int { return 0 })
	if err != nil {
		t.Errorf("%v", err)
	}

	c.Set("key", 1)
	val, ok := c.Delete("key")
	if !ok {
		t.Errorf("%v", err)
	}

	if val != 1 {
		t.Fatalf("expected: %d, actual: %d", 1, val)
	}
}

func TestCacheDelete_NotFound(t *testing.T) {
	c, err := NewCache[string, int](10, func(key string) int { return 0 })
	if err != nil {
		t.Errorf("%v", err)
	}

	_, ok := c.Delete("key")
	if ok {
		t.Errorf("%v", err)
	}
}
