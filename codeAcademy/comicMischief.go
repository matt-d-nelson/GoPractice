package comicMischief

import (
	"fmt"
)

func PrintComics() {

	var publisher string
	var writer string
	var artist string
	var title string

	var year int
	var pageNumber int

	var grade float32

	publisher = "DizzyBooks Publishing Inc."
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "published by", publisher, "written by", writer, "drawn by", artist, "published in", year, "page number:", pageNumber, "grade:", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "published by", publisher, "written by", writer, "drawn by", artist, "published in", year, "page number:", pageNumber, "grade:", grade)
}
