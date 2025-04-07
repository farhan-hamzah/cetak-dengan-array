package main
import "fmt"

func main(){
	const NMAX int = 10
	var i, n int
	var A [NMAX]string
	fmt.Scan(&n)
	for i = 0; i<n;i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i<n; i++{
		fmt.Println(A[i])
	}
}