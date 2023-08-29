package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Set your Confluence credentials and other details
	username := "EMAIL"
	password := "API_TOKEN"
	filePath := "report.csv"
	comment := "File attached via REST API"
	attachmentURL := "CONF_BASE_URL"

	// Open the file to be uploaded
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to store the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add the file and comment fields to the multipart form
	filePart, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(filePart, file)
	if err != nil {
		fmt.Println("Error copying file data:", err)
		return
	}
	writer.WriteField("comment", comment)

	// Close the multipart writer
	writer.Close()

	// Create a new HTTP request
	request, err := http.NewRequest("POST", attachmentURL, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set basic authentication header
	request.SetBasicAuth(username, password)

	// Set the appropriate content type for the request
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("X-Atlassian-Token", "nocheck")

	// Perform the HTTP request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Display the response status
	fmt.Println("Response Status:", response.Status)
}
