package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		resposta := rand.Intn(100) + 1

		fmt.Println("Adivinhe o número entre 1 e 100:")

		for {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			palpite, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Palpite inválido. Digite um número inteiro")
				continue
			}

			if palpite == resposta {
				fmt.Println("Resposta correta!")
				break
			} else if palpite < resposta {
				fmt.Println("Errou. O número é maior.")
			} else {
				fmt.Println("Errou. O número é menor.")
			}
		}

		fmt.Println("Deseja jogar novamente? (s/n)")
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "s" {
			break
		}
	}

	fmt.Println("Obrigado por jogar!")
}
