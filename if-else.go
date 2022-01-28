package main

func main() {
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	if num := 9; num < 0 {
		println(num, "is negative")
	} else if num < 10 {
		println(num, "had 1 digit")
	} else {
		println(num, "had multiple digits")
	}
}
