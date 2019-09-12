package main

import (
	"testing"

	"github.com/test/myint"
)

func main() {

}

func TestAdd(t *testing.T) {
	data := []struct {
		title  string
		value  myint.MyInt
		param  int
		should myint.MyInt
	}{
		{"A", 1, 1, 2},
		{"B", 2, 1, 3},
		{"C", 9, 1, 10},
	}
	for _, v := range data {
		mi := v.value
		mi.Add(v.param)
		if mi != v.should {
			t.Error("for", v.title, "got", mi, "should got", v.should)
		}
	}
}
