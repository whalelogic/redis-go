package main

import (
		"fmt"
		"sort"
)


type Sequence []int
type StringSequence []string

func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

func (s Sequence) String() string {
    s = s.Copy() 
    sort.Sort(s)
    str := "["
    for i, elem := range s { 
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}

func modify(val *int) {
    *val = *val * *val 
}


func Finder(str []string, target string) (string, bool) {
	for i := 0; i < len(str); i++ {
		if str[i] == target {
			fmt.Printf("Found ==== %s\n", str[i])
			return str[i], true
		}

	}
	return "Not Found.", false

}


func main() {

	var seq Sequence = []int{3,4,13,63,1,9,44,2,5}
	var seq2 Sequence = []int{9,43,13,63,11,89,34,28,50}
	fmt.Println("Original:", seq)
	fmt.Println("Sorted:  ", seq.String())
	fmt.Println(seq.Len())
	fmt.Println(seq.Less(4, 5))

	fmt.Println("Copied Sequence 2:", seq2)
	fmt.Println("Sorted Copied Sequence:", seq2.String())
	fmt.Println(seq2.Len())


	var stringSeq StringSequence = []string{"apple", "cherry", "orange", "lime", "date", "elderberry", "fig"}
	fmt.Println("Length of stringSeq: ", len(stringSeq))
	fmt.Println("Second element of stringSeq: ", stringSeq[1])

	s4 := &seq2
	fmt.Println("s4: ", s4)
	var s5 int = 10
	modify(&s5)
	fmt.Println("s5: ", s5)

	r, e := Finder(stringSeq, "cherry")
	fmt.Println("r: ", r)
	fmt.Println("e: ", e)




}


























