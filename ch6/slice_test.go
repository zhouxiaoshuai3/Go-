package array_test

import "testing"

func TestSliceInit(t *testing.T)  {
	var s []int
	// 打印 slice 的 len 长度 和 cap 容量
	t.Log(len(s), cap(s)) // 0 0
	// 往 slice 里添加元素,并打印 len cap
	s1 := append(s, 1)
	t.Log(len(s1), cap(s1)) //1 1
	// 声明并初始化slice
	s2 := []int{1, 2, 3, 4, 5}
	t.Log(s2) // [1 2 3 4 5]
	// 使用 make 创建 slice
	// 使用make 创建 slice ， 第一个参数是 类型[type]，第二个参数是 slice 的长度 len ,
	// 第三个参数 是 slice 的容量 cap
	s3 := make([]int, 3, 5)
	t.Log(len(s3), cap(s3)) // 3 5
}

func TestSliceGrowing(t *testing.T)  {
	// slice 是如何可变长的
	s := []int{}
	for i := 0; i < 10 ;i++  {
		s = append(s, i)
		t.Log(len(s), cap(s))
		// 当 slice 中的 len 超过原有的 cap 时，cap 的增长是前一次的 2 倍
		// 当 slice 在不断增长是，在内部会创建一个新的存储空间来保存 新的slice，
		// 并把原有的存储空间拷贝到新的存储空间中 所以需要赋值给新的slice
		/**
		 	slice_test.go:27: 1 1
		    slice_test.go:27: 2 2
		    slice_test.go:27: 3 4
		    slice_test.go:27: 4 4
		    slice_test.go:27: 5 8
		    slice_test.go:27: 6 8
		    slice_test.go:27: 7 8
		    slice_test.go:27: 8 8
		    slice_test.go:27: 9 16
		    slice_test.go:27: 10 16
		 */
	}
}

func TestSliceShareMemory(t *testing.T)  {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]

	t.Log(Q2, len(Q2), cap(Q2))
	// [Apr May Jun] 3 9  len(Q2) 是 3 容易理解，截取下标是 3 的元素到下标是 6 的元素[不包含]
	// cap(Q2) 为什么是 9 呢？ 是因为 从 数组下标 3 开始到数组的最后一个元素，指向了内部的连续存储空间 一个 9 个元素；
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// [Jun Jul Aug] 3 7 同理
	summer[0] = "Unkown"
	t.Log(Q2) // [Apr May Unkown] 两个切片 共同的部分被影响；因为修改了同一个存储空间
	t.Log(year) // [Jan Feb Mar Apr May Unkown Jul Aug Sep Oct Nov Dec]
}

func TestSliceComparing(t *testing.T)  {
	s := []int{}
	s1 := []int{}
	t.Log(s == s1) // invalid operation: s == s1 (slice can only be compared to nil)
}