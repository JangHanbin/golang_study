package main

import (
	"fmt"
	"rsc.io/quote"
)

const (
	i int = 18
)

func test() {
	i := 10
	fmt.Println(i)
}
func main() {
	fmt.Println(quote.Go())
	test()
	fmt.Println(i)
	fmt.Println("숨겨왔던 나의~ " +
		"수줍은 마음 모두~")
	var str string = "Hello Go lang"
	fmt.Println(str)

	var k int = 20

	if true {
		fmt.Println("Good")
		fmt.Println(k)
	}

	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}

	s := []int{0, 1, 2, 3, 4, 5}
	d := s
	println(&s)
	println(&(d))
	fmt.Println(s)
	s = s[2:5] // 2, 3, 4
	fmt.Println(s)
	s = s[1:]      // 3, 4
	fmt.Println(s) // 3, 4 출력
	fmt.Println(d)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)
	s = append(s, 80)

	//s = append(s, 80)

	fmt.Println(s)
	println(&s)
	println(&s[1])
	println(*&s[1])
	fmt.Println(cap(s))

	println(GetMusic("Alicia Keys"))
	//song := GetMusic("Alicia Keys")
	//println(song)
	// Due to experiment, slice maybe consists of like a vector in c++
}
