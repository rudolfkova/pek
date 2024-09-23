package Line

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	v1 := NewVec(0, 0, 12, 12)
	v2 := NewVec(15, 15, 15, 30)
	xc, yc, err := v1.Intersect(v2)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("X:", xc, "Y:", yc)
	}
}
