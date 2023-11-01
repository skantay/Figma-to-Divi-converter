package converter

import (
	"encoding/json"
	"fmt"
)

type FigmaResponse struct {
	Document Document `json:"document"`
}

type Document struct {
	Children []Component `json:"children"`
}

type Component struct {
	Name     string      `json:"name"`
	Children []Component `json:"children"`
}

// ConvertToFigmaLayout converts Figma JSON data to Divi HTML/CSS.
func ConvertToFigmaLayout(figmaData []byte) (string, error) {
	var figmaResponse FigmaResponse
	err := json.Unmarshal(figmaData, &figmaResponse)
	if err != nil {
		return "", err
	}
	htmlOutput := ""

	
	for i, child := range figmaResponse.Document.Children {
		htmlOutput += fmt.Sprintf("<div class='et_pb_section et_pb_section_%d'>", i)
		htmlOutput += generateHTML(child)
		htmlOutput += "</div>\n"
	}

	cssOutput := ".et_pb_section { background-color: #ffffff; }"
	// Apply CSS styles based on figmaResponse.

	return htmlOutput + "\n\n" + cssOutput, nil
}

func generateHTML(component Component) string {
	// Implement logic to map Figma components to Divi HTML elements.
	// Example mapping (for demonstration purposes, actual implementation may vary):
	i := 0

	fmt.Println(component.Name)

	switch component.Name {
	case "header":
		return "<div id=\"header\" >" + generateChildrenHTML(component.Children) + "</div>"
	case "section":
		return fmt.Sprint("<div class=\"et_pb_row et_pb_row_%d\">"+generateChildrenHTML(component.Children)+"</div>", i)
	case "footer":
		return "<div id=\"footer\" >" + generateChildrenHTML(component.Children) + "</div>"
	default:
		return "NOTHING"
	}
}

func generateChildrenHTML(children []Component) string {
	// Implement logic to generate HTML for child components.
	// Iterate through children and call generateHTML recursively.
	// Example code (for demonstration purposes, actual implementation may vary):
	html := ""
	for _, child := range children {
		html += generateHTML(child)
	}
	return html
}
