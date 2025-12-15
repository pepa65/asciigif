package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/pepa65/asciigif/frames"
)

var version = "0.2.0"

var NotFoundMessage = map[string]string{
	"error": "Frameset not found. Navigate to /list for list of framesets. Navigate to https://github.com/pepa65/asciigif to submit new framesets.",
}

var NotCurledMessage = map[string]string{
	"error": "You almost ruined a good surprise. Come on, curl (or wget -O-) it in a terminal.",
}

var availableFrames []string
var defaultFrameRateMS int

func init() {
	for k := range frames.FrameMap {
		availableFrames = append(availableFrames, k)
	}
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

func listHandler(w http.ResponseWriter, r *http.Request) {
	writeJson(w, r, map[string][]string{"frames": availableFrames}, http.StatusOK)
	glog.Infof("- List request")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeJson(w, r, NotFoundMessage, http.StatusNotFound)
}

func notCurledHandler(w http.ResponseWriter, r *http.Request) {
	writeJson(w, r, NotCurledMessage, http.StatusExpectationFailed)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cn := w.(http.CloseNotifier)
	flusher := w.(http.Flusher)

	vars := mux.Vars(r)
	frameSource := vars["frameSource"]
	glog.Infof("= Frameset: %s", frameSource)
	frameRateMS := defaultFrameRateMS
	framerate, err := strconv.Atoi(r.URL.Query().Get("framerate"))
	if err == nil {
		frameRateMS = framerate
	}
	glog.Infof("- Framerate: %v", frameRateMS)

	frames, ok := frames.FrameMap[frameSource]
	if !ok {
		glog.Infof("# Frameset not found: %v", frameSource)
		notFoundHandler(w, r)
		return
	}

	userAgent := r.Header.Get("User-Agent")
	glog.Infof("- User-Agent: %v", userAgent)
	if !strings.Contains(userAgent, "curl") && !strings.Contains(userAgent, "Wget") {
		glog.Infof("# Unapproved User-Agent")
		notCurledHandler(w, r)
		return
	}

	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	i := 0
	for {
		select {
		// Handle client disconnects
		case <-cn.CloseNotify():
			glog.Infof("- Client stopped listening")
			return
		default:
			if i >= frames.GetLength() {
				i = 0
			}
			// Wait between reponses
			time.Sleep(time.Millisecond * time.Duration(frameRateMS))

			// Clear screen
			clearScreen := "\033[2J\033[H"
			newLine := "\n"

			// Write frames
			fmt.Fprint(w, clearScreen+frames.GetFrame(i)+newLine)
			//fmt.Fprint(w, strconv.Itoa(i))
			i++

			// Send some data.
			flusher.Flush()
		}
	}
}

// Server.
func main() {
	flag.IntVar(&defaultFrameRateMS, "framerate", 70, "Length of time to display each frame in milliseconds")
	port := flag.Int("port", 8080, "Port number to serve on")
	flag.Parse()
	// Don't write to /tmp - doesn't work in docker scratch
	flag.Set("logtostderr", "true")

	r := mux.NewRouter()
	r.HandleFunc("/list", listHandler).Methods("GET")
	r.HandleFunc("/{frameSource}", handler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(*port),
		// Set unlimited read/write timeouts
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	glog.Infof("* asciigif v%v serving on port %d with default framerate %d", version, *port, defaultFrameRateMS)
	glog.Fatal(srv.ListenAndServe())
}
