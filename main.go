


package main

import (
	"errors"
	"fmt"
	"os"
)

func addNumber(numbers []int n int) []int {
	numbers = append(numbers, n)
	return numbers
}

func getMin(numbers []int) int {
	min := numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

func getMax(numbers []int) int {
	max := numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func getAverage(numbers []int) float64 {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return float64(sum) / float64(len(numbers))
}

func getStatistics(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("list is empty")
	}
	// Calling granular functions
	return getMin(numbers), getMax(numbers), getAverage(numbers), nil
}

func listNumbers(numbers []int) {
	if len(numbers) == 0 {
		fmt.Println("The list is empty.")
		return
	}
	fmt.Printf("Numbers: %v\n", numbers)
}

func removeByIndex(numbers []int, index int) ([]int, error) {
	if index < 0 || index >= len(numbers) {
		return nil, errors.New("index out of bounds")
	}
	return append(numbers[:index], numbers[index+1:]...), nil
}

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}


func main() {
    var numbers []int 
    var option int

    for {
        fmt.Println("\n--- Gerenciador de Números ---")
        fmt.Println("1) Adicionar número")
        fmt.Println("0) Sair")
        fmt.Print("Escolha uma opção: ")
        fmt.Scan(&opton)

        if option == 0 {
            break
        }

        switch option {
        case 1:
            var n int
            fmt.Print("Qual número você quer adicionar?")
            fmt.scan(&n)
            list = addNumbers(numbers, n)

        case 2:
            listNumbers(numbers)
        
        case 3:
            var idx int
			fmt.Print("Digite o índice que deseja remover: ")
			fmt.Scan(&idx)
			newList, err := removeByIndex(numbers, idx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				numbers = newList
			}

			
		case 4:
			min, max, avg, err := getStatistics(numbers)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Min: %d | Max: %d | Avg: %.2f\n", min, max, avg)
			}
		
		case 5:
			var n1, n2 int
			fmt.Print("Enter two numbers (a b): ")
			fmt.Scan(&n1, &n2)
			res, err := safeDivision(n1, n2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", res)
			}
		
		case 6:
			numbers = []int{}
			fmt.Println("Lista limpa.")

        default:
            fmt.Println("Opção inválida!")
        }
    }
}