package baseLogic

import (
	"math"

	"github.com/g3n/engine/math32"
)

/*
	Доступное поле зрения на нем фокус для которого есть
	1. нейронка которая учиться его двигать
	2. нейронка которая запоминает паттерны
	3. потребность которая заставляет работать первую на основе данных со второй
*/

/*
	для каждого органа свой полюсный пузырь
	такие пузыри пересекаються так как находяться в одном общем пузыре
	на местах пересечений создаються связи
*/

type Logic struct {
	LightCube map[math32.Vector3]*Neuron
	//LightCub map[math32.Vector3]float32
	//VisionChan chan *[][]float32
	//VisionChan chan map[math32.Vector3]float32
	VisionChan chan map[math32.Vector3]math32.Vector3
	SoundChan  chan *[][]float32

	//DataVision       [][]float32
	DataVision       map[math32.Vector3]math32.Vector3
	DataSound        [][]float32

	flagReady  bool
	ActiveVectors map[math32.Vector3]bool
	ColorsLayer map[math32.Vector3]float32
}

func InitLogic() *Logic {
	return &Logic{
		LightCube: make(map[math32.Vector3]*Neuron),
		//VisionChan: make(chan *[][]float32, 1),
		VisionChan: make(chan map[math32.Vector3]math32.Vector3, 1),
		SoundChan:  make(chan *[][]float32, 1),
		ActiveVectors: make(map[math32.Vector3]bool),
		ColorsLayer:  make(map[math32.Vector3]float32),
		//DataVision: make(map[math32.Vector3]float32),
	}
}

func (l *Logic) Update() {
	l.getDataFromVisionChan()
	l.DeactivateNeurons()
	l.UpdLightCube()
}

func (l *Logic) getDataFromVisionChan() {
	select {
	case l.DataVision = <-l.VisionChan:
		//l.DataVision = *dataKey
		l.flagReady = true
	default:
		l.flagReady = false
	}
}

// отдавать данные для рисования сцени
//func (l *Logic) GetDataVision() *[][]float32 {
func (l *Logic) GetDataVision() *map[math32.Vector3]math32.Vector3 {
	return &l.DataVision
}

func (l *Logic) GetReady() bool{
	return l.flagReady
}

func (l *Logic) UpdLightCube(){
	var vector math32.Vector3
	for key, value := range l.DataVision {
		if (key.X >= 305 && key.Z >= 225) && (key.X <= 335 && key.Z <= 255){
			valueTemp := (((value.X + value.Y + value.Z) / 3))
			valueTemp = Round(float64(valueTemp), 2)
			l.ColorsLayer[value] = 1
			// создавать нейрон по общему правилу *****************************************!
			// на дендриты кто то должен повлиять для активации/создания, даже для появления представителя входящего сигнала
			// создать карту где будем хранить представителей цвета
			// на основе местоположение цвета в карте и положения пикселя на сетчатке создавать представителя в LightCube
			vector.Set(key.X, key.Y + valueTemp + 10, key.Z)
			if _, ok := l.LightCube[vector]; ok {

				if valueTemp > 0.2{
					l.LightCube[vector].SetActive(true)
					l.LightCube[vector].Power = valueTemp
					l.ActiveVectors[vector] = true
				} else {
					l.LightCube[vector].SetActive(false)
				}
			} else {
				l.LightCube[vector] = NewNeuron()
				l.LightCube[vector].SetActive(true)
				l.LightCube[vector].Power = valueTemp
				l.ActiveVectors[vector] = true
			}
		}
	}
}

func (l *Logic) DeactivateNeurons(){
	for key,_ := range l.ActiveVectors {
		l.LightCube[key].SetActive(false)
		delete(l.ActiveVectors, key)
	}
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
