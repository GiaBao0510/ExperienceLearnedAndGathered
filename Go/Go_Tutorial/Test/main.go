package main

import (
	"fmt"
)

// func Test(s string) string{
	
// }

func reverse(s string) string{
	rns := []rune(s)

	for l, r := 0, len(rns)-1; l < r; l, r = l+1, r-1{
		rns[l], rns[r] = rns[r], rns[l]
	}

	return string(rns)
}

func Test(s string, k int) string{

	result := ""
	i := 0
	fmt.Printf("\t string: %s, k: %d, Len_string: %d\n", s, k, len(s))

	for  {

		fmt.Printf(" \n---- Vòng lặp i: %d -> i: %d\n", i, i + 2*k - 1)
		if 2*k + i < len(s){
			fmt.Printf(" \n---- Vòng lặp i: %d (%v) -> i: %d(%v)\n", i, string(s[i]), i + 2*k, string(s[i+2*k]))
			result += s[i: k + i]
			result += s[ k + i : 2*k +i]

			fmt.Printf("[%d, %d] = %s\n", i, k + i, s[i:k + i])
			fmt.Printf("[%d, %d] = %s\n", k + i , 2*k +i , s[k + i: 2*k +i ])

			i+= 2*k + i
		} else {
			fmt.Println("\n\tI: ",i)
			break
		}
	}

	

	if i < len(s){
		fmt.Println("finally i: ",i)
		result += s[i:]
	}

	return result
}

func main() {
	s := "abcdefghkij"
	fmt.Println(s[6:8])
}