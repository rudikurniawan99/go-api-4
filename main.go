package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rudikurniawan99/go-api-4/src"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not loader properly")
	}
	log.Println("read .env from file")

	wg := sync.WaitGroup{}
	once := sync.Once{}

	srv := src.InitServer()
	wg.Add(1)

	go func() {
		defer wg.Done()
		once.Do(func() {
			srv.Run()
		})
	}()

	wg.Wait()
}
