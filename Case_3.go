package main

import "fmt"

// é usada para calcular o comprimento da maior substring sem caracteres
// repetidos em uma string fornecida. A função recebe uma string s como entrada e retorna
// um número que representa o comprimento da maior substring sem caracteres repetidos.
func lengthOfLongestSubstring(s string) int { 
	charIndexMap := make(map[rune]int) // Cria um map para armazenar o índice de cada caractere
	maxTamanho := 0 									 // Inicializa o tamanho máximo da substring como 0
	inicio := 0 											 // Inicializa o índice de início da substring como 0

	for ate, char := range s { // Percorre a string
		if index, ok := charIndexMap[char]; ok && index >= inicio { // Verifica se o caractere já foi encontrado
			inicio = index + 1 // Atualiza o índice de início da substring
		}
		charIndexMap[char] = ate // Atualiza o índice do caractere
		if ate-inicio+1 > maxTamanho {	// Verifica se o tamanho da substring é maior que o tamanho máximo
			maxTamanho = ate - inicio + 1 // Atualiza o tamanho máximo da substring
		}
	}

	return maxTamanho // Retorna o tamanho máximo da substring
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
}
