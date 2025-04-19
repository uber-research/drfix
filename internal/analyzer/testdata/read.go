package main

var racyVar0 int

func main() {
	_ = racyVar0 // This should not be detected as a write
}
