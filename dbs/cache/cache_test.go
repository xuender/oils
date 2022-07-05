package cache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs/cache"
)

func TestCache(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[string, string](1 * time.Second)

	_ = data.Set("key1", "value1")
	_ = data.Set("key2", "value2")

	assert.Equal(t, 2, data.Size(), "Size")

	value, has := data.Get("key2")
	assert.Equal(t, "value2", value, "Get")
	assert.True(t, has, "Get ok")

	value, has = data.GetString("keyx")
	assert.False(t, has, "not find")
	assert.Equal(t, "", value)

	_ = data.Set("key2", "value3")
	value, has = data.Get("key2")
	assert.Equal(t, "value3", value, "Get")
	assert.True(t, has, "Get ok")

	_ = data.Del("key1")
	assert.Equal(t, 1, data.Size(), "Size")

	assert.Equal(t, 1, len(data.Keys()), "Keys")

	time.Sleep(2 * time.Second)
	assert.Equal(t, 0, data.Size(), "time")
}

func BenchmarkCacheSize(b *testing.B) {
	c := cache.NewCache[int, int](20 * time.Second)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Size()
	}
}

func BenchmarkCacheSet(b *testing.B) {
	c := cache.NewCache[string, string](20 * time.Second)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = c.Set("key1", "value1")
	}
}

func BenchmarkMapPut(b *testing.B) {
	c := map[string]string{}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c["key1"] = "value1"
	}
}

func ExampleCache() {
	data := cache.NewCache[string, string](3 * time.Second)

	_ = data.Set("key1", "value1")
	_ = data.SetByDuration("key2", "value2", time.Second)
	_ = data.Set("key3", "value3")

	fmt.Println(data.Get("key1"))
	fmt.Println("init size:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println(data.Get("key3")) // reset expire time.
	fmt.Println("2 second:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println("4 second:", data.Size())

	// Output:
	// value1 true
	// init size: 3
	// value3 true
	// 2 second: 2
	// 4 second: 1
}

func ExampleNewCallbackCache() {
	data := cache.NewCallbackCache(3*time.Second, func(key, value string) {
		fmt.Println("del:", key, value)
	})

	_ = data.Set("key1", "value1")
	_ = data.SetByDuration("key2", "value2", time.Second)
	_ = data.Set("key3", "value3")

	fmt.Println(data.Get("key1"))
	fmt.Println("init size:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println(data.Get("key3")) // reset expire time.
	fmt.Println("2 second:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println("4 second:", data.Size())

	// Output:
	// value1 true
	// init size: 3
	// value3 true
	// del: key2 value2
	// 2 second: 2
	// del: key1 value1
	// 4 second: 1
}

func TestCacheSet1(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[string, string](2 * time.Second)
	_ = data.Set("3", "3")

	time.Sleep(3 * time.Second)
	assert.Equal(t, 0, data.Size(), "exists")
}

func TestCacheSet2(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[string, string](time.Second)
	_ = data.Set("3", "3")

	time.Sleep(2 * time.Second)

	assert.Equal(t, 0, data.Size(), "exists")
	_ = data.SetByTime("3", "3", time.Now().Add(time.Minute))
	time.Sleep(2 * time.Second)
	assert.Equal(t, 1, data.Size(), "exists")
}

func TestCacheSetByTime(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](time.Minute)
	assert.Nil(t, data.SetByTime(1, 2, time.Now()))
	data.Close()
	assert.NotNil(t, data.SetByTime(1, 2, time.Now()))
}

func TestCacheSetByDuration(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](time.Minute)
	assert.Nil(t, data.SetByDuration(1, 2, time.Second))
	data.Close()
	assert.NotNil(t, data.SetByDuration(1, 2, time.Second))
}

func TestCacheDel(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](time.Minute)
	_ = data.Set(1, 1)
	_ = data.Set(2, 2)
	assert.Equal(t, 2, data.Size())
	assert.Nil(t, data.Del(1))
	assert.Equal(t, 1, data.Size())
	data.Close()
	assert.NotNil(t, data.Del(2))
}

func TestCacheRest(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](2 * time.Second)
	_ = data.Set(1, 1)
	assert.Equal(t, 1, data.Size())
	time.Sleep(1 * time.Second)
	assert.Nil(t, data.Reset(1))
	time.Sleep(1 * time.Second)
	assert.Nil(t, data.Reset(1))
	time.Sleep(1 * time.Second)
	assert.Nil(t, data.Reset(1))
	assert.Equal(t, 1, data.Size())
	data.Close()
	assert.NotNil(t, data.Reset(1))
}

func TestCacheSet3(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](2 * time.Second)
	_ = data.Set(1, 1)

	time.Sleep(1 * time.Second)
	data.Get(1)
	time.Sleep(1 * time.Second)
	data.Get(1)
	time.Sleep(1 * time.Second)
	assert.Equal(t, 1, data.Size(), "LRU")
}

func TestCacheGetOnly(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](2 * time.Second)
	_ = data.Set(1, 1)

	time.Sleep(1 * time.Second)
	data.GetOnly(1)
	time.Sleep(1 * time.Second)
	data.GetOnly(1)
	time.Sleep(1 * time.Second)
	assert.Equal(t, 0, data.Size())
}

func TestCacheGetOrSet(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](2 * time.Second)
	_ = data.Set(1, 1)
	one, has := data.GetOrSet(1, 2)
	assert.Equal(t, 1, one)
	assert.True(t, has)

	two, has := data.GetOrSet(2, 2)
	assert.Equal(t, 2, two)
	assert.False(t, has)
}

func TestCacheGetOrLoad(t *testing.T) {
	t.Parallel()

	data := cache.NewCache[int, int](2 * time.Second)
	_ = data.Set(1, 1)
	one, has := data.GetOrLoad(1, func(key int) int { return 2 })
	assert.Equal(t, 1, one)
	assert.True(t, has)

	two, has := data.GetOrLoad(2, func(key int) int { return 2 })
	assert.Equal(t, 2, two)
	assert.False(t, has)
}

func TestCacheSetClose(t *testing.T) {
	t.Parallel()

	cach := cache.NewCache[int, int](time.Minute)
	assert.Nil(t, cach.Set(1, 1))
	cach.Close()
	assert.NotNil(t, cach.Set(1, 2))
}

func TestCacheSetString(t *testing.T) {
	t.Parallel()

	cach := cache.NewCache[int, int](2 * time.Second)

	_ = cach.Set(1, 1)
	str, _ := cach.GetString(1)
	assert.Equal(t, "1", str)
}
