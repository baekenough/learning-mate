package day01

import (
	"fmt"
	"slices"
	"strconv"
)

/*
## 2. [**자연수 뒤집어 배열로 만들기**](https://school.programmers.co.kr/learn/courses/30/lessons/12932)

### **문제 설명**

자연수 n을 뒤집어 각 자리 숫자를 원소로 가지는 배열 형태로 리턴해주세요. 예를들어 n이 12345이면 [5,4,3,2,1]을 리턴합니다.

### 제한 조건

- n은 10,000,000,000이하인 자연수입니다.
*/

func solution02(n int64) []int {
	convStr := strconv.Itoa(int(n))
	var intArray []int
	for i := range convStr {
		intConvStr := int(convStr[i]) - int('0')
		intArray = append(intArray, intConvStr)
	}
	slices.Reverse(intArray) // array를 뒤집는다
	return intArray
}

func main() {
	fmt.Println(solution02(12345))

}
