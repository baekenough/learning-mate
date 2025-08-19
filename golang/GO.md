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
- day005. 문법 - methods, interfaces(https://baekenough.oopy.io/2537ecd4-3d73-80c5-bec9-d1e262739561)
- day006. 문법 - goroutines, channels(https://baekenough.oopy.io/2547ecd4-3d73-80e8-b94e-dea76dd12d73)

## 4. 과제(write by Teacher)
- **day006: goroutines와 channels 실습**
  1. **Part A - 기본 goroutines**: `say` 함수를 고루틴으로 3번 호출하되, 각각 다른 문자열("Hello", "World", "Go")을 전달합니다.
  2. **Part B - channels 활용**: `fibonacci` 함수를 만들어 채널을 통해 피보나치 수열의 첫 10개 값을 전송하고, `main` 함수에서 받아서 출력합니다.
  3. `say` 함수는 문자열을 받아 1초 간격으로 3번 출력합니다.
  4. `fibonacci` 함수는 `chan int` 타입의 채널을 받아 피보나치 수열을 전송합니다.
  5. 과제 코드는 `golang/homeworks/homework-06.go` 파일을 생성하여 그 안에 작성해주세요.

    
## 5. 과제 피드백(write by Homework Coach)
- **Part A (기본 goroutines)**
    - **칭찬할 점:** `sync.WaitGroup`을 사용하여 모든 고루틴이 종료될 때까지 기다리는 로직을 정확하게 구현했습니다. 이는 동시성 프로그래밍에서 매우 중요한 부분입니다.
    - **개선할 점:** `say` 함수의 요구사항은 "문자열을 받아 1초 간격으로 3번 출력"하는 것이었지만, 현재 코드는 한 번만 출력하고 있습니다. `for` 루프와 `time.Sleep`을 사용하여 요구사항을 만족시킬 수 있습니다.

- **Part B (channels 활용)**
    - **칭찬할 점:** 피보나치 수열을 생성하여 채널로 보내고, `range`를 이용해 채널이 닫힐 때까지 값을 수신하는 흐름을 완벽하게 구현했습니다. `close` 함수를 사용하여 채널을 닫아주는 부분도 정확합니다.

**총평:**
`goroutine`과 `channel`의 기본적인 사용법과 동시성 제어의 필요성을 잘 이해하고 계십니다. 특히 `WaitGroup`의 사용법과 채널을 활용한 데이터 통신 방법은 매우 훌륭합니다. `say` 함수의 세부 요구사항을 놓친 점이 아쉽지만, 전반적으로 Go의 동시성 모델을 성공적으로 적용하셨습니다.
    