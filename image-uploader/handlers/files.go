package handlers

import (
	"go-microservices/image-uploader/files"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Files struct {
	log   *log.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{store: s, log: l}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fname := vars["filename"]

	f.log.Printf("[INFO] Handle POST, \"id\": %s, \"filename\": %s\n", id, fname)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fname, rw, r)
}

func (f *Files) invalidURI(URI string, rw http.ResponseWriter) {
	f.log.Printf("[ERROR] Invalid path, \"path\": %s\n", URI)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Printf("Save file for \"id\": %s, \"path\": %s\n", id, path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Printf("Unable to save file \"error\": %s\n", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
