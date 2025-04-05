package main

import "fmt"

type (
	Camera interface {
		TakePicture()
	}

	DslrCamera struct {
		lens             string
		isHaveViewFinder bool
		maxSdCardSize    int
	}

	CompactCamera struct {
		weight          int
		isCanSwitchLens bool
	}
)

func (d DslrCamera) TakePicture() {
	fmt.Println("Taking picture with DSLR camera with lens:", d.lens)
	fmt.Println("Max SD card size:", d.maxSdCardSize)
	fmt.Println("Has viewfinder:", d.isHaveViewFinder)
}

func (c CompactCamera) TakePicture() {
	fmt.Println("Taking picture with compact camera with weight:", c.weight)
	fmt.Println("Can switch lens:", c.isCanSwitchLens)
}

func main() {
	// Creating instances of dslCamera and compactCamera
	dslr := DslrCamera{
		lens:             "50mm f/1.8",
		isHaveViewFinder: true,
		maxSdCardSize:    128,
	}

	compact := CompactCamera{
		weight:          300,
		isCanSwitchLens: false,
	}

	// Calling the TakePicture method for each camera type
	dslr.TakePicture()
	compact.TakePicture()
}
