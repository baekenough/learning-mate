package day01

import (
	"fmt"
)

/*
## 5. [**x만큼 간격이 있는 n개의 숫자**](https://school.programmers.co.kr/learn/courses/30/lessons/12954?language=go)

### **문제 설명**

함수 solution은 정수 x와 자연수 n을 입력 받아, x부터 시작해 x씩 증가하는 숫자를 n개 지니는 리스트를 리턴해야 합니다. 다음 제한 조건을 보고, 조건을 만족하는 함수, solution을 완성해주세요.

### **제한 조건**

- x는 -10000000 이상, 10000000 이하인 정수입니다.
- n은 1000 이하인 자연수입니다.
*/

func solution05(x int, n int) []int64 {
	var array []int64
	for i := 0; i < n; i++ {
		count := (i + 1) * x
		array = append(array, int64(count))
	}
	return array
}

func main() {
	fmt.Println(solution05(-4, 2))

}
