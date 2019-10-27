package baseLogic

type Logic struct{
	VisionChan chan [][]byte // include from baselevel to button-WebCam

}

func InitLogic() *Logic{
	return &Logic{
		VisionChan: make(chan [][]byte),
	}
}

func(l *Logic) Update(){

}