package vision

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

var OnOff = false

/*
Тут создадим обьект
инициализируем для начала в мейне(потом в спец пакете)
тут будет создан канал в который будем ложить двумерный массив изображения
а в мейне читать его - присваивать массиву представителей которые будем выставлять с координатами
*/

//func StartWebCam(ch chan<- *[640][480]byte) {
func StartWebCam() {
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

	//var massVis [640][480]byte

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
		//fmt.Println(imgVision)

		//massVis = massiveImg(imgVision)
		//massVis2 := &massVis
		//ch <- massVis2

		//write jpg file
		//gocv.IMWrite("file.jpg", img)

		//window.WaitKey(1)
		if OnOff == true {
			//fmt.Println("Stop WebCam")
			break
		}
	}
}

func massiveImg(img image.Image) [640][480]byte {
	bounds := img.Bounds()
	//fmt.Println(bounds.Max.X, bounds.Max.Y)
	var data [640][480]byte
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			data[i][j] = byte((r + g + b) / 3)
		}
	}

	//for i:= 0; i < len(data[i]); i++{
	//	for j:= 0; j < len(data[j]); j++{
	//		fmt.Print(data[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	return data
}

func myVision(img image.Image) []byte {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	data := make([]byte, 0, x*y*3)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			data = append(data, byte(b>>8), byte(g>>8), byte(r>>8))
		}
	}
	return data
}
