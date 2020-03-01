package main

import (
	"log"
	"time"
)

/*
@Time : 2020/3/1
@Author : hejun
*/

func main() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			log.Println("xxxxxxxxxx")
		}
	}
}
