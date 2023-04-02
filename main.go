package main

import (
	"flag"

	"github.com/joho/godotenv"
)

var (
	port = flag.Int("port", 8888, "server port")
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("failed load .env")
	}
}

func main() {

}
