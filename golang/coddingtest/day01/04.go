package day01

import (
	"fmt"
)

/*
## 4. [평균 구하기](https://school.programmers.co.kr/learn/courses/30/lessons/12944?language=go)

### **문제 설명**

정수를 담고 있는 배열 arr의 평균값을 return하는 함수, solution을 완성해보세요.

### **제한사항**

- arr은 길이 1 이상, 100 이하인 배열입니다.
- arr의 원소는 -10,000 이상 10,000 이하인 정수입니다.
*/

func solution04(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	average := float64(sum) / float64(len(arr))
	return average
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(solution04(arr))

}
