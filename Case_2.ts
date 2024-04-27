/**
 * Mescla dois objetos e retorna um novo objeto que contém propriedades de ambos os objetos.
 * Se existir uma propriedade em ambos os objetos, o valor do segundo objeto substituirá o valor do primeiro objeto.
 * Objetos aninhados são mesclados recursivamente.
 *
 * @template T – O tipo do primeiro objeto.
 * @template U – O tipo do segundo objeto.
 * @param {T} obj1 – O primeiro objeto a ser mesclado.
 * @param {U} obj2 – O segundo objeto a ser mesclado.
 * @returns {T & U} – Um novo objeto que contém propriedades de ambos os objetos.
 */
function mergeObjects<T, U>(obj1: T, obj2: U): T & U {
  // T & U é a interseção de T e U
  const result: any = { ...obj1 }; // Cria um novo objeto com as propriedades de obj1

  for (const key in obj2) {
    // Loop através das propriedades de obj2
    if (Object.prototype.hasOwnProperty.call(obj2, key)) {
      // Verifica se a propriedade é do próprio objeto
      if (typeof obj2[key] === "object" && obj2[key] !== null) {
        // Verifica se a propriedade é um objeto
        if (typeof result[key] === "object" && result[key] !== null) {
          // Verifica se a propriedade é um objeto
          result[key] = mergeObjects(result[key], obj2[key]); // Mescla objetos aninhados recursivamente
        } else {
          // Se a propriedade não é um objeto
          result[key] = obj2[key]; // Substitui o valor da propriedade
        }
      } else {
        // Se a propriedade não é um objeto
        result[key] = obj2[key]; // Substitui o valor da propriedade
      }
    }
  }

  return result as T & U; // Retorna o novo objeto
}

// Exemplos de uso
console.log(mergeObjects({ foo: true }, { bar: false })); // { foo: true, bar: false }
console.log(mergeObjects({ foo: true, nested: { bar: 0 } }, { foo: false })); // { foo: false, nested: { bar: 0 } }
console.log(mergeObjects({ foo: true }, { foo: false })); // { foo: false }
console.log(mergeObjects({ array: [2, 4] }, { array: [1, 3] })); // { array: [1, 3] }
