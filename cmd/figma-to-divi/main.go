package main

import (
	"fidi/internal/figma"
	"fmt"
)

func main() {

	figmaData, err := figma.FetchFigmaDesign()

	if err != nil {
		fmt.Println("Error fetching Figma design:", err)
		return
	}
	val, _ := figma.PrettyPrintJSON(figmaData)
	fmt.Println(val)
}
