// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

/*** Without any added goroutines or changes the program's runtime is aprroximately around 19-20s on my stationary computer.
	The idea, to begin with, is to create one goroutine for each function/file aka 8 goroutines etc.
	So putting a goroutine for each call to createPng, however this also requires  waitgroup and it also requires a way to prevent data race
	on variable n and fn in the main fuctions for-loop.

	Having done this program runs at approximately 14 seconds, 4-5 seconds improvement.

	The idea is then to split the image into sub-images and let goroutines work on each subimage before returning the merged image.
	The improvement after having each of the 8 "function goroutines" create 4 goroutine for 4 subimages, the program ran at
	approximately 6 seconds, 8 second improvement.

	The idea was to create function that create x amount of subimages with handling nxn pixels, but I didn't get that far so  settled at a static 4, hope that's enough.

	Forcing the program to run on one processing unit increases the runtime to 19+ sec
----------------------------------------------------
	How much faster is your parallell version?
	Approximately 13-14 seconds faster

	How many CPUs does you program use?
	All/4

***/

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
	//procs := runtime.NumCPU()
	//runtime.GOMAXPROCS(procs)
	before := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(8)
	for n, fn := range Funcs {
		i := n
		in := fn
		go func() {
			err := CreatePng("picture-"+strconv.Itoa(i)+".png", in, 1024)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("time:", time.Now().Sub(before))
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	return
}

/* MODIFIED VERSION OF Julia(f ComplexFunc, n int) image.Image */
// Julia returns an image of size n x n of the Julia set for f.

func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	s := float64(n / 4)
	wg2 := new(sync.WaitGroup) //wait group that will tell when all computations for all the subimages are done
	wg2.Add(4)
	//splits image into 4 sub-images (of size 512x512) that are computed in seperate goroutines befeore final image is returned.
	for k := 0; k < 2; k++ {
		sk := k //copy for data race. This variable had to be sent as a parameter to go functions and not accessed globally as that had an effect on runtime.
		go func(k int) {
			for i := bounds.Min.X; i < bounds.Min.X+512; i++ {
				for j := bounds.Min.Y + k*512; j < bounds.Min.Y+(k+1)*512; j++ {
					n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
					r := uint8(0)
					g := uint8(0)
					b := uint8(n % 32 * 8)
					img.Set(i, j, color.RGBA{r, g, b, 255})
				}
			}
			wg2.Done()
		}(sk)

		go func(k int) {
			for i := 0; i < bounds.Max.X; i++ {
				for j := bounds.Min.Y + k*512; j < bounds.Min.Y+(k+1)*512; j++ {
					n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
					r := uint8(0)
					g := uint8(0)
					b := uint8(n % 32 * 8)
					img.Set(i, j, color.RGBA{r, g, b, 255})
				}
			}
			wg2.Done()
		}(sk)
	}
	wg2.Wait()
	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}
