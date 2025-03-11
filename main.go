package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var Host string = "localhost"
var Port int = 3333
var BaseURL string = "http://localhost:3333"
var InboxDir string = "./inbox"
var InboxPath string = "/inbox/"
var InboxWritable bool = true
var InboxPublic bool = true

type Message struct {
	Context  string   `json:"@context"`
	Id       string   `json:"@id"`
	Contains []string `json:"contains"`
}

func doInbox(w http.ResponseWriter, r *http.Request) {
	Logger.Printf("[INFO] %v %v %v\n", r.RemoteAddr, r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		doInboxGET(w, r)
	case http.MethodPost:
		if InboxWritable {
			doInboxPOST(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	case http.MethodHead:
		doInboxHEAD(w, r)
	case http.MethodOptions:
		doInboxOPTIONS(w, r)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}

func doInboxGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == InboxPath {
		doInboxListing(w, r)
	} else if InboxPublic {
		doInboxFile(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func doInboxFile(w http.ResponseWriter, r *http.Request) {
	localPath := strings.TrimPrefix(r.URL.Path, InboxPath)

	filePath := InboxDir + "/" + localPath

	content, err := os.ReadFile(filePath)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Page not found"))
		return
	}

	w.Header().Set("Content-Type", "application/ld+json")

	io.WriteString(w, string(content))
}

func doInboxListing(w http.ResponseWriter, r *http.Request) {
	m := Message{
		Context: "http://www.w3.org/ns/ldp",
		Id:      "http://localhost:3333/inbox/",
	}

	if InboxPublic {
		m.Contains = lsDir(InboxDir, BaseURL+InboxPath)
	} else {
		m.Contains = []string{}
	}

	b, _ := json.Marshal(m)

	io.WriteString(w, string(b))
}

func doInboxPOST(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != InboxPath {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if !(contentType == "application/ld+json" ||
		contentType == "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Need a content-type 'application/ld+json'")
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	path := storeBody(body)

	w.Header().Set("Location", BaseURL+InboxPath+path)
	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, "Accepted "+BaseURL+InboxPath+path)
}

func doInboxHEAD(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == InboxPath {
		meta := readMeta(InboxDir + "/.meta")

		if meta != nil {
			for key, value := range meta {
				w.Header().Set(key, fmt.Sprintf("%v", value))
			}
		}

		w.Header().Set("Content-Type", "application/ld+json")
	} else {
		localPath := strings.TrimPrefix(r.URL.Path, InboxPath)
		filePath := InboxDir + "/" + localPath

		if fileExists(filePath) {
			meta := readMeta(filePath)

			if meta != nil {
				for key, value := range meta {
					w.Header().Set(key, fmt.Sprintf("%v", value))
				}
			}

			w.Header().Set("Content-Type", "application/ld+json")
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}

func doInboxOPTIONS(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == InboxPath {
		allow := []string{"GET", "HEAD", "OPTIONS"}

		if InboxWritable {
			allow = append(allow, "POST")
			w.Header().Set("Allow-Post", "application/ld+json")
		}

		w.Header().Set("Allow", strings.Join(allow, ","))
		w.Header().Set("Content-Type", "application/ld+json")
	} else {
		localPath := strings.TrimPrefix(r.URL.Path, InboxPath)
		filePath := InboxDir + "/" + localPath

		if fileExists(filePath) {
			allow := []string{"GET", "HEAD", "OPTIONS"}

			w.Header().Set("Allow", strings.Join(allow, ","))
			w.Header().Set("Content-Type", "application/ld+json")
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}

func lsDir(dirPath string, baseUrl string) []string {
	matches, err := filepath.Glob(filepath.Join(dirPath, "*.jsonld"))

	if err != nil {
		fmt.Printf("Glob error: %v\n", err)
		return []string{}
	}

	result := []string{}

	for _, match := range matches {
		path := filepath.Base(match)
		result = append(result, baseUrl+path)
	}

	return result
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func storeBody(body []byte) string {
	hash := md5.New()
	io.WriteString(hash, string(body))
	id := fmt.Sprintf("%x", hash.Sum(nil))

	newPath := InboxDir + "/" + id + ".jsonld"

	os.WriteFile(newPath, body, 0444)

	return id + ".jsonld"
}

func readMeta(path string) map[string]interface{} {
	file, err := os.Open(path)

	if err != nil {
		return nil
	}

	content, _ := io.ReadAll(file)

	var result map[string]interface{}

	if err := json.Unmarshal(content, &result); err != nil {
		return nil
	}

	return result
}

func main() {
	host := flag.String("host", Host, "Hostname")
	port := flag.Int("port", Port, "Port")
	base := flag.String("base", BaseURL, "Base URL")
	inboxDir := flag.String("inboxDir", InboxDir, "Local path to your inbox")
	inboxPath := flag.String("inboxPath", InboxPath, "URL path to your inbox")
	public := flag.Bool("public", InboxPublic, "World readable inbox")
	writable := flag.Bool("writable", InboxWritable, "World appendable inbox")

	flag.Parse()

	BaseURL = *base
	InboxDir = *inboxDir
	InboxPath = *inboxPath
	InboxPublic = *public
	InboxWritable = *writable

	mux := http.NewServeMux()
	mux.HandleFunc(InboxPath, doInbox)

	Logger.Printf("Starting server on %v:%v\n", *host, *port)

	connection := fmt.Sprintf("%v:%v", *host, *port)

	err := http.ListenAndServe(connection, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
