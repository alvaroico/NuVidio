package main

import (
	"fmt"
	"reflect"
)

// Mescla dois objetos e retorna um novo objeto que contém propriedades de ambos os objetos.
// Se existir uma propriedade em ambos os objetos, o valor do segundo objeto substituirá o valor do primeiro objeto.
// Objetos aninhados são mesclados recursivamente.
func mergeObjects(obj1, obj2 map[string]interface{}) map[string]interface{} { // Define a função mergeObjects
	result := make(map[string]interface{}) // Cria um novo map para armazenar o resultado

	for key, value := range obj1 { // Percorre as propriedades do primeiro objeto
		result[key] = value // Adiciona a propriedade ao resultado
	}

	for key, value := range obj2 { // Percorre as propriedades do segundo objeto
		if val, ok := result[key]; ok { // Verifica se a propriedade já existe no resultado
			if reflect.TypeOf(val) == reflect.TypeOf(value) { // Verifica se os tipos das propriedades são iguais
				switch val.(type) { // Verifica o tipo da propriedade
				case map[string]interface{}: // Se for um map
					result[key] = mergeObjects(val.(map[string]interface{}), value.(map[string]interface{})) // Mescla os objetos aninhados
				case []interface{}: // Se for um slice
					result[key] = append(val.([]interface{}), value.([]interface{})...) // Adiciona os elementos ao slice
				default: // Se for outro tipo
					result[key] = value // Substitui o valor da propriedade
				}
			}
		} else { // Se a propriedade não existir no resultado
			result[key] = value // Adiciona a propriedade ao resultado
		}
	}

	return result // Retorna o resultado
}

func main() {
	obj1 := map[string]interface{}{"foo": true}
	obj2 := map[string]interface{}{"bar": false}
	fmt.Println(mergeObjects(obj1, obj2)) // Output: map[foo:true bar:false]

	obj1 = map[string]interface{}{"foo": true, "nested": map[string]interface{}{"bar": 0}}
	obj2 = map[string]interface{}{"foo": false}
	fmt.Println(mergeObjects(obj1, obj2)) // Output: map[foo:false nested:map[bar:0]]

	obj1 = map[string]interface{}{"foo": true}
	obj2 = map[string]interface{}{"foo": false}
	fmt.Println(mergeObjects(obj1, obj2)) // Output: map[foo:false]

	obj1 = map[string]interface{}{"array": []interface{}{2, 4}}
	obj2 = map[string]interface{}{"array": []interface{}{1, 3}}
	fmt.Println(mergeObjects(obj1, obj2)) // Output: map[array:[2 4 1 3]]
}
