package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (pc pc) ping() {
	fmt.Println(pc.brand, "pong")
}

func (pc *pc) duplicateRam() {
	pc.ram = pc.ram * 2
}

func main() {
	a := 50
	b := &a
	fmt.Println(*b)
	fmt.Println(b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "Test"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)
}
