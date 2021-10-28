package pc

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func New(ram, disk int, brand string) *pc {
	return &pc{ram, disk, brand}
}
func (pc pc) Ping() {
	fmt.Println(pc.brand, " Pong")
}

func (myPc *pc) DuplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc *pc) SetRam(ram int) {
	myPc.ram = ram
}
