package vision

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"log"
	"os"
)

// StartWebCam
func StartWebCam(chFlag chan bool) {
	// set to use a video capture device 0
	deviceID := 0
	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println(err)
		webcam.Close()
		return
	}
	defer webcam.Close()

	// open display window
	//window := gocv.NewWindow("WebCam")
	//defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	imgVision, err := img.ToImage()
	if err != nil {
		fmt.Println("ImgMat not convert")
	}
	//imgVision.Bounds().

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}
		//window.IMShow(img)
		imgVision, _ = img.ToImage()

		fmt.Println(imgVision.Bounds().Size())

		//write jpg file
		//gocv.IMWrite("C:\Users\synth\go\src\github\SynthBrain\synthBrain\assets\webCam.jpg", img)

		//window.WaitKey(1)
		if len(chFlag) > 0 {
			if ok := <-chFlag; ok {
				fmt.Println("Thread WebCam Close")
				break
			}
		}
	}
}

func ReadImg(dataDir string) {
	dataDir = dataDir + "/assets/0.jpg"

	reader, err := os.Open(dataDir)
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	Print2DSlice(ImgToDataSlice(m))
	fmt.Println()
}

func ImgToDataSlice(img image.Image) [][]byte {
	bounds := img.Bounds()

	data := make([][]byte, bounds.Size().X) // create 1D slice size columns
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		data[i] = make([]byte, bounds.Size().Y) // create 2D slice size rows
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			//data[i][j] = byte((r + g + b) / 3)
			data[i][j] = byte(((b >> 8) + (g >> 8) + (r >> 8)) / 3)
		}
	}
	return data
}

func Print2DSlice(data [][]byte) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[j][i], " ")
		}
		fmt.Println()
	}
}
