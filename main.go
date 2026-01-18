package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pepa65/asciigif/frames"
)

var version = "0.16.0"
var clientnumber = 0
var repo = "github.com/pepa65/asciigif"
var defaultFramepause = 70
var defaultPort = 8080
var NoFrameMessage = map[string]string{"error": "No frameset given, after domain give /frameset. Give /list for a list of framesets."}
var NotFoundMessage = map[string]string{"error": "Frameset not present. Give /list for a list of framesets."}
var NotCurledMessage = map[string]string{"error": "Not for graphical browsers, use curl (or wget -qO-) in a terminal."}
var availableFrames []string
var defFramepause int

func usage() {
	fmt.Println("asciigif v" + version + ` - Ascii-gifs served for terminal consumption
Repo:  ` + repo + `
Usage: asciigif [--ms MS] [--port PORT] [--list] [--version] [-h|--help]
    --ms MS:      Display-time of each frame in milliseconds (default ` + strconv.Itoa(defaultFramepause) + `)
    --port PORT:  Port number to serve on (default ` + strconv.Itoa(defaultPort) + `)
    --list:       Show available framesets
    --version:    Show version
    --help:       Show this help text`)
}

func init() {
	for k := range frames.FrameMap {
		availableFrames = append(availableFrames, k)
	}
	sort.Strings(availableFrames)
}

func writeJson(w http.ResponseWriter, r *http.Request, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(res)
	if err != nil {
		return
	}
	w.WriteHeader(status)
	fmt.Fprint(w, string(data) + "\n")
}

func noFrameHandler(w http.ResponseWriter, r *http.Request) {
	clientnumber += 1
	writeJson(w, r, NoFrameMessage, http.StatusNotFound)
	fmt.Fprintf(os.Stderr, "### No frame requested (%d)\n", clientnumber)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cn := w.(http.CloseNotifier)
	flusher := w.(http.Flusher)

	clientnumber += 1
	client := clientnumber
	userAgent := strings.Split(r.Header.Get("User-Agent"), " ")
	vars := mux.Vars(r)
	frameSource := vars["frameSource"]
	if frameSource == "list" {
		writeJson(w, r, map[string]interface{}{"repo": repo, "version": version, "frames": availableFrames}, http.StatusOK)
		fmt.Fprintf(os.Stderr, "--- List request from User-Agent: %v (%d)\n", userAgent[0], client)
		return
	}

	// No /list
	if !strings.Contains(userAgent[0], "curl") && !strings.Contains(userAgent[0], "Wget") {
		fmt.Fprintf(os.Stderr, "### Request '%v' from unapproved User-Agent: %v (%d)\n", frameSource, userAgent[0], client)
		writeJson(w, r, NotCurledMessage, http.StatusExpectationFailed)
		return
	}

	// Approved User-Agent
	frames, ok := frames.FrameMap[frameSource]
	if !ok {
		fmt.Fprintf(os.Stderr, "### Frameset '%v' not found for User-Agent %v (%d)\n", frameSource, userAgent[0], client)
		writeJson(w, r, NotFoundMessage, http.StatusNotFound)
		return
	}

	// Existing frameset
	framepause, err := strconv.Atoi(r.URL.Query().Get("ms"))
	if err != nil {
		framepause = defFramepause
	}
	fmt.Fprintf(os.Stderr, "--- Request '%v' at framepause %d with User-Agent: %v (%d)\n", frameSource, framepause, userAgent[0], client)

	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	i := 0
	for {
		select {
		// Handle client disconnects
		case <-cn.CloseNotify():
			fmt.Fprintf(os.Stderr, "### User-Agent %v (%d) stopped listening\n", userAgent[0], client)
			return
		default:
			if i >= frames.GetLength() {
				i = 0
			}
			clearScreen, newline := "\033[2J\033[H", "\n"
			fmt.Fprint(w, clearScreen+frames.GetFrame(i)+newline)
			flusher.Flush()
			time.Sleep(time.Millisecond * time.Duration(framepause))
			i++
		}
	}
}

// Server
func main() {
	flag.Usage = usage
	flag.IntVar(&defFramepause, "ms", defaultFramepause, "Display time of each frame in milliseconds")
	port := flag.Int("port", defaultPort, "Port number to serve on")
	list := flag.Bool("list", false, "Show available framesets")
	vers := flag.Bool("version", false, "Show version")
	flag.Parse()
	if *vers {
		fmt.Println("asciigif v" + version)
		return
	}
	if *list {
		fmt.Println("asciigif v" + version + " framesets:")
		fmt.Println(strings.Join(availableFrames, " "))
		return
	}
	// Don't write to /tmp - doesn't work in docker scratch
	flag.Set("logtostderr", "true")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(noFrameHandler)
	r.HandleFunc("/{frameSource}", handler).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(*port),
		// Set unlimited read/write timeouts
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	fmt.Fprintf(os.Stderr, "=== asciigif v%v serving on port %d with default framepause %d\n", version, *port, defFramepause)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
