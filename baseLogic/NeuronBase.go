package baseLogic


import (
	"github.com/g3n/engine/math32"
)

// v, ok := m["Answer"] //проверка на наличие ключа в карте
// fmt.Println("The value:", v, "Present?", ok)
// uint8 = 0 -> 255

// Neuron - base struct for neuron 
type Neuron struct{
	Coord math32.Vector3 			// местоположение в пространстве
	coordStimulus []math32.Vector3  // список координат тех кто способствовал активаци в момент времени - для отрисовки связей
	specVector int8				    // специализация вектора в каком направлении отдаёт медиатор (полный радиус или только вверх, или в лево...)
	radius int8 					// максимальная радиус-дистанция куда передаст медиатор, всем кто попал в этот радиус будет передан медиатор

	dendrite []uint8                 // массив где индекс означает сторону с которой пришёл заряд, а значение его частоту активаций
	Neuroplasticity []uint8         // массив где индекс означает сторону - значение хранит нейропластичность для каждого направления

	/*
		 мапа хранит в ключе строковое представление сформированного нейромедиатора, 
		 а значением будет частота повторений этого нейромедиатора,

		 нейропластичность будет расчитываться от частоты и силы, так же и специализация
	*/
	dendrMapTemp []uint8

	/*
		список наборов нейромедиаторов - специализация, реагирует только на такую комбинацию
		может иметь несколько наборов

		в набор должно входить сторона и нейромедиатор того кто посылает
	*/
	dendrSpec []uint8
	

	/*
	мапа что хранит в себе всех соседей в заданном радиусе и векторном направлении и 
	которым будет раздавать свой нейромедиатор
	*/
	axonMap []uint8
}


// NewNeuron create new obj
func NewNeuron() *Neuron{
	return &Neuron{
		Coord: 			    *math32.NewVector3(0, 0, 0),
		coordStimulus:    	make([]math32.Vector3, 0),
		specVector: 	  	0,
		radius:			   	0,
		dendrite: 		  	make([]uint8, 0),
		Neuroplasticity:   	make([]uint8, 0),
		dendrMapTemp: 	  	make([]uint8, 0),
		dendrSpec: 		  	make([]uint8, 0),
		axonMap: 		  	make([]uint8, 0),
	}
}


// SetPosition add new coord
func(n *Neuron) SetPosition(x, y, z float32){
	n.Coord.X = x
	n.Coord.Y = y
	n.Coord.Z = z
}