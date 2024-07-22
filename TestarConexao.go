package main

import (
	"fmt"
	"net/http"
)

func requisicao(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site :", site, "foi carregado com sucesso.", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}

}

/*func main() {
	fmt.Println("----------------Testa host----------------")
	requisicao("https://www.google.com")

}
*/
