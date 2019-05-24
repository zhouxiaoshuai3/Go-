package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	// 声明数组，初始化零值
	var arr [3]int
	// 声明数组并且初始化
	arr1 := [3]int {1, 2, 3}
	// 声明多元素的数组，数组的长度是初始值的长度
	arr2 := [...]int {1, 2, 3, 4, 5}
	t.Log(arr[0], arr[1])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T)  {
	// 普通使用 for 遍历数据
	arr3 := [...]int{1, 2, 3, 4, 5}
	len := len(arr3)
	for i := 0;i < len ;i++  {
		t.Log(arr3[i])
		/**
			array_test.go:21: 1
		    array_test.go:21: 2
		    array_test.go:21: 3
		    array_test.go:21: 4
		    array_test.go:21: 5
		 */
	}
	// 类似 foreach 的用法，使用 range 关键字
	// 其中 ixd 为数组的下标，可以使用 _ 占位符进行省略
	for ixd, val := range arr3 {
		t.Log(ixd, val)
		/**
			array_test.go:24: 0 1
		    array_test.go:24: 1 2
		    array_test.go:24: 2 3
		    array_test.go:24: 3 4
		    array_test.go:24: 4 5
		 */
	}
}

func TestMultiArrayTravel(t *testing.T)  {
	// 多维数组的声明
	// 数组中有 3 个 数组元素，每个数组元素中有 2 个元素
	// 第一个[]是 包含的数组个数，第二个是 每个数组中元素的个数
	arr4 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for _, v := range arr4 {
		//t.Log(i, v)
		/**
		 	array_test.go:47: 0 [1 2]
		    array_test.go:47: 1 [3 4]
		    array_test.go:47: 2 [5 6]
		 */
		for vi, vv := range v {
			t.Log(vi, vv)
			/**
				array_test.go:57: 0 1
			    array_test.go:57: 1 2
			    array_test.go:57: 0 3
			    array_test.go:57: 1 4
			    array_test.go:57: 0 5
			    array_test.go:57: 1 6
			 */
		}
	}
}

func TestSubArray(t *testing.T)  {
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	t.Log(arr5[1:2]) // [2]
	t.Log(arr5[1:3]) // [2 3]
	t.Log(arr5[:]) // [1 2 3 4 5 6 7]
	t.Log(arr5[:5]) // [1 2 3 4 5]
	t.Log(arr5[2:]) // [3 4 5 6 7]
}