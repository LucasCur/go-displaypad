package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ImageRequest struct {
	ImageData string `json:"image_data"`
}

func main() {
	// Replace this with your actual code to handle HTTP requests
	http.HandleFunc("/send_image", handleSendImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSendImage(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var imageRequest ImageRequest
	err := json.NewDecoder(r.Body).Decode(&imageRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Get the image data from the request
	imageData := imageRequest.ImageData

	// Convert the image data from base64 to bytes
	imageBytes, err := decodeBase64(imageData)
	if err != nil {
		http.Error(w, "Failed to decode image data", http.StatusBadRequest)
		return
	}

	// Replace this code with your actual code to send the image to the display pad
	err = sendImageToDisplay(imageBytes)
	if err != nil {
		http.Error(w, "Failed to send image to display", http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprint(w, "Image sent successfully")
}

func decodeBase64(data string) ([]byte, error) {
	// Remove the data URL prefix if present
	data = strings.TrimPrefix(data, "data:image/png;base64,")

	// Decode the base64 data
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}

func sendImageToDisplay(imageBytes []byte) error {
	// Replace this code with your actual code to send the image bytes to the display pad
	// Example:
	// displayPad.SendImage(imageBytes)
	// Replace "displayPad" with your display pad library or implementation

	return nil
}
