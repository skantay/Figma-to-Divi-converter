package figma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchFigmaDesign() ([]byte, error) {
	url := "https://api.figma.com/v1/files/" + FigmaFileID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Figma-Token", FigmaAPIToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Figma API request failed with status code: %d", resp.StatusCode)
	}

	byteVal, _ := io.ReadAll(resp.Body)

	return byteVal, nil
}

func PrettyPrintJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON.Bytes()), nil
}
