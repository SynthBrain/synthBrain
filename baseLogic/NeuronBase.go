package baseLogic


import (
	"github.com/g3n/engine/math32"
)

// v, ok := m["Answer"] //проверка на наличие ключа в карте
// fmt.Println("The value:", v, "Present?", ok)
// uint8 = 0 -> 255

// Neuron - base struct for neuron 
type Neuron struct{
	Power float32
	position math32.Vector3
	active bool

	activationLevel uint8
	axons map[math32.Vector3]uint8
	dendrites map[math32.Vector3]uint8
	dendritesPlasticity map[math32.Vector3]uint8






	PowerActivation float32			// Сила Активации для InputNeurons
	Coord math32.Vector3 			// местоположение в пространстве
	CoordStimulus []math32.Vector3  // список координат тех кто способствовал активаци в момент времени - для отрисовки связей
	specVector int8				    // специализация вектора в каком направлении отдаёт медиатор (полный радиус или только вверх, или в лево...)
	radius int8 					// максимальная радиус-дистанция куда передаст медиатор, всем кто попал в этот радиус будет передан медиатор

	dendrite [][]uint8              // массив где индекс означает сторону с которой пришёл заряд, а значение его нейромедиаторный код
	dendritePeriodicity []uint8     // массив где индекс означает сторону с которой пришёл заряд, а значение его частоту активаций
	Neuroplasticity []uint8         // массив где индекс означает сторону - значение хранит нейропластичность для каждого направления

	/*
		 мапа хранит в ключе строковое представление сформированного нейромедиатора, 
		 а значением будет частота повторений этого нейромедиатора,

		 нейропластичность будет расчитываться от частоты и силы, так же и специализация
	*/
	DendrMapTemp []uint8

	/*
		список наборов нейромедиаторов - специализация, реагирует только на такую комбинацию
		может иметь несколько наборов

		в набор должно входить сторона и нейромедиатор того кто посылает
	*/
	DendrSpec []uint8
	

	/*
	мапа что хранит в себе всех соседей в заданном радиусе и векторном направлении и 
	которым будет раздавать свой нейромедиатор
	*/
	AxonMap []uint8
}


// NewNeuron create new obj
func NewNeuron() *Neuron{
	return &Neuron{
		Coord: 			    *math32.NewVector3(0, 0, 0),
		CoordStimulus:    	make([]math32.Vector3, 0),
		//dendrite: 		  	make([]uint8, 0),
		Neuroplasticity:   	make([]uint8, 0),
		DendrMapTemp: 	  	make([]uint8, 0),
		DendrSpec: 		  	make([]uint8, 0),
		AxonMap: 		  	make([]uint8, 0),
	}
}


// SetPosition add new coord
func(n *Neuron) SetPosition(x, y, z float32){
	n.Coord.X = x
	n.Coord.Y = y
	n.Coord.Z = z
}

// SetPosition add new coord
func(n *Neuron) SetPowerActivation(x float32){
	n.PowerActivation = x
}

// SetActive add new coord
func(n *Neuron) SetActive(active bool){
	n.active = active
}

// GetActive add new coord
func(n *Neuron) GetActive() bool{
	return n.active
}


func(n *Neuron) CreatePattern(LightCub *map[math32.Vector3]Neuron, DataVision *[][]float32){
	/*
	tempVision := *DataVision
	var vertex math32.Vector3

	var tempI float32
	*/

	// так как нейрон имеет радиус действия пробегаться только по его маленькому радиусу
	// call for everyone neuron

	//for i := 0; i < len(tempVision); i++ {
	//	for j := 0; j < len(tempVision[0]); j++ {
	//		if (j % 2.0 == 1) {
	//			tempI = 0.5
	//		}
	//		//color := tempVision[i][j]
	//		vertex.Set(
	//			float32(j),
	//			0,
	//			float32(i)+ tempI,
	//		)
	//		tempI = 0
	//	}
	//}
}