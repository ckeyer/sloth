package views

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"strings"

	logpkg "github.com/ckeyer/go-log"
)

var (
	log = logpkg.GetDefaultLogger("goci-view")
)

type Views struct {
	Names map[string]struct{}
	Index string
}

func New() *Views {
	v := &Views{make(map[string]struct{}), "index.html"}
	return v
}

func (v *Views) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("route in static file")
	name := strings.TrimLeft(r.URL.Path, "/")
	if name == "" {
		name = v.Index
	}
	data, err := os.Open("assets/" + name)
	if err != nil && strings.HasSuffix(err.Error(), "no such file or directory") {
		log.Debug("no such file or directory")
		data, err = os.Open("assets/" + v.Index)
	}

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, data)
	return

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
