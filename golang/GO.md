## 1. 언어
- golang(1.25)

## 2. 참고자료
- A Tour to Go(https://go.dev/tour/list)
- Go packages(https://pkg.go.dev/)
- effective-go(https://github.com/golangkorea/effective-go?tab=readme-ov-file)
- 가장 빨리 만나는 Go 언어(https://pyrasis.com/go.html)

## 3. 학습내용
- day001. 문법 - 타입, 상수, 연산자(https://baekenough.oopy.io/24e7ecd4-3d73-8027-a24b-c9c3bcfcb653)
- day002. 문법 - 제어문(if, for, switch, defer)(https://baekenough.oopy.io/2517ecd4-3d73-80e6-abd5-cda3844ab937)
- day003. 문법 - pointer, struct, array, slice(https://baekenough.oopy.io/2517ecd4-3d73-80c5-a309-ef2e4142bd7b)
- day004. 문법 - maps, functions(https://baekenough.oopy.io/2527ecd4-3d73-80f4-a46d-cd364f5b8737)
- day005. 문법 - methods, interfaces(https://go.dev/tour/methods/1)

## 4. 과제(write by Teacher)
## 4. 과제(write by Teacher)
- **day005: 도형의 넓이 계산하기**
  1. `Shape`라는 이름의 인터페이스를 정의하고, 이 인터페이스는 `area() float64` 메서드를 가져야 합니다.
  2. `Rectangle` (사각형) 구조체를 정의하세요. `width`와 `height` 필드(모두 `float64`)를 가져야 합니다.
  3. `Circle` (원) 구조체를 정의하세요. `radius` 필드(`float64`)를 가져야 합니다.
  4. `Rectangle`과 `Circle` 타입이 `Shape` 인터페이스를 만족하도록 `area()` 메서드를 각각 구현하세요. (원의 넓이 = π * 반지름 * 반지름)
  5. `main` 함수에서 `Rectangle`과 `Circle` 인스턴스를 만들고 각 도형의 넓이를 출력하세요.
  6. 과제 코드는 `golang/homeworks/homework-05.go` 파일을 생성하여 그 안에 작성해주세요.

    
## 5. 과제 피드백(write by Homework Coach)
- **day004:**
  - **잘한 점:** `map`과 `for...range`를 사용하여 요구사항의 핵심 기능을 정확히 구현했으며, `printScores` 함수로 기능을 분리하여 코드 구조를 잘 설계했습니다.
  - **개선 제안:** 과제는 `map[string]int`를 요구했으나 `map[string]Score` 구조체를 사용했습니다. 의도와는 달랐지만 구조체에 대한 이해도를 보여주었습니다. 더 간결한 코드를 위해 `int` 타입을 직접 사용하는 것을 권장합니다. 또한, `for...range` 루프에서 변수 이름을 `name`, `score` 등으로 더 명확하게 지정하면 가독성이 향상될 것입니다.
  - **총평:** 요구사항을 성공적으로 완수했으며, `maps`와 `functions`의 개념을 잘 이해하고 활용했습니다.
    