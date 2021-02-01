package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/osohq/go-oso"
)

type Expense struct {
	hello string
}

type App struct {
	expenses map[int]Expense
	oso      oso.Oso
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	resource := parts[1]
	id := parts[2]
	if resource == "expenses" {
		i, _ := strconv.Atoi(id)
		expense, ok := a.expenses[i]
		if ok {
			fmt.Fprintf(w, "Got one", expense)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404: not found")
}

func main() {
	o, _ := oso.NewOso()
	o.LoadFile("policy.polar")
	results, _ := o.QueryStr("hello(x)")
	for result := range results {
		fmt.Println("Hello,", result["x"])
	}

	expenses := make(map[int]Expense)
	expenses[0] = Expense{hello: "world"}
	app := App{expenses: expenses, oso: o}

	http.Handle("/", &app)
	http.ListenAndServe(":8080", nil)
}
