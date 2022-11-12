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
/*
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
*/

package main

import "fmt"

func main(){
    var i, n, div, mod, a, j1, j2 int
    
    fmt.Scan(&n)
    for i = 0; i < n; i++ {
        fmt.Scan(&a)
        div = a/3
        mod = a%3
        j1 += div
        j2 += mod
    }
    fmt.Println(j1, j2)
}