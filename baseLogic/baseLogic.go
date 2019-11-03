package baseLogic

import "github.com/g3n/engine/math32"

type Logic struct {
	VisionChan chan *[][]float32
	data [][]float32
	mapData map[math32.Vector3]float32
}

func InitLogic() *Logic {
	return &Logic{
		VisionChan: make(chan *[][]float32, 1),
	}
}

func (l *Logic) Update() {
	l.getDataFromChan()
}

func (l *Logic) getDataFromChan () {
	select {
	case dataKey := <-l.VisionChan:
		l.data = *dataKey
		count := 0
		//vision.ImgToDataSlice(data)
		l.mapData = make(map[math32.Vector3]float32, len(l.data)*len(l.data[0]))
		tempPosition := *math32.NewVector3(0, 0, 0)
		for i := 0; i < len(l.data); i++ {
			for j := 0; j < len(l.data[i]); j++ {
				//fmt.Println("Start 2 ", j)
				//fmt.Print(data[i][j], " ")
				tempPosition.Set(float32(i), float32(j), 0)
				l.mapData[tempPosition] = l.data[i][j]
				//coords[count].Set(float32(i), float32(j), 0) //data[i][j]
				count++
			}
		}
		//fmt.Println(count)
		//app.Scene().RemoveAll(true) // NOTBAD
		//level.make3DLayer(0, count, coords, app)
	default:
		// залишаєм щоб не стопати фрейм рейт
	}
}

// отдавать данные для рисования сцени
func (l *Logic) GetData() *map[math32.Vector3]float32{
	return &l.mapData
}
