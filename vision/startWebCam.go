package vision

import (
	"fmt"
	"image"
	"log"
	"os"
	"time"

	"gocv.io/x/gocv"
)

// type Data struct{
// 	Image image.Image
// 	Slice [][]byte
// }

var dataImage image.Image
var dataSlice [][]byte

// StartWebCam
func StartWebCam(chFlag chan bool, visionChan chan *[][]byte) {
	//data := new(Data)

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
	//(sizeX := 3840, sizeY:= 2160)

	// open display window
	//window := gocv.NewWindow("WebCam")
	//defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	//imgVision, err := img.ToImage()
	dataImage, err = img.ToImage()
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
		dataImage, _ = img.ToImage()
		select {
		case visionChan <- ImgToDataSlice(&dataImage):
		default:
			time.Sleep(50 * time.Millisecond)
		}

		//Print2DSlice(ImgToDataSlice(imgVision))
		//fmt.Println(imgVision.Bounds().Size())

		//write jpg file
		//gocv.IMWrite("C:/Users/synth/go/src/github.com/SynthBrain/synthBrain/assets/webCam.jpg", img)

		//window.WaitKey(1)
		if len(chFlag) > 0 {
			if ok := <-chFlag; ok {
				fmt.Println("Thread WebCam Close")
				break
			}
		}
	}
}

func ReadImg(dataDir string, name string) {
	dataDir = dataDir + name

	reader, err := os.Open(dataDir)
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	Print2DSlice(*ImgToDataSlice(&m))
	fmt.Println()
}

// ImgToDataSlice convert image to slice
func ImgToDataSlice(img *image.Image) *[][]byte {
	//data := dataSlice
	imgTemp := *img
	bounds := imgTemp.Bounds()

	dataSlice = make([][]byte, bounds.Size().Y) // create 1D slice size columns
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		dataSlice[y] = make([]byte, bounds.Size().X) // create 2D slice size rows
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := imgTemp.At(x, y).RGBA()
			dataSlice[y][x] = byte(((b >> 8) + (g >> 8) + (r >> 8)) / 3)
		}
	}
	return &dataSlice
}

func Print2DSlice(data [][]byte) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
