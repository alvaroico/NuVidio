/**
 * Calcula o comprimento da substring mais longa sem repetir caracteres em uma determinada string.
 * @param s - O input e tipo string.
 * @returns O comprimento da substring mais longa sem caracteres repetidos.
 *
 * lengthOfLongestSubstring()
 * é usada para calcular o comprimento da maior substring sem caracteres
 * repetidos em uma string fornecida. A função recebe uma string s como entrada e retorna
 * um número que representa o comprimento da maior substring sem caracteres repetidos.
 */
function lengthOfLongestSubstring(s: string): number {
  let maxTamanho = 0; // Comprimento da substring mais longa sem caracteres repetidos
  let inicio = 0; // Início da substring
  const charIndexMap: { [key: string]: number } = {}; // Mapeia o caractere para o seu índice

  for (let ate = 0; ate < s.length; ate++) {
    // Loop através da string
    const char = s[ate]; // Caractere atual
    if (charIndexMap[char] !== undefined) {
      // Se o caractere já existe no mapa
      inicio = Math.max(inicio, charIndexMap[char] + 1); // Atualiza o início da substring
    }
    charIndexMap[char] = ate; // Atualiza o índice do caractere
    maxTamanho = Math.max(maxTamanho, ate - inicio + 1); // Atualiza o comprimento da substring
  }

  return maxTamanho;
}

// Exemplos de uso
console.log(lengthOfLongestSubstring("abcabcbb")); // 3
console.log(lengthOfLongestSubstring("pwwkew")); // 3
