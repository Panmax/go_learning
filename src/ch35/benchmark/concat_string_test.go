package benchmark_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	elems := []string{"1", "2", "3", "4", "5", "6"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal(t, "123456", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5", "6"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal(t, "123456", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
