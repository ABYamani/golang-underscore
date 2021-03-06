package underscore

import (
	"testing"
	"time"
)

func TestEach(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	Each(arr, func(r TestModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func TestChain_Each(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	Chain(arr).Each(func(r TestModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func TestChain_Parallel_Each(t *testing.T) {
	arr := []int{1, 2, 3}
	beginUnix := time.Now().Unix()
	Chain(arr).AsParallel().Each(func(n, i int) {
		time.Sleep(time.Second)
	}).Value()
	endUnix := time.Now().Unix()
	if int(endUnix-beginUnix) > len(arr) {
		t.Error("wrong")
	}
}
