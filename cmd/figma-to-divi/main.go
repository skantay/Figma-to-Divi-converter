package main

import (
	converter "fidi/internal/divi"
	"fidi/internal/figma"
	"fmt"
)

func main() {

	figmaData, err := figma.FetchFigmaDesign()

	if err != nil {
		fmt.Println("Error fetching Figma design:", err)
		return
	}
	
	result, err := converter.ConvertToFigmaLayout(figmaData)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
