package main

import (
	fib "basid_algos_server/Fibonacci"
	task "basid_algos_server/task_handler"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var asd task.Taskhandler

var validPath = regexp.MustCompile("^/(fibonacci|main)/([a-zA-Z0-9]*)$")

func fibonacci(w http.ResponseWriter, r *http.Request, title string) {
	switch r.Method {
	case "GET":
		v, err := asd.GetWorks()
		if err != nil {
			fmt.Fprintf(w, "asd")
		}
		fmt.Fprintf(w, strconv.Itoa(v))
	case "POST":
		v, err := strconv.Atoi(r.FormValue("n"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "fibonacci(%d) = %d", v, fib.Work(v))
	}
}

func enableCors(w *http.ResponseWriter) {
	// This is a hack. Replace eventually
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	asd = *new(task.Taskhandler)
	http.HandleFunc("/fibonacci/", makeHandler(fibonacci))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
