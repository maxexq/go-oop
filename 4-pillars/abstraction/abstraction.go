package main

type (
	Camera interface {
		takePicture()
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

func takePictureWithCamera(c Camera) {
	c.takePicture()
}

func (d DslrCamera) takePicture() {
	println("Taking picture with DSLR camera with lens:", d.lens)
	println("Max SD card size:", d.maxSdCardSize)
	println("Has viewfinder:", d.isHaveViewFinder)
}

func (c CompactCamera) takePicture() {
	println("Taking picture with compact camera with weight:", c.weight)
	println("Can switch lens:", c.isCanSwitchLens)
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

	// Taking pictures with both cameras
	// The takePictureWithCamera function accepts any type that implements the Camera interface
	takePictureWithCamera(dslr)
	takePictureWithCamera(compact)
}
