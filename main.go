package main

import (
	"fmt"
	"strings"
)

var exV1 string = "asd"

// ERRROR : exV2 := "asd"

func lean1() {
	//상수 재활당 금지
	const name string = "bh"
	//변수 선언
	var v1 string = "value1"
	//축약형 축약형은 외부에서 사용 할 수 없음
	v2 := "v2"

	fmt.Println(exV1)
	fmt.Println(name)
	fmt.Println(v1)
	fmt.Println(v2)

}

func multply(a, b int) int {
	return a * b
}

func lendAddUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repreatMe(words ...string) {
	fmt.Println(words)
}

func functionLean() {
	fmt.Println(multply(3, 5))
	totalLeng, upperName := lendAddUpper("김병희")
	totalLeng2, _ := lendAddUpper("김병희")
	fmt.Println(totalLeng, upperName)
	fmt.Println(totalLeng2)

	repreatMe("hi", "dep", "taa")

}

func lenAndUpper(name string) (length int, upper string) {
	defer fmt.Println("I', done") //함수를 끝내고 실행함
	length = len(name)
	//Error length := len(name) 상위에서 이미 선언을 해주었기 때문에 에러가 남
	upper = strings.ToUpper(name)
	return
}

//반복문
func superAdd(numbers ...int) int {
	total := 0
	// Range 는 Index값을 출력함
	//
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("for Loop", numbers[i])
	}

	for _, number := range numbers {
		total += number
	}

	return total

}

func carIDrink(age int) bool {

	//IF 문에서 변수를 쓰는 방식 Express variable
	//koranAge := age + 2;
	/*
		if koranAge := age + 2; koranAge < 18 {
			return false
		}
		return true
	*/
	switch kareang := age + 2; kareang {
	case 18:
		fmt.Println("Iam 18")

	}

	switch {
	case age < 18:
		return false
	case age >= 18:
		return true
	}
	return false

}

func pointer() {
	/*
		a := 2
		//b := a //값 복사
		b := a
		a = 10

		fmt.Println(&a, &b) //메모리주소를 보는법
	*/

	c := 2
	d := &c
	//c = 3
	*d = 10
	//* 포인트는 주소를 훍어보는것
	fmt.Println(*d, c)

}

//Array
func arraySlice() {
	/* 크기있는 array */
	names :=
		[5]string{"gi", "hi", "sdsf", "asdsad"}
	names[4] = "fs"
	su := [5]string{}
	su[1] = "a"

	// 크기없는 slice
	namseSlice := []string{"nuco", "lynn", "da"}
	namseSlice = append(namseSlice, "can")

}

func mapsTest() {
	kims := map[string]string{"name": "kim", "age": "12"}
	fmt.Println(kims)

	for _, value := range kims {
		fmt.Println(value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	foode := []string{"sada", "sad", "sad"}
	kims := person{"kims", 24, foode}
	kims2 := person{name: "kims", age: 44, favFood: foode}
	fmt.Println(kims)
	fmt.Println(kims2)

}
