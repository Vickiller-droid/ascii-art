package main

func main() {
	input, banner := validate()
	lines := loadBanner(banner)
	generate(input, lines)
}
