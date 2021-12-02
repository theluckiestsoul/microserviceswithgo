package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/theluckiestoul/microserviceswithgo/product-images/files"
)

// Files is a handler for reading and writting files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadREST handler implements the http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fn, rw, r)
}

// UploadMultipart form data
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart from data", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form for id", "id", id)

	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected integer id", http.StatusBadRequest)
		return
	}

	_, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}
	f.saveFile(r.FormValue("id"), mh.Filename, rw, r)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
