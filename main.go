package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/webview/webview"

	"net"
)

//go:embed public
var embededFiles embed.FS

// main function
func main() {
	// channel to get the web prefix
	prefixChannel := make(chan string)
	// run the web server in a separate goroutine
	go app(prefixChannel)
	prefix := <-prefixChannel
	// create a web view
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Go webapp template")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(prefix + "/public/html/index.html")
	w.Run()
}

// web app
func app(prefixChannel chan string) {
	fsys, err := fs.Sub(embededFiles, "public")
	if err != nil {
		panic(err)
	}
	fmt.Println()

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(fsys))))

	// get an ephemeral port, so we're guaranteed not to conflict with anything else
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	portAddress := listener.Addr().String()
	prefixChannel <- "http://" + portAddress
	listener.Close()
	server := &http.Server{
		Addr:    portAddress,
		Handler: mux,
	}
	server.ListenAndServe()
}
