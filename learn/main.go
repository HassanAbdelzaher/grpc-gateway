package main

import (
	"log"
	"regexp"
)

func main() {
	pattern := "abc"
	str := "abc/"
	match, err := regexp.MatchString(pattern, str)
	log.Println(match, err)
}
