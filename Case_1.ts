/**
 * Converte quaisquer valores indefinidos em um objeto ou array em nulo.
 *
 * @param input – O objeto ou array de entrada.
 * @returns O objeto de entrada ou matriz com valores indefinidos substituídos por nulos.
 */
function undefinedToNull<T>(input: T): T {
  // T é o tipo de input
  if (Array.isArray(input)) {
    // Verifica se input é uma matriz
    return input.map((item) => (item === undefined ? null : item)) as T; // Substitui valores indefinidos por nulos
  } else if (typeof input === "object" && input !== null) {
    // Verifica se input é um objeto
    const result: any = {}; // Cria um novo objeto
    for (const key in input) {
      // Loop através das propriedades do objeto
      if (Object.prototype.hasOwnProperty.call(input, key)) {
        // Verifica se a propriedade é do próprio objeto
        result[key] = input[key] === undefined ? null : input[key]; // Substitui valores indefinidos por nulos
      }
    }
    return result as T; // Retorna o novo objeto
  } else {
    // Se input não é um objeto ou matriz
    return input; // Retorna input sem alterações
  }
}

// Exemplos de uso
console.log(undefinedToNull({ a: undefined, b: 3 })); // { a: null, b: 3 }
console.log(undefinedToNull([undefined, undefined])); // [null, null]
