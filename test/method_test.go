package test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"testing"
)

func TestSum(t *testing.T) {
	var i I
	i = 5
	fmt.Println(i.Sum(7))
}

func floatImageTest() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func exact() {
	f := 1.984
	i := int(f)
	fmt.Println(f, i)
	f = 3.12
	fmt.Println(int(f))
}

func floatTest() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func complexTest() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x + y)
	fmt.Println(y - x)
	fmt.Println(x * y)       // "(-5+10i)"
	fmt.Println(real(x * y)) // "-5"
	fmt.Println(imag(x * y)) // "10"
}

func stringTest() {
	var a string = "hello,\n world"
	// c := a[len(a)-1]
	// a[0]='g';
	fmt.Println(a)
	fmt.Println(string(0x4eac))
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	var e string = "189243944564513"
	fmt.Println(comma(e))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Flags Flags
type Flags uint

const (
	// FlagUp is up
	FlagUp Flags = 1 << iota
	// FlagBroadcast supports broadcast access capability
	FlagBroadcast
	// FlagLoopback is a loopback interface
	FlagLoopback
	// FlagPointToPoint belongs to a point-to-point link
	FlagPointToPoint
	// FlagMulticast supports multicast access capability
	FlagMulticast
)

func iotaTest() {
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)
	fmt.Println(FlagMulticast | FlagUp)
}
