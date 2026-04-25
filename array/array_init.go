package main

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for i := range len(a) {
		println(a[i])
	}
}
