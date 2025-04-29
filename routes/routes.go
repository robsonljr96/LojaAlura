package routes

import (
	"Loja1/controlers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controlers.Index)
	http.HandleFunc("/new", controlers.New)
	http.HandleFunc("/insert", controlers.Insert)
	http.HandleFunc("/delete", controlers.Delete)
	http.HandleFunc("/edit", controlers.Edit)
	http.HandleFunc("/update", controlers.Update)

}
