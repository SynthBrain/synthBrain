package neurons


import (
	"github.com/g3n/engine/math32"
)

// v, ok := m["Answer"] //проверка на наличие ключа в карте
// fmt.Println("The value:", v, "Present?", ok)

type neuron struct{
	coord math32.Vector3 			// местоположение в пространстве
	coordStimulus []math32.Vector3  // список координат тех кто способствовал активаци в момент времени - для отрисовки связей
	specVector int8				    // специализация вектора в каком направлении отдаёт медиатор (полный радиус или только вверх, или в лево...)
	radius int8 					// максимальная радиус-дистанция куда передаст медиатор, всем кто попал в этот радиус будет передан медиатор

	dendrite []int8                 // массив где индекс означает сторону с которой пришёл заряд, а значение его частоту активаций
	Neuroplasticity []int8          // массив где индекс означает сторону - значение хранит нейропластичность для каждого направления

	/*
		 мапа хранит в ключе строковое представление сформированного нейромедиатора, 
		 а значением будет частота повторений этого нейромедиатора,

		 нейропластичность будет расчитываться от частоты и силы, так же и специализация
	*/
	dendrMapTemp []int8 

	/*
		список наборов нейромедиаторов - специализация, реагирует только на такую комбинацию
		может иметь несколько наборов

		в набор должно входить сторона и нейромедиатор того кто посылает
	*/
	dendrSpec []int8
	

	/*
	мапа что хранит в себе всех соседей в заданном радиусе и векторном направлении и 
	которым будет раздавать свой нейромедиатор
	*/
	axonMap []int8
}

/*
dotGeom := geometry.NewCircle(0, 3)
dotMat := material.NewPhong(math32.NewColor("White"))
dotMesh := graphic.NewPoints(dotGeom, dotMat)
dotMesh.SetPosition(
	float32(rand.Int31n(20)),
	float32(rand.Int31n(20)),
	float32(rand.Int31n(20)))
app.Scene().Add(dotMesh)
*/
func(n *neuron) NewNeuron() *neuron{
	return &neuron{
		coord: *math32.NewVector3(0, 0, 0),
	}
}

func(n *neuron) SetPosition(x, y, z float32){
	n.coord.X = x
	n.coord.Y = y
	n.coord.Z = z
}