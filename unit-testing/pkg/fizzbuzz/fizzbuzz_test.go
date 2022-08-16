package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type FizzBuzzInp struct {
	Total  int64
	FizzAt int64
	BuzzAt int64
}

var expectedValues = map[FizzBuzzInp][]string{
	{Total: 3, FizzAt: 3, BuzzAt: 5}: {"1", "2", "Fizz"},
	{Total: 5, FizzAt: 3, BuzzAt: 5}: {"1", "2", "Fizz", "4", "Buzz"},
	{Total: 8, FizzAt: 3, BuzzAt: 5}: {"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8"},
}

func TestFizzBuzz__Success(t *testing.T) {
	t.Run("FizzBuzz--success", func(t *testing.T) {
		for inp, out := range expectedValues {
			resp := FizzBuzz(inp.Total, inp.FizzAt, inp.BuzzAt)
			require.Equal(t, resp, out)
		}
	})
}

func TestFizzBuzz__Fail_Negative(t *testing.T) {
	t.Run("FizzBuzz--fail:negative-total", func(t *testing.T) {
		resp := FizzBuzz(-1, 3, 5)
		require.Nil(t, resp)
	})
	t.Run("FizzBuzz--fail:negative-fizzat", func(t *testing.T) {
		resp := FizzBuzz(5, -3, 5)
		require.Nil(t, resp)
	})
	t.Run("FizzBuzz--fail:negative-buzzat", func(t *testing.T) {
		resp := FizzBuzz(5, 3, -5)
		require.Nil(t, resp)
	})
}

func TestFizzBuzz__Fail_TooLarge(t *testing.T) {
	t.Run("FizzBuzz--fail:negative-total", func(t *testing.T) {
		resp := FizzBuzz(1000000000000000000, 3, 5)
		require.Nil(t, resp)
	})
}
