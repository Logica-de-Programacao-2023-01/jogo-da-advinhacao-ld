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

	var matrixTentativas [][]int

	for {
		resposta := rand.Intn(100) + 1
		tentativas := []int{}

		fmt.Println("Adivinhe o número entre 1 e 100:")

		for {
			palpite, err := obterPalpite()
			if err != nil {
				fmt.Println(err)
				continue
			}

			tentativas = append(tentativas, palpite)

			if palpite == resposta {
				fmt.Println("Resposta correta!")
				break
			} else if palpite < resposta {
				fmt.Println("Errou. O número é maior.")
			} else {
				fmt.Println("Errou. O número é menor.")
			}
		}

		matrixTentativas = append(matrixTentativas, tentativas)
		fmt.Printf("Número de tentativas: %d\n", len(tentativas))

		if !jogarNovamente() {
			break
		}
	}

	fmt.Println("\nTodas as tentativas:")
	for i, tentativas := range matrixTentativas {
		fmt.Printf("Jogada %d: %v (Número de tentativas: %d)\n", i+1, tentativas, len(tentativas))
	}
}

func obterPalpite() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	palpiteStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	palpite, err := strconv.Atoi(strings.TrimSpace(palpiteStr))
	if err != nil {
		return 0, fmt.Errorf("Palpite incorreto. Digite outro número")
	}

	return palpite, nil
}

func jogarNovamente() bool {
	var jogarNovamente string
	fmt.Println("Deseja jogar novamente? (s/n)")

	reader := bufio.NewReader(os.Stdin)
	jogarNovamente, _ = reader.ReadString('\n')

	jogarNovamente = strings.TrimSpace(strings.ToLower(jogarNovamente))
	return jogarNovamente == "s"
}
