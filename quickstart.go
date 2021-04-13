package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/osohq/go-oso"
	"github.com/osohq/go-oso/interfaces"
)

type User string

func (u User) EndsWith(pattern string) bool {
	return strings.HasSuffix(string(u), pattern)
}

func (u User) Equal(other interfaces.Comparer) bool {
	if v, ok := other.(User); ok {
		return u == v
	} else {
		return false
	}
}
func (u User) Lt(other interfaces.Comparer) bool {
	return false
}

type Expense struct {
	Amount      int
	Description string
	SubmittedBy User
}

type App struct {
	expenses map[int]Expense
	oso      oso.Oso
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Not Authorized!\n")
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found!\n")
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
			allowed, ok := a.oso.IsAllowed(User(actor), action, expense)
			if ok == nil && allowed {
				fmt.Fprintf(w, "Expense{%v}\n", expense)
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
	o.RegisterClass(reflect.TypeOf(Expense{}), nil)
	o.RegisterClass(reflect.TypeOf(User("")), nil)
	o.LoadFile("expenses.polar")

	expenses := make(map[int]Expense)
	expenses[1] = Expense{
		Amount: 500, Description: "coffee", SubmittedBy: User("alice@example.com"),
	}
	expenses[2] = Expense{
		Amount: 5000, Description: "software", SubmittedBy: User("alice@example.com"),
	}
	expenses[3] = Expense{
		Amount: 50000, Description: "flight", SubmittedBy: User("bhavik@example.com"),
	}
	app := App{expenses: expenses, oso: o}

	fmt.Println("server running on port 5050")
	http.Handle("/", &app)
	http.ListenAndServe(":5050", nil)
}
