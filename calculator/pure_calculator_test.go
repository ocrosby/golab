package calculator

import "testing"

var numberset = []struct {
	x      int
	y      int
	result int
}{
	{1, 2, 3},
	{2, 2, 4},
	{3, 3, 6},
}

func TestAdd(t *testing.T) {
	result := Add(1, 3)
	if result != 4 {
		// t.Fail()
		t.Errorf("Expect 1 + 3 == 4, got %d instead", result)
	}
}

func TestAdds(t *testing.T) {
	for _, set := range numberset {
		result := Add(set.x, set.y)

		if result != set.result {
			t.Errorf("Expected %d + %d == %d, got %d instead.", set.x, set.y, set.result, result)
		}
	}
}
