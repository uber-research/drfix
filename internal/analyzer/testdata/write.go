package main

var racyVar0 int

func main() {
	racyVar0 = 42 // This should be detected as a write
}
