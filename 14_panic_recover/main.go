package main

import "fmt"

func tryRecover() {
	fmt.Println("trying recover code...")

	if rec := recover(); rec != nil {
		fmt.Println("recover code success!")
	}
}

func media(n1, n2 int) string {
	defer tryRecover()
	media := (n1 + n2) / 2

	if media < 6 {
		return "reprovado"
	} else if media > 6 {
		return "aprovado"
	}

	panic("media exatamente igual a 6, entrei em chok!!")
}

func main() {
	fmt.Println(media(4, 8))
}
