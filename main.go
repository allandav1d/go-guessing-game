package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número será sorteado. Tente acertar. O número é um inteiro entre 0 e 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Digite o seu chute: ")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		fmt.Println("Seu chute foi: ", chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro.")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf("Parabéns! Você acertou! O número sorteado era %d.\n"+
				"Você acertou em %d tentativas.\n"+
				"Estes foram os seus chutes: %v\n", x, i+1, chutes[:i])
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf("Infelizmente você não acertou o número. O número sorteado foi: %d. \nVocê teve 10 chutes. \nEstes foram os seus chutes: %v\n", x, chutes)
}
