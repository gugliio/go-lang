package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (pc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de Ram, %d GB de disco y es una %s", pc.ram, pc.disk, pc.brand)
}

func main() {
	myPc := pc{16, "Asus", 100}
	fmt.Println(myPc)
}
