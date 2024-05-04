package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const tmpl = `./pages/index.html`

func main() {
	start_site()
}

func start_site() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pages/index.html")
	t.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(r.FormValue("number"))
	if err != nil {
		http.Error(w, "Error converting number", http.StatusBadRequest)
		return
	}

	fmt.Println(number)

	x1, x2 := findClosestFibonacciNumbers(number)

	fmt.Fprintf(w, "You entered: %d\nClosest fib: %d, %d", number, x1, x2)
}

func findClosestFibonacciNumbers(n int) (int, int) {
	first, second := 0, 1
	for second < n {
		fmt.Println(n, second)
		first, second = second, first+second
	}

	fmt.Println(n, second)
	fmt.Println("frst:")
	fmt.Println(first)

	return first, second
}
