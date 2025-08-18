package main

func getScore(score int) string {
	if score >= 90 {
		return "A"
	}
	if score >= 80 {
		return "B"
	}
	if score >= 70 {
		return "C"
	}
	return "F"
}

func getEvenNumber() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
}

func getDayOfWeek(day string) string {
	switch day {
	case "Monday":
		return "월요일"
	case "Tuesday":
		return "화요일"
	default:
		return "기타"
	}
}

func main() {
	//**과제 1:** `if`문을 사용하여, 주어진 정수형 변수 `score`가 90점 이상이면 "A", 80점 이상이면 "B", 70점 이상이면 "C", 그 외에는 "F"를 출력하는 코드를 작성하세요.
	println(getScore(90))
	//**과제 2:** `for`문을 사용하여, 1부터 10까지의 숫자 중에서 짝수만 출력하는 코드를 작성하세요.
	getEvenNumber()
	//**과제 3:** `switch`문을 사용하여, 주어진 문자열 변수 `day`가 "Monday"면 "월요일", "Tuesday"면 "화요일", 그 외에는 "기타"를 출력하는 코드를 작성하세요.
	println(getDayOfWeek("Monday"))
}
