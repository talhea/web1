package main

import (
	"fmt"
	"os"
)

func main() {
	// 입력 받기
	var rows int
	fmt.Print("출력할 별의 라인 수를 입력하시오: ")
	_, err := fmt.Scan(&rows)
	if err != nil {
		fmt.Println("숫자가 아닙니다")
		os.Exit(1)
	}

	// 별 출력
	// for i := 0; i < rows; i++ {
	// 	for j := rows - i; j > 0; j-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := rows - i; k > 0; k-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < rows; i++ {
	// 	for j := rows - i - 1; j > 0; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 2*i + 1; k > 0; k-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 2*(rows-i) - 1; k > 0; k-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// half := rows / 2
	// for i := 0; i < rows; i++ {
	// 	repeat := i - half
	// 	if repeat < 0 {
	// 		repeat *= -1
	// 	}

	// ver. 1
	// 	for j := 0; j < repeat; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := rows - 2*repeat; k > 0; k-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// ver. 2
	totalRows := rows*2 - 1 // 전체 라인 수 == 라인 당 자리 수 == 라인 최대 별 수
	half := totalRows / 2
	var absNumber int
	var k int
	for number := half * -1; number <= half; number++ { // 총 진행 라인
		if number < 0 { // 현 진행 라인을 이용한 절대값을 계산해서, 라인 당 출력할 총 자리 수와 그 라인에 찍힐 빈칸들의 수 확인에 이용
			absNumber = number * -1
		} else {
			absNumber = number
		}
		k = 0 // absNumber를 이용해 빈공간 또는 "*" 출력 판별에 이용

		for j := 0; j < totalRows-absNumber; j++ { // 라인 당 출력해야 할 자리 수
			if k < absNumber { // 빈칸 혹은 별("*") 출력 판별
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
			k++
		}

		fmt.Println()
	}
	fmt.Println("ㅎㅎ")
}
