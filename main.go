package main

import (
	"fmt"
)

func groupAnagrams(arr []string) [][]string {
	hasil := make([][]string, 0)
	tampung := make(map[[26]byte]int)

	for i := range arr {
		/**
		 * 26 adalah alphabet, contoh: 'abcdefghijklmnopqrstuvwxyz'
		 */
		list := [26]byte{}
		for j := range arr[i] {
			/**
			 * 'a' adalah rune dari "a", dgn ASCII code 97 (lowercase)
			 * misal: huruf nya adalah a, maka words[i][j]-'a' == 0, contoh: 97 - 97 == 0
			 * maka huruf nya diletakkan pada index 0
			 * jika huruf nya adalah b, maka words[i][j]-'a' == 1, maka huruf b akan diletakkan di index 1
			 * lakukan berulang kali (increment)
			 */
			list[arr[i][j]-'a']++
		}

		/**
		 * jika list menemukan kata yang sama, maka masukkan ke 'tampung'
		 * value tampung akan menjadi index, lakukan append
		 */
		if idx, ok := tampung[list]; ok {
			hasil[idx] = append(hasil[idx], arr[i])
		} else {
			hasil = append(hasil, []string{arr[i]})
			tampung[list] = len(hasil) - 1
		}
	}
	return hasil
}

func main() {
	var arr = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	result := groupAnagrams(arr)
	fmt.Println("[")
	for _, res := range result {
		fmt.Printf("%v, \n", res)
	}
	fmt.Println("]")
}
