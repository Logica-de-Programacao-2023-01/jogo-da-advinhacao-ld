package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	contador := 0
	var matrizPalpites [][2]int

	for {
		contador++
		var palpite int
		fmt.Println("Adivinhe um número que está entre 1 e 100")
		fmt.Scan(&palpite)
		rand.Seed(time.Now().UnixNano())
		var resposta int
		resposta = rand.Intn(100) + 1
		var historicoPalpites []int

		for {
			if palpite == resposta {
				fmt.Println("Parabéns, você acertou! ")
				fmt.Println("O número de tentativas foi:", len(historicoPalpites))
				break
			} else {
				fmt.Println("Você errou, tente novamente.")
				historicoPalpites = append(historicoPalpites, palpite)

				if palpite > resposta {
					fmt.Println("O número resposta é menor.")
				} else {
					fmt.Println("O número resposta é maior.")
				}

				fmt.Scan(&palpite)
			}
		}

		matrizPalpites = append(matrizPalpites, [2]int{contador, len(historicoPalpites)})

		fmt.Println("Você deseja jogar novamente? (s/n)")
		var jogarNovamente string
		fmt.Scan(&jogarNovamente)

		if jogarNovamente == "n" {
			fmt.Println("Obrigado por jogar!")
			fmt.Println(matrizPalpites)
			break
		}
	}
}
