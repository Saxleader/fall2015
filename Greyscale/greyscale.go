package main

import (
	"image"
	"image/jpeg"
	"os"
	"fmt"
//	"image/color"
)

func main() {
	f1, err := os.Open("first.jpg")
	if err != nil {
		fmt.Printf("Error opening file: %v",err)
		return
	}
	defer f1.Close()

	i, err := jpeg.Decode(f1)
	if err != nil {
		fmt.Printf("Error decoding image: %v",err)
		return
	}

	i1, ok := i.(*image.YCbCr)
	if !ok {
		fmt.Println("Image 1 is not YCbCr.")
		return
	}

	b1 := i1.Bounds()

	f2, err := os.Open("second.jpg")
	if err != nil {
		fmt.Printf("Error opening file: %v",err)
		return
	}
	defer f2.Close()

	i, err = jpeg.Decode(f2)
	if err != nil {
		fmt.Printf("Error decoding image: %v",err)
		return
	}

	i2, ok := i.(*image.YCbCr)
	if !ok {
		fmt.Println("Image 2 is not YCbCr.")
		return
	}

	b2 := i2.Bounds()

	if b2.Dx()*b2.Dy() > b1.Dx()*b1.Dy() {
		b1,b2 = b2,b1
		i1,i2 = i2,i1
	}

	o, err := os.Create("output.jpg")
	if err != nil {
		fmt.Printf("Error creating file: %v",err)
		return
	}
	defer o.Close()

	iout := image.NewYCbCr(b1,i1.SubsampleRatio)
//	iout := image.NewRGBA(b1)

//	multi := 1

//	if iout.SubsampleRatio == image.YCbCrSubsampleRatio420{
//		multi = 4
//	}

//	colors := make([]color.RGBA,b1.Dx()*b1.Dy())

//	var ycbcrcolor color.YCbCr
//	var r,g,b,a uint32
//
//	for n := b1.Min.Y; n < b1.Max.Y; n++  {
//		for m := b1.Min.X; m < b1.Max.X; m++ {
//			r,g,b,a = i1.At(m,n).RGBA()
//			iout.Pix[((n-b1.Min.Y)*b1.Dx()+m-b1.Min.X)*4] = uint8(r)
//			iout.Pix[((n-b1.Min.Y)*b1.Dx()+m-b1.Min.X)*4+1] = uint8(g)
//			iout.Pix[((n-b1.Min.Y)*b1.Dx()+m-b1.Min.X)*4+2] = uint8(b)
//			iout.Pix[((n-b1.Min.Y)*b1.Dx()+m-b1.Min.X)*4+3] = uint8(a)
//		}
//	}
//	fmt.Println(r)
//	fmt.Println(g)
//	fmt.Println(b)
//	fmt.Println(a)
//	fmt.Println(uint(r))
//	fmt.Println(uint(g))
//	fmt.Println(uint(b))
//	fmt.Println(uint(a))
	//	dx := b1.Dx()
//	dy := b1.Dy()
//
//	if b1.Dx() > b2.Dx() {
//		dx = b2.Dx()
//	}
//	if b1.Dy() > b2.Dy() {
//		dy = b2.Dy()
//	}
//
//	var r1, g1, bl1, a1, r2, g2, bl2, a2 uint32
//
//	for n := 0; n < dy ; n++  {
//		for m := 0; m < dx; m++ {
//			r1, g1, bl1, a1 = i1.At(b1.Min.X+b1.Dx()/2-dx/2+m,b1.Min.Y+b1.Dy()/2-dy/2+n).RGBA()
//			r2, g2, bl2, a2 = i2.At(b2.Min.X+b2.Dx()/2-dx/2+m,b2.Min.Y+b2.Dy()/2-dy/2+n).RGBA()
//			r,g,b,a = mixpixel(r1, g1, bl1, a1, r2, g2, bl2, a2)
//			iout.Pix[(b1.Dy()/2-dy/2+n)*b1.Dy()+(b1.Dx()/2-dx/2+m)*4] = uint8(r)
//			iout.Pix[(b1.Dy()/2-dy/2+n)*b1.Dy()+(b1.Dx()/2-dx/2+m)*4+1] = uint8(g)
//			iout.Pix[(b1.Dy()/2-dy/2+n)*b1.Dy()+(b1.Dx()/2-dx/2+m)*4+2] = uint8(b)
//			iout.Pix[(b1.Dy()/2-dy/2+n)*b1.Dy()+(b1.Dx()/2-dx/2+m)*4+3] = uint8(a)
//		}
//	}

//	for n := 0; n < b1.Dy(); n++  {
//		for m := 0; m < b1.Dx(); m++ {
//			iout.Y[n*b1.Dx()+m],iout.Cb[(n*b1.Dx()+m)/4],iout.Cr[(n*b1.Dx()+m)/4] = color.RGBToYCbCr(colors[n*b1.Dx()+m].R,colors[n*b1.Dx()+m].G,colors[n*b1.Dx()+m].B)
//		}
//	}

//	for n := 0; n < b1.Dy(); n++  {
//		for m := 0; m < b1.Dx(); m++ {
//			iout.Y[n*b1.Dx()+m] = i1.Y[n*b1.Dx()+m]
//			iout.Cb[(n*b1.Dx()+m)/multi] = i1.Cb[(n*b1.Dx()+m)/multi]
//			iout.Cr[(n*b1.Dx()+m)/multi] = i1.Cr[(n*b1.Dx()+m)/multi]
//		}
//	}


	iout.Y,iout.Cb,iout.Cr = i1.Y,i1.Cb,i1.Cr


//	for n := 0; n < len(i); n++ {
//		iout.Y[len(iout.Y)/2+n],iout.Cb[len(iout.Cb)/2+n],iout.Cr[len(iout.Cr)/2+n] = iout.Y[len(iout.Y)/2+n]+10,iout.Cb[len(iout.Cb)/2+n]+10,iout.Cr[len(iout.Cr)/2+n]+70
//	}

	err = jpeg.Encode(o,iout,nil)
	if err != nil {
		fmt.Printf("Error encoding image: %v",err)
		return
	}
	fmt.Println(b1.Dx())
	fmt.Println(b1.Dy())
	fmt.Println(b2.Dx())
	fmt.Println(b2.Dy())
	fmt.Println(len(i1.Y))
	fmt.Println(len(i1.Cb))
	fmt.Println(len(i1.Cr))
	fmt.Println(i1.SubsampleRatio)
	fmt.Println("Success")
}

func mixpixel(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
	return r2,g2,b2,a2
}