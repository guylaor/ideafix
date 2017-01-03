package main

import "log"

func checkError(e error) {
	if e != nil {
		log.Panicf("Error: %s", e)
	}
}

func checkWarnning(e error) {
	if e != nil {
		log.Printf("Warning: %s \n", e)
	}
}

func main() {
	loadConfigValues()

	//loadYoutube()
	loadFacebook()
}
