package views

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	uiDir = "assets/"
)

// Views
type Views struct {
	Names map[string]struct{}
	Index string
}

// New
func New() *Views {
	v := &Views{make(map[string]struct{}), "index.html"}
	return v
}

// ServeHTTP: webui
func (v *Views) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimLeft(r.URL.Path, "/")
	if name == "" {
		name = v.Index
	}
	data, err := os.Open(uiDir + name)
	if err != nil && strings.HasSuffix(err.Error(), "no such file or directory") {
		data, err = os.Open(uiDir + v.Index)
	} else {
		if strings.HasSuffix(name, "css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(name, "js") {
			w.Header().Set("Content-Type", "application/javascript")
		}
	}

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(200)
	io.Copy(w, data)

	// TODO
	if true {
		return
	}

	hdr := r.Header.Get("Accept-Encoding")
	if strings.Contains(hdr, "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		io.Copy(w, data)
	} else {
		gz, err := gzip.NewReader(data)
		if err != nil {
			log.Error(err)
			w.Write([]byte(err.Error()))
			return
		}
		io.Copy(w, gz)
		gz.Close()
	}
}

func SetUIDir(path string) {
	uiDir = strings.TrimSuffix(path, "/") + "/"
	log.Infof("set ths ui dir: %s", uiDir)
}
