package main

import (
	"log"
	"os"

	"SpamMasker/internal/service"
)

func main() {

	inputPath := "input.txt" // дефолтное название входного ф-ла
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	outputPath := "output.txt" // дефолтное для выходного
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	prod := service.NewFileProducer(inputPath)
	pres := service.NewFilePresenter(outputPath)
	service := service.NewService(prod, pres)

	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
