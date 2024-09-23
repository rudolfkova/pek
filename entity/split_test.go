package entity

import (
	"fmt"
	"image/color"
	"testing"
)

func TestSplit(t *testing.T) {
	o := NewObject(100, 100, 40, 400, color.RGBA{0, 0, 0, 255})
	so, err := o.Split()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Slice len:", len(so))
	fmt.Println("First. Width:", so[0].Width, "Height:", so[0].Height, "X:", so[0].X, "Y:", so[0].Y)
	fmt.Println("Last. Width:", so[len(so)-1].Width, "Height:", so[len(so)-1].Height, "X:", so[len(so)-1].X, "Y:", so[len(so)-1].Y)
	
}
