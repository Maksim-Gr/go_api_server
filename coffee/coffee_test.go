package coffee

import "testing"

func init() {
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			{"Latte", 2.5},
			{"Flat White", 2},
			{"Cappuccino", 2.25},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	type testCase struct {
		coffeeType string
		want       bool
	}

	cases := []testCase{
		{"lat", false},
		{"Latte", true},
		{"", false},
		{"cappuccino", false},
	}

	for _, tc := range cases {
		got := IsCoffeeAvailable(tc.coffeeType)
		if tc.want != got {
			t.Errorf("Expected '%v', but got '%v'", tc.want, got)
		}
	}
}
