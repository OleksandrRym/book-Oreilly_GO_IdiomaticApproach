package main

func main() {

}

func init() {
	type Score int
	type Converter func(string) Score
	type TeamScores map[string]Score
}
