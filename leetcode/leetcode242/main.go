package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false

	fmt.Println(isAnagram("oso", "soo"))
}

/*Given two strings s and t, return true if t is an anagram of s, and false otherwise.



Example 1:

Input: s = "anagram", t = "nagaram"

Output: true

Example 2:

Input: s = "rat", t = "car"

Output: false

 Un anagrama significa:
👉 ambas palabras tienen las mismas letras
👉 en igual cantidad
👉 pero pueden estar en diferente orden

Ejemplos:

"anagram" y "nagaram" → true

"rat" y "car" → false

✅ ¿Qué hace este código?
✔️ 1. Crear un mapa para contar letras
count := make(map[rune]int)


Esto crea un mapa donde:

clave (key) = una letra del string (rune)

valor (value) = cuántas veces aparece esa letra (int)



*/
