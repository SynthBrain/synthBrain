package vision

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"time"

	"github.com/g3n/engine/math32"
	"gocv.io/x/gocv"
)

// type Data struct{
// 	Image image.Image
// 	Slice [][]byte
// }

var dataImage image.Image
var dataSlice [][]float32
var dataMap map[math32.Vector3]math32.Vector3

// StartWebCam
//func StartWebCam(chFlag chan bool, visionChan chan *[][]float32) {
func StartWebCam(chFlag chan bool, visionChan chan map[math32.Vector3]math32.Vector3) {
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

	// init Part of Memory for img Data ************************
	if ok := webcam.Read(&img); !ok {
		//fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		return
	} else {
		dataImage, _ = img.ToImage()
		initMemory(&dataImage)
	}
	//**************************************


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
			//ImgChan <- &dataImage
			gocv.IMWrite("C:/Users/synth/go/src/github.com/SynthBrain/synthBrain/data/webCam.jpg", img)
			//gocv.IMWrite("~/data/webCam.jpg", img)
		default:
			time.Sleep(50 * time.Millisecond)
		}

		//Print2DSlice(*ImgToDataSlice(&dataImage))
		//fmt.Println(imgVision.Bounds().Size())

		//write jpg file
		//gocv.IMWrite("C:/Users/synth/go/src/github.com/SynthBrain/synthBrain/assets/webCam.jpg", img)
		
		//gocv.IMWrite(app.DirData() + "webCam.jpg", img)

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
	//m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//Print2DSlice(*ImgToDataSlice(&m))
	fmt.Println()
}

// ImgToDataSlice convert image to slice
//func ImgToDataSlice(img *image.Image) *[][]float32 {
func ImgToDataSlice(img *image.Image) map[math32.Vector3]math32.Vector3 {
	var vertex math32.Vector3
	var color math32.Vector3
	var tempI float32
	dataMap = make(map[math32.Vector3]math32.Vector3)
	//data := dataSlice
	imgTemp := *img
	bounds := imgTemp.Bounds()
	//temp := 0
	//dataSlice = make([][]float32, bounds.Size().Y) // create 1D slice size columns
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		//dataSlice[y] = make([]float32, bounds.Size().X) // create 2D slice size rows
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//dataSlice[y][x] = byte(((b >> 8) + (g >> 8) + (r >> 8)) / 3)
			r, g, b, _ := imgTemp.At(x, y).RGBA()
			red := float32(r) / 255
			green := float32(g) / 255
			blue := float32(b) / 255

			red = Round(float64(red), 0)
			green = Round(float64(green), 0)
			blue = Round(float64(blue), 0)

			red = Round(float64(red / 255), 2)
			green = Round(float64(green / 255), 2)
			blue = Round(float64(blue / 255), 2)

			//dataSlice[y][x] = ((red + green + blue) / 3) / 255

			if (x % 2.0 == 1) {
				tempI = 0.5
			}
			vertex.Set(
				float32(x),
				-1,
				float32(y) + tempI,
			)
			color.Set(red, green, blue)
			//dataMap[vertex] = ((red + green + blue) / 3) / 255
			dataMap[vertex] = color
			tempI = 0
		}
	}
	//return &dataSlice
	return dataMap
}

func Print2DSlice(data [][]float32) {
	for i := 0; i < len(data)-460; i++ {
		for j := 0; j < len(data[i])-620; j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}

func initMemory(img *image.Image) {
	//dataMap = make(map[math32.Vector3]float32)
	//imgTemp := *img
	// bounds := imgTemp.Bounds()
	// dataSlice = make([][]float32, bounds.Size().Y) // create 1D slice size columns
	// for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
	// 	dataSlice[y] = make([]float32, bounds.Size().X) // create 2D slice size rows
	// }
}

func Round(x float64, prec int) float32 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	result := float32(rounder / pow)

	return result
}
