package array_test

import "testing"

func TestArrayTest(t *testing.T) {
	arr1 := [5]int{1, 2, 3}
	arr2 := arr1
	arr1[4] = 88
	arr2[4] = 99
	t.Log(arr1)
	t.Log(arr2)

	s1 := arr1[:]
	s1[4] = 77
	t.Log(arr1)
	t.Log(s1)
	s1 = append(s1, 99)
	s1[4] = 66
	t.Log(arr1)
	t.Log(s1)

	s2 := s1[:len(s1)-1]
	s2[4] = 55
	t.Log(s1)
	t.Log(s2)

	s2 = append(s2, 44)
	t.Log(s1)
	t.Log(s2)

}

func TestSlice(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s2 := s1[3:5]
	s2 = append(s2, 42)
	t.Log(s1)
	t.Log(s2)

	s3 := s1[3:]
	s3 = append(s3, 42)
	t.Log(s1)
	t.Log(s2)
	t.Log(s3)
}
