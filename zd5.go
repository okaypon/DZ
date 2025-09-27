package main

import (
	"fmt"
	"time"
)

func main() {
	cache := map[string]string{}

	cache["polzovatel"] = "чувак"
	go func() {
		time.Sleep(2 * time.Second)
		delete(cache, "polzovatel")
		fmt.Println("чувак удален")
	}()

	fmt.Println("есть:", cache["polzovatel"])
	time.Sleep(1 * time.Second)
	fmt.Println("всё ешё есть:", cache["polzovatel"])
	time.Sleep(2 * time.Second)
	fmt.Println("уже нету чувака(:", cache["polzovatel"])
}

/*
func main() {
	cache := map[string]string{}
	timer := time.NewTimer(2 * time.Second)

	cache["polzovatel"] = "чувак"

	go func() {
		<-timer.C
		delete(cache, "polzovatel")
		fmt.Println("чувак удален")
	}()

	fmt.Println("есть:", cache["polzovatel"])

	timer.Stop()
	timer = time.NewTimer(2 * time.Second)

	time.Sleep(3 * time.Second)
	fmt.Println("перезагрузка он еще сдесь:", cache["polzovatel"])
}
*/
