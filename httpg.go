package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/drag0ns3c/httpg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	var port string
	var host string
	var dirPath string
	var help bool

	flag.BoolVar(&help, "help", false, "print cli help text")
	flag.StringVar(&dirPath, "dirPath", ".", "the directory to serve")
	flag.StringVar(&port, "port", "8080", "the port to listen on")
	flag.StringVar(&host, "host", "0.0.0.0", "the IP to listen on")
	flag.Parse()

	if help {
		fmt.Println("Usage: httpg [options]\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	r := mux.NewRouter()

	// sys info
	//templatesBox := packr.New("templates", "./templates")
	r.Handle("/__info__", handlers.SystemInfoHandler())

	// file server
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		log.Printf("failed to determine absolute path of dirPath: %s", dirPath)
		os.Exit(1)
	}

	fileHandler := handlers.CacheControl("no-store", http.FileServer(http.Dir(absPath)))
	r.PathPrefix("/").Handler(fileHandler)

	bindAddr := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Addr:         bindAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("httpg listening on %s", bindAddr)
		log.Printf("serving the contents of directory %s", absPath)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
