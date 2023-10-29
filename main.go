package main

import "kp-final/controllers"

func main() {
	e := controllers.Init()
	e.Logger.Fatal(e.Start(":1234"))

}
