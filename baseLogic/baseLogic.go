package baseLogic

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
	VisionChan chan *[][]float32
	Data       [][]float32

	flagReady  bool
}

func InitLogic() *Logic {
	return &Logic{
		VisionChan: make(chan *[][]float32, 1),
	}
}

func (l *Logic) Update() {
	l.getDataFromChan()
}

func (l *Logic) getDataFromChan() {
	select {
	case dataKey := <-l.VisionChan:
		l.flagReady = true
		l.Data = *dataKey
	default:
		l.flagReady = false
	}
}

// отдавать данные для рисования сцени
func (l *Logic) GetData() *[][]float32 {
	return &l.Data
}

func (l *Logic) GetReady() bool{
	return l.flagReady
}