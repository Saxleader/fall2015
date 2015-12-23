package main

import (
	"image"
	"image/jpeg"
//	"os"
//	"fmt"
	"io"
)

func greyscale(f io.Reader) (*image.YCbCr, error) {
//	fmt.Println("Which document would you like to scan:")
//	var input string
//	fmt.Scanln(&input)

//	f, err := os.Open(input)
//	if err != nil {
//		fmt.Printf("Error opening file: %v",err)
//		return
//	}
//	defer f.Close()

	i, err := jpeg.Decode(f)
	if err != nil {
//		fmt.Printf("Error decoding image: %v",err)
		return nil,err
	}

	im, ok := i.(*image.YCbCr)
	if !ok {
//		fmt.Println("Image not YCbCr.")
		return nil,err
	}

//	b := im.Bounds()

//	o, err := os.Create("output.jpg")
//	if err != nil {
//		fmt.Printf("Error creating file: %v",err)
//		return nil,err
//	}
//	defer o.Close()

	for key, _ := range im.Cb {
		im.Cb[key]=128
	}

	for key, _ := range im.Cr {
		im.Cr[key]=128
	}

//	err = jpeg.Encode(o,im,nil)
//	if err != nil {
//		fmt.Printf("Error encoding image: %v",err)
//		return nil,err
//	}

	return im,nil
//	fmt.Println(b.Dx())
//	fmt.Println(b.Dy())
//	fmt.Println(len(im.Y))
//	fmt.Println(len(im.Cb))
//	fmt.Println(len(im.Cr))
//	fmt.Println(im.SubsampleRatio)
//	fmt.Println("Success")
}