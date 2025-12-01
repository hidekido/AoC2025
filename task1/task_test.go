import (
    "testing"
    "regexp"
)

func TestSolve(t *testing.T) {
    raw := "L68
			L30
			R48
			L5
			R60
			L55
			L1
			L99
			R14
			L82"
	test := make([]string, 10)
	for _, val := range strings.Split(raw, '\n') {
		test = append(test, string.Trim(val))
	}
	fmt.Println(test)
}