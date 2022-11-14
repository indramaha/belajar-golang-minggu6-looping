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

/*
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

*/

//tugas1
/*
package main

import "fmt"

func main() {
    var a, b int

    fmt.Scan(&a, &b)
    for i:=a; a <= b; i++ {
		fmt.Print(i)
    }
}
*/

//tugas2
/*
package main
import "fmt"

func main() {
    var a int
    var str string

    fmt.Scan(&a, &str)
    for i:=1; i<=a; i++ {
        fmt.Println(str)
    }
}
*/

//tugas3
/*
package main
import "fmt"

func main() {
    var n, m int
    var hasil float64

    fmt.Scan(&n, &m)
    for i:=n; i<=m; i++ {
        hasil += (4/float64(i))
    }
    fmt.Printf("%.2f", hasil)
}
*/

// tugas4
/*
package main

import "fmt"

func main(){
    var n, m, hasil, b1, b2 int

    fmt.Scanln(&n)
    for i:=1; i<=n; i++ {
        fmt.Scanln(&m)
        b1= ((m/1000)%10)
        b2= ((m/10)%10)
        hasil += b1 +b2
    }
    fmt.Println(hasil)
}
*/

// soal penjumlahan terus menerus sampai diatas 100
/*
package main

import "fmt"

func main(){
    var x, hasil, i int

    hasil = 0
    for i=0; hasil <= 100; i++ {
        fmt.Scan(&x)
        hasil += x
    }
    fmt.Println(hasil, i)
}
*/

// soal input persegi
/*
package main

import "fmt"

func main(){
    var n, p, l, kll int

    fmt.Scan(&n)
    for i:=1; i <= n; i++ {
        fmt.Scan(&p)
        l = p *p
        kll = 4*p
        fmt.Println(l, kll)
    }

}
*/

// jumlahkan deret 1 3 5 sampai 99 DENGAN FOR
/*
package main

import "fmt"

func main(){
    var jumlah int
    jumlah = 0
    i:= 1
    for i<100 {
        jumlah += i
        i += 2
    }
    fmt.Println(jumlah)
}
*/

// soal 1/1, 1/2, ..., 1/n

package main

import "fmt"

func main(){
    var input,hasil float64

    hasil = 1
    fmt.Scan(&input)
    for input != 0 {
        hasil *= (1/input)
        fmt.Scan(&input)
    }
    fmt.Printf("%.5f", hasil)
}