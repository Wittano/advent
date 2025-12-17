package six

import "testing"

func TestSumHomeworks(t *testing.T) {
	data := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	sum, err := SumHomeworks(data)
	if err != nil {
		t.Fatal(err)
	}

	const exp = 4277556
	if sum != exp {
		t.Errorf("SumHomeworks got %f, want %d", sum, exp)
	}
}
