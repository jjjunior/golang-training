package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // analisa argumentos, você precisa chamar isso sozinho
	fmt.Println(r.Form) // imprime as informações do formulário no lado do servidor
	fmt.Println("caminho", r.URL.Path)
	fmt.Println("esquema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("chave:", k)
		fmt.Println("valor:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Olá, astaxie!") // envia dados para o lado do cliente
}

func main() {
	http.HandleFunc("/", sayhelloName)       // define o roteador
	err := http.ListenAndServe(":9090", nil) // define a porta de escuta
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
