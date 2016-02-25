package dollar

import "testing"

func TestDollar_Times(t *testing.T) {
  dollar := Dollar{2}
  if dollar.Times(5) != 10 {
    t.Fatalf("Expected dollar.Times(5) to equal 10")
  }
}
