package main

func main() {
	//**과제 1:** `int`, `float64`, `string`, `bool` 타입을 사용하여 각각 나이, 원주율, 이름, 학습 완료 여부를 나타내는 변수를 선언하고 출력하는 코드를 작성하세요.
	var age int = 30
	var pi float64 = 3.1415926
	var Name string = "Sangyi Baek"
	var complete bool = true
	println(age, pi, Name, complete)

	//**과제 2:** `const`를 사용하여 `Pi`라는 이름의 상수를 `3.14159`로 선언한 뒤, 이 값을 변경하려고 시도할 때 발생하는 컴파일 에러를 확인하고 그 이유를 주석으로 설명하세요.
	const Pi = 3.14159
	Pi = 3.14 //.\homework-01.go:13:2: cannot assign to Pi (neither addressable nor a map index expression)
	//컴파일 타임에 상수의 값이 고정되어있기에, 컴파일러가 이를 체크하여 빌드 자체가 실패한다.
	//만약 값을 변경하고 싶다면 const가 아닌 var로 선언해야 한다.

	//**과제 3:** 두 `int` 변수 `a=20`, `b=5`를 사용하여 덧셈, 뺄셈, 곱셈, 나눗셈, 나머지 연산 결과를 출력하세요. 또한, `a > b` 와 `a == b` 비교 연산의 결과도 출력하세요.
	var a int = 20
	var b int = 5
	println(a+b, a-b, a*b, a/b, a%b)
	println(a > b, a == b)
}
