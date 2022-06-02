package split

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	got := Split("bababef", "b")
// 	want := []string{"", "a", "a", "ef"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("want %v, but got %v", want, got)
// 	}
// }

//使用slice进行的测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := []testCase{
		{"bababef", "b", []string{"", "a", "a", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "a"}},
	}
	for _, tr := range testGroup {
		got := Split(tr.str, tr.sep)
		if !reflect.DeepEqual(got, tr.want) {
			t.Fatalf("faild")
		}
	}
}

//子测试，使用map来进行测试
func TestSplit1(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": {"bababef", "b", []string{"", "a", "a", "ef"}},
		"case2": {"a:b:c", ":", []string{"a", "b", "a"}},
	}
	for name, tr := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tr.str, tr.sep)
			if !reflect.DeepEqual(got, tr.want) {
				t.Fatalf("faild")
			}
		})
	}
}
