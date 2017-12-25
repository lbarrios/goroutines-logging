package main

import(
	"github.com/lbarrios/logging/logger"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var log logger.Logger
	log.Init()

	log.Println("Testing 1")

	wg.Add(1)
	go func(log *logger.Logger, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("Testing 2")
	}(&log, &wg)

	output(&log, 3)

	wg.Add(1)
	go outputRoutine(&log, &wg, 4)

	wg.Wait()
}

func output(log *logger.Logger, d int){
	log.Println("Holi", d)
}

func outputRoutine(log *logger.Logger, wg *sync.WaitGroup, d int){
	defer wg.Done()
	log.Println("Holi", d)
}
