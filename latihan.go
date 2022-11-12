// Pseudocode soal login
/*
program Login
kamus
	un, pw : string

algoritma
	imput(un, pw)

*/

/*
package main

import "fmt"

func main(){
	var xi, total int
	total=0
	fmt.Scan(&xi)
	for xi % 2 == 0 {
		total = total + xi
		fmt.Scan(&xi)
	}
	fmt.Println(total)
}

*/

package main

import "fmt"

func main(){
	var a, b, i int
	i=1
	fmt.Scan(&a, &b)
	for !((a==b) || ((a/10)==b) || ((a%10)==b)) {	
		i++
		fmt.Scan(&b)
	}
	fmt.Println(i)
}