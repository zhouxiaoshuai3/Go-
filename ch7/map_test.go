package map__test

import "testing"

func TestMapInit(t *testing.T)  {
	// [string] 是key 的类型, []string 是 value 的类型
	m1 := map[string]string{"one": "one", "two":"two", "three":"three"}
	m := map[string]int{"one":1, "two":2, "three":3}
	m2 := map[string]int{}
	t.Logf("len m=%d", len(m)) // 3
	t.Logf("len m1=%d", len(m1)) // 3
	m2["one"] = 100
	t.Logf("len m2=%d", len(m2)) // 1
	t.Log(m, m1, m2) // map[one:1 three:3 two:2] map[one:one three:three two:two] map[one:100]
	m3 := make(map[int]int, 10) // 10 是 cap  容量大小
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNoExistingkey(t *testing.T)  {
	// 声明一个空 map
	m := map[int]int{}
	// 访问一个不存在的 value
	t.Log(m[1]) // 0
	m[2] = 0
	// 访问一个 value 为 0 的key
	t.Log(m[2]) // 0
	// 如何区分访问的 value 不存在还是 值为0 呢？
	// 使用 OK  patten 判断 该 key 是否存在，OK 为 true 则存在
	if v, ok := m[1]; ok {
		t.Logf("m[1] 的 value 是 %d", v)
	} else {
		t.Log("m[1] value is  not  exists")
	}
}

func TestTravelMap(t *testing.T)  {
	m1 := map[string]string{"key1": "value1", "key2":"value2", "key3":"value3"}
	for key, val := range m1{
		t.Log(key, val)
	}
	/**
		map_test.go:39: key1 value1
	    map_test.go:39: key2 value2
	    map_test.go:39: key3 value3
	 */
}