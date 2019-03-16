package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThan100Error = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 2 || n > 100 {
	//	return nil, errors.New("n should be in [2, 100]")
	//}

	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThan100Error
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1000); err != nil {
		t.Error(err)
		if err == LessThanTwoError {
			fmt.Println("LessThanTwoError")
		} else if err == LargerThan100Error {
			fmt.Println("LargerThan100Error")
		}
	} else {
		t.Log(v)
	}
}
