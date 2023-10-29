package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

func validExt(filename string) bool {
	validExt := "(?i)\\.(jpg|jpeg|png)$"
	rgx := regexp.MustCompile(validExt)
	return rgx.MatchString(filename)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Couldn't get file from form data", http.StatusBadRequest)
			return
		}
		defer file.Close()

		if handler.Size > 1024 {
			http.Error(w, "File terlalu besar", http.StatusBadRequest)
		}

		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		filePath := filepath.Join(uploadDir, handler.Filename)

		// open new file to store data from uploaded file
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Couldn't create file", http.StatusBadRequest)
			return
		}
		defer newFile.Close()

		// copy data from uploaded file to new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Couldn't copy file data", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully")

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :2020")
	http.ListenAndServe(":2020", nil)
}
