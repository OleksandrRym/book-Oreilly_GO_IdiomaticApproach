package main

import "fmt"

func main() {
	initMap()
	idiomOk()
}
func nySet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
}
func idiomOk() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok) // ok - boolean type which say - map contain element or not
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)
}
func initMap() {
	var nilMap map[string]int              // nil
	var emptyLiteralMap = map[string]int{} // empty
	fmt.Println(nilMap == nil)
	fmt.Println(emptyLiteralMap == nil)

	teams := map[string][]string{
		"oleks":  []string{"java", "go"},
		"ivan":   []string{"ruby", "js"},
		"stepan": []string{"ts", "go"},
	}
	fmt.Println(teams)

	ages := make(map[string]int, 5)

	fmt.Println(ages)
}
