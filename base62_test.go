package base62

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	testData := []struct {
		n        int64
		expected string
	}{
		{0, "0"},
		{10, "a"},
		{630, "aa"},
		{2222821365901088, "abc123EFG"},
		{3781504209452600, "hjNv8tS3K"},
	}
	for _, testCase := range testData {
		r := Encode(testCase.n)
		if r != testCase.expected {
			t.Fatalf("encode expected '%v', but got '%v'", testCase.expected, r)
		}
	}
}

func ExampleEncode() {
	fmt.Println(Encode(99))
	// Output: 1B
}

func TestDecode(t *testing.T) {
	testData := []struct {
		key      string
		expected int64
	}{
		{"0", 0},
		{"a", 10},
		{"aa", 630},
		{"abc123EFG", 2222821365901088},
		{"hjNv8tS3K", 3781504209452600},
	}
	for _, testCase := range testData {
		n, err := Decode(testCase.key)
		if n != testCase.expected {
			t.Fatalf("expected %v, but got %v", testCase.expected, n)
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

func ExampleDecode() {
	i, err := Decode("1B")
	fmt.Println(i, err)
	// Output: 99 <nil>
}

var result interface{}

func BenchmarkEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = Encode(int64(123456789098765432))
	}
}

func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _ = Decode("abcdefghilljkml1234567890")
	}
}
