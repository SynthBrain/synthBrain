package samples

import "fmt"

func notTesmain(){
	fmt.Println("Start")
	process(7)

	go process(7)

	go func() {
		fmt.Println("Anonimus")
	}()

	for i:=0; i< 10; i++{
		go process(i)
	}

	// ждём пока отработаёт горутины
	fmt.Scanln() // ввод строки в консоль
}

func process(i int)  {
	fmt.Println("Штурмовик номер: ",i)
}