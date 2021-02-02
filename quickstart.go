package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/osohq/go-oso"
)

type Lib struct{}

func (Lib) HasSuffix(x string, y string) bool {
	return strings.HasSuffix(x, y)
}

type Expense struct {
	Amount      int
	Description string
	SubmittedBy string
}

type App struct {
	expenses map[int]Expense
	oso      oso.Oso
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Not Authorized!")
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found!")
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	resource := parts[1]
	id := parts[2]
	if resource == "expenses" {
		i, _ := strconv.Atoi(id)
		expense, ok := a.expenses[i]
		if ok {
			actor := r.Header.Get("user")
			action := r.Method
			allowed, ok := a.oso.IsAllowed(actor, action, expense)
			if ok == nil && allowed {
				fmt.Fprintf(w, "Expense{%v}", expense)
			} else {
				unauthorized(w)
			}
			return
		}
	}
	notFound(w)
}

func main() {
	o, _ := oso.NewOso()
	lib := Lib{}
	o.RegisterClass(reflect.TypeOf(lib))
	o.RegisterConstant(lib, "Lib")
	o.RegisterClass(reflect.TypeOf(Expense{}))
	o.LoadFile("policy.polar")
	results, _ := o.QueryStr("hello(x)")
	for result := range results {
		fmt.Println("Hello,", result["x"])
	}

	expenses := make(map[int]Expense)
	expenses[1] = Expense{
		Amount: 500, Description: "coffee", SubmittedBy: "alice@example.com",
	}
	expenses[2] = Expense{
		Amount: 5000, Description: "software", SubmittedBy: "alice@example.com",
	}
	expenses[3] = Expense{
		Amount: 50000, Description: "flight", SubmittedBy: "bhavik@example.com",
	}
	app := App{expenses: expenses, oso: o}

	http.Handle("/", &app)
	http.ListenAndServe(":8080", nil)
}
