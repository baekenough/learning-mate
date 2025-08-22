package day01

import (
	"fmt"
	"strconv"
)

/*
## 1. [**자릿수 더하기**](https://school.programmers.co.kr/learn/courses/30/lessons/12931?language=go)

### **문제 설명**

자연수 N이 주어지면, N의 각 자릿수의 합을 구해서 return 하는 solution 함수를 만들어 주세요.

예를들어 N = 123이면 1 + 2 + 3 = 6을 return 하면 됩니다.

### 제한사항

- N의 범위 : 100,000,000 이하의 자연수
*/

func solution01(n int) int {
	answer := 0
	numStr := strconv.Itoa(n) // 하나씩 슬라이스하기 위해 string으로 컨버팅한다

	for i := range numStr { // string 타입은 그 길이로 lenth를 제공한다
		temp := 0
		temp = int(numStr[i]) - int('0') // 아스키코드값 1부터 시작하기 위함
		answer += temp
	}

	return answer
}

func main() {
	fmt.Println(solution01(123))

}
