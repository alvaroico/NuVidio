package main

import (
	"fmt"
)

// undefinedToNull converte recursivamente todas as ocorrências de `nil`(Nada) em uma determinada entrada para `nil`.
// Ele percorre mapas e fatias para encontrar valores `nil`(Nada) e os substitui por `nil`.
// Para qualquer outro tipo, retorna o valor como está.

func undefinedToNull(input interface{}) interface{} { // Define a função undefinedToNull
	switch v := input.(type) { // Verifica o tipo da entrada
	case map[string]interface{}: // Se for um map
		result := make(map[string]interface{}) // Cria um novo map para armazenar o resultado
		for key, val := range v { // Percorre as propriedades do map
			if val == nil { // Verifica se o valor é nil
				result[key] = nil // Adiciona nil ao resultado
			} else { // Se o valor não for nil
				result[key] = undefinedToNull(val) // Chama a função recursivamente
			}
		}
		return result // Retorna o resultado
	case []interface{}: // Se for um slice
		result := make([]interface{}, len(v)) // Cria um novo slice para armazenar o resultado
		for i, val := range v { // Percorre os elementos do slice
			if val == nil { // Verifica se o valor é nil
				result[i] = nil	// Adiciona nil ao resultado
			} else { // Se o valor não for nil
				result[i] = undefinedToNull(val) // Chama a função recursivamente
			}
		}
		return result // Retorna o resultado
	case nil: // Se for nil
		return nil // Retorna nil
	default: // Se for outro tipo
		return v // Retorna o valor
	}
}

func main() {
	fmt.Println(undefinedToNull(map[string]interface{}{"a": nil, "b": 3})) // Output: map[a:<nil> b:3]
	fmt.Println(undefinedToNull([]interface{}{nil, nil}))                   // Output: [<nil> <nil>]
}
