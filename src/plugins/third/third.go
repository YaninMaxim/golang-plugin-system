package main

func Something(a, b int) int {
	return MyBestFunc(a, b)
}

func MyBestFunc(a, b int) int {
	return a + (a * b) + b
}
