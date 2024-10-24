package main

import (
	"fmt"
	"log"
	"strconv"

	//"log"
	//"strconv"

	//"math"
	"strings"
)

/*
func main() {
	var a float64
	var b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(math.Sqrt(a*a + b*b))
}
*/

///////////////////////
/*
func main() {
	var str string
	fmt.Scanln(&str)
	fmt.Print(strings.Join(strings.Split(str, ""), "*"))

}
*/
////////////////////////
/*
func T(w float64) float64 {
	t := 6 / w
	return t
}

func W(k, m float64) float64 {
	w := math.Sqrt(k / m)
	return w
}

func M(p, v float64) float64 {
	var m1 = p * v
	return m1
}

func main() {
	var k, p, v float64
	fmt.Scan(&k, &p, &v)
	m := M(p, v)
	w := W(k, m)
	t := T(w)
	fmt.Println(t)
}
*/
/*
func main() {
	var intgr int
	fmt.Scan(&intgr)
	var str string = strconv.Itoa(intgr)
	var str1 []string = strings.Split(str, "")
	for i := 0; i < len(str1); i++ {
		a, err := strconv.Atoi(str1[i])
		if err != nil {
			log.Fatal(err)
		}
		f := strconv.Itoa(a * a)
		str1[i] = f

	}
	fmt.Println()
	for i := 0; i < len(str1); i++ {

		fmt.Printf(str1[i])
	}
	fmt.Println()
}
*/
/////////////////////////

func main() {

	var str string
	fmt.Scan(&str)
	str1 := strings.Split(str, "")
	var max int = 0
	for i := 0; i < len(str); i++ {
		a, err := strconv.Atoi(str1[i])
		if err != nil {
			log.Fatal(err)
		}
		if a > max {
			max = a
		}
	}
	fmt.Println(max)
}
