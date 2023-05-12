package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type Processor struct {
	ProcessorName string
	Cores         int
}
type Memory struct {
	MemoryCapacity int
	MemoryType     string
}

type Computer struct {
	Processor
	Memory
	Price int
}

func main() {
	// Computer Struct Embedding In Golang

	computer1 := Computer{}
	computer1.ProcessorName = "Intel I3"
	computer1.Cores = 4
	computer1.MemoryCapacity = 8
	computer1.MemoryType = "DDR4"
	computer1.Price = 50000

	fmt.Println(computer1)

	computer2 := Computer{
		Processor: Processor{
			ProcessorName: "Intel i5",
			Cores:         8,
		},
		Memory: Memory{
			MemoryCapacity: 16,
			MemoryType:     "DDR6",
		},
		Price: 70000,
	}

	fmt.Println(computer2)
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
