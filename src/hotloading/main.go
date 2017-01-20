package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"plugin"
	"strings"
	"sync"
	"time"
)

const watchDir = "/go/bin/watch/"

var (
	mu      sync.Mutex
	plugins map[string]*plugin.Plugin
)

func init() {
	plugins = make(map[string]*plugin.Plugin)
}

func watch() {
	for {
		time.Sleep(time.Second)

		files, err := ioutil.ReadDir(watchDir)
		if err != nil {
			log.Print(err)
			continue
		}

		for _, file := range files {
			n := file.Name()
			mu.Lock()
			_, ok := plugins[n]
			mu.Unlock()
			if !ok {
				p, err := plugin.Open(watchDir + n)
				if err != nil {
					log.Print(err.Error())
					continue
				}
				mu.Lock()
				plugins[n] = p
				mu.Unlock()
			}
		}
	}
}

func main() {
	go watch()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pluginName := strings.Trim(r.URL.Path, "/")
		mu.Lock()
		p, ok := plugins[pluginName]
		mu.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "unable to find plugin")
			return
		}
		s, err := p.Lookup("Foo")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "unable to find symbol Foo")
			return
		}
		f, ok := s.(func() string)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "unable to type assert symbol")
			return
		}
		fmt.Fprintf(w, f())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
