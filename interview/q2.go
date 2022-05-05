package main

func main() {
	var s = "string"
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] == s[i] {
				println(false)
				return
			}
		}
	}

	println(true)
}
