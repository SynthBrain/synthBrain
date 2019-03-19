package samples

import "SynthBrainGo/samples/models"

func notOopmain() {
	user := models.NewUser("Synth", "Brain", 7000)
	go user.GetSalary()

	user2 := models.NewUser("Jack", "Power", 4000)
	go user2.GetSalary()
	go user2.GetPrintNameUser()
	//fmt.Scanln()
}
