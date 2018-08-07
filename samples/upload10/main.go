package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {

	// Set up the HTTP server:
	serverMUX := http.NewServeMux()
	serverMUX.HandleFunc("/upload", handlerUpload)

	server := &http.Server{}
	server.Addr = ":50000"
	server.Handler = serverMUX
	server.SetKeepAlivesEnabled(true)
	server.ReadTimeout = 60 * 120 * time.Second // 2 hours
	server.WriteTimeout = 16 * time.Second

	// Start the server:
	go func() {
		log.Println("The HTTP web server starts now on http://127.0.0.1:50000")
		if errHTTP := server.ListenAndServe(); errHTTP != nil {
			log.Printf("Was not able to start the HTTP server: %s\n", errHTTP.Error())
			os.Exit(2)
		}
	}()

	//
	// Here, the client code starts. It could be in a separate program.
	//
	time.Sleep(6 * time.Second)

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	if formWriter, formErr := bodyWriter.CreateFormFile("file", "my file.jpg"); formErr != nil {
		log.Printf("Could not create the form: %s\n", formErr.Error())
	} else {

		// Read the source file into memory:
		if fileData, fileErr := ioutil.ReadFile(filepath.Join("./fox.jpg")); fileErr != nil {
			log.Printf("Could not read the source file: %s\n", fileErr.Error())
		} else {
			// Write the data to the form:
			formWriter.Write(fileData)
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Write the form to the server:
	if _, errPOST := http.Post("http://127.0.0.1:50000/upload", contentType, bodyBuffer); errPOST != nil {
		log.Printf("Was not able to send the data: %s", errPOST.Error())
	} else {
		log.Println("Client: File was send.")
	}

	// Wait forever.
	for {
		time.Sleep(1 * time.Second)
	}
}

func handlerUpload(response http.ResponseWriter, request *http.Request) {
	if file, fileHeader, fileError := request.FormFile(`file`); fileError != nil {
		log.Printf("Was not able to access the uploaded file: %s", fileError.Error())
	} else {

		// Close the file afterwards:
		defer file.Close()

		// Get the original filename:
		sourceFilename := fileHeader.Filename

		// Read the entire file into memory:
		if data, dataErr := ioutil.ReadAll(file); dataErr != nil {
			log.Printf("Error while reading file from client: %s\n", dataErr.Error())
		} else {

			// Write the data into a new file on server's side:
			log.Printf("Creating file at : %s\n", "./"+sourceFilename)
			ioutil.WriteFile("./"+sourceFilename, data, 0600)
			log.Println("Server: File was read from client and written to disk.")
		}
	}
}
