package main

import(
	"github.com/lbarrios/logging/logger"
)

func main() {
	log := new(logger.Logger)
	log.Init()

	log.Println("Testing 1")
	go log.Println("Testing 2")

	output(log, 3)
	go output(log, 4)
}

func output(log *logger.Logger, d int){
	log.Println("Holi", d)
}
