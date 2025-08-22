package day01

import (
	"fmt"
)

/*
## 3. [짝수와 홀수](https://school.programmers.co.kr/learn/courses/30/lessons/12937?language=go)

정수 num이 짝수일 경우 "Even"을 반환하고 홀수인 경우 "Odd"를 반환하는 함수, solution을 완성해주세요.

### 제한 조건

- num은 int 범위의 정수입니다.
- 0은 짝수입니다.
*/

func solution03(num int) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	fmt.Println(solution03(12345))

}
