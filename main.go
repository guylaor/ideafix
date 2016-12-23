package main

import "log"

func checkError(e error) {
	if e != nil {
		log.Panicf("Error: %s", e)
	}
}

func main() {
	loadConfigValues()

	loadYoutube()
}
