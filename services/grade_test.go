package services_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-unit-integration-test/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "success grade A", score: 80, expected: "A"},
		{name: "success grade B", score: 70, expected: "B"},
		{name: "success grade C", score: 60, expected: "C"},
		{name: "success grade D", score: 50, expected: "D"},
		{name: "success grade F", score: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			assert.Equal(t, c.expected, grade)

			//if grade != c.expected {
			//	t.Errorf("got %v expected %v", grade, c.expected)
			//}
		})
	}
}

// go test go-unit-integration-test/services -cover -v -bench=. -benchmem
func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

// godoc -http=:8000
// then go to localhost:8000
func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}
