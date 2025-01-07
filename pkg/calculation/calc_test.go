package calculation

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"3 + 5", 8, false},
		{"10 - 2 * 3", 4, false},
		{"(1 + 2) * (3 + 4)", 21, false},
		{"8 / 2 + 3", 7, false},
		{"10 / (2 + 3)", 2, false},
		{"10 / 0", 0, true}, // Деление на ноль
		{"2 * (3 + (4 - 1))", 12, false},
		{"", 0, true}, // Пустое выражение
		{"3 + a", 0, true}, // Неверный символ
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		
		if test.expectErr {
			if err == nil {
				t.Errorf("expected an error for expression %q but got none", test.expression)
			}
			continue
		}

		if err != nil {
			t.Errorf("unexpected error for expression %q: %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("for expression %q expected %v but got %v", test.expression, test.expected, result)
		}
	}
}