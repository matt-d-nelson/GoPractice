package main

import (
	"fmt"

	"github.com/matt-d-nelson/GoPractice/bankHeist"
	comicMischief "github.com/matt-d-nelson/GoPractice/comicMischief"
	"github.com/matt-d-nelson/GoPractice/practice"
)

func main() {
	practice.HelloWorld()
	fmt.Println(practice.Solution("mynameismatt"))
	comicMischief.PrintComics()
	bankHeist.RunHeist()
	interstellarTravel.FlyToVenus()
}
