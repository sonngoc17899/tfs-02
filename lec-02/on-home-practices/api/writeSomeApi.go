package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/sum", apiSum)
	http.HandleFunc("/sub", apiSub)
	http.HandleFunc("/mul", apiMul)
	http.HandleFunc("/div", apiDiv)
	http.ListenAndServe(":8080", nil)
}

type Input struct {
	A string `json:"a"`
	B string `json:"b"`
}

type Sum struct {
	Sum string `json:"sum"`
}

type Sub struct {
	Sub int `json:"sub"`
}

type Mul struct {
	Mul int `json:"mul"`
}

type Div struct {
	Div int `json:"div"`
}

func apiSum(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	i := Sum{
		Sum: a[0] + b[0],
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiSub(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	intA, err := strconv.Atoi(a[0])
	intB, err := strconv.Atoi(b[0])
	i := Sub{
		Sub: intA - intB,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiMul(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	intA, err := strconv.Atoi(a[0])
	intB, err := strconv.Atoi(b[0])
	i := Mul{
		Mul: intA * intB,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiDiv(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	intA, err := strconv.Atoi(a[0])
	intB, err := strconv.Atoi(b[0])
	i := Div{
		Div: intA / intB,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}
