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
- day007. 문법 - select, sync.Mutex(https://baekenough.oopy.io/2557ecd4-3d73-80fa-bc7a-fcc40921c822)

## 4. 과제(write by Teacher)
- **day007: 동시성 웹 스크레이퍼 시뮬레이터**
  여러 웹사이트를 동시에 크롤링하고, 그 결과를 안전하게 저장하는 간단한 시뮬레이터를 만들어 보겠습니다. 이 과제를 통해 `sync.Mutex`와 `select`를 함께 사용하는 실전적인 방법을 익힐 수 있습니다.

  **요구사항:**

  1.  **`CrawlResult` 구조체 정의:**
      -   `url` (string): 크롤링한 URL
      -   `body` (string): 크롤링 결과 (간단한 문자열로 대체)

  2.  **`SafeCrawlStore` 구조체 정의:**
      -   `mu` (`sync.Mutex`): 저장소를 보호할 뮤텍스
      -   `results` (`map[string]string`): URL을 키로, 크롤링 결과(body)를 값으로 저장하는 맵

  3.  **`worker` 함수 구현:**
      -   `id` (int), `jobs` (`<-chan string`), `store` (`*SafeCrawlStore`), `wg` (`*sync.WaitGroup`)을 인자로 받습니다.
      -   `jobs` 채널로부터 URL을 받아 "Crawled [URL]" 라는 문자열을 생성합니다.
      -   `store`의 메소드를 사용해 결과를 안전하게 저장합니다.
      -   작업이 끝나면 `WaitGroup`에 완료 신호를 보냅니다.

  4.  **`main` 함수 로직:**
      -   5개의 URL을 담은 슬라이스를 생성합니다.
      -   `jobs` 채널을 생성하고, 고루틴을 사용해 URL들을 `jobs` 채널에 보냅니다. 모든 URL을 보낸 후 채널을 닫습니다.
      -   `SafeCrawlStore` 인스턴스를 생성합니다.
      -   3개의 `worker` 고루틴을 실행합니다.
      -   `time.NewTicker`를 사용해 50밀리초마다 한 번씩 `store`에 저장된 결과의 개수를 출력합니다.
      -   `select` 문을 사용합니다.
          -   `case` 1: Ticker가 울릴 때마다 결과 개수를 출력합니다.
          -   `case` 2: 모든 `worker`의 작업이 완료되면 (WaitGroup을 활용) Ticker를 멈추고 "All workers done." 메시지를 출력한 뒤, 루프를 종료합니다.
      -   최종적으로 `store`에 저장된 모든 결과를 출력합니다.

  5.  **파일 생성:**
      -   과제 코드는 `golang/homeworks/homework-07.go` 파일을 생성하여 그 안에 작성해주세요.

- **day006: goroutines와 channels 실습**
  1. **Part A - 기본 goroutines**: `say` 함수를 고루틴으로 3번 호출하되, 각각 다른 문자열("Hello", "World", "Go")을 전달합니다.
  2. **Part B - channels 활용**: `fibonacci` 함수를 만들어 채널을 통해 피보나치 수열의 첫 10개 값을 전송하고, `main` 함수에서 받아서 출력합니다.
  3. `say` 함수는 문자열을 받아 1초 간격으로 3번 출력합니다.
  4. `fibonacci` 함수는 `chan int` 타입의 채널을 받아 피보나치 수열을 전송합니다.
  5. 과제 코드는 `golang/homeworks/homework-06.go` 파일을 생성하여 그 안에 작성해주세요.

    
## 5. 과제 피드백(write by Homework Coach)
- **day007: 동시성 웹 스크레이퍼 시뮬레이터 (재평가)**
    - **칭찬할 점**: 이전 피드백을 정말 훌륭하게 반영하셨습니다! Deadlock 문제, `wg.Done()` 사용, `select`에서 `done` 채널을 기다리는 패턴, `wg.Add(1)` 등 핵심적인 부분들을 모두 정확하게 수정하셨습니다. 정말 큰 발전을 이루셨습니다!
    - **개선할 점**: 이제 개별적인 기능은 모두 올바르게 동작하니, 전체적인 로직을 완성할 차례입니다. 현재 코드는 3개의 작업을 처리하면 프로그램이 종료될 가능성이 높고, `worker`가 작업을 하나만 처리하고 사라지는 점이 아쉽습니다.
        1.  **`worker`가 계속 일하게 만들기**: `worker` 함수가 URL 하나를 처리하고 종료되는 대신, `for range` 루프를 사용해 `jobs` 채널이 닫힐 때까지 계속해서 작업을 가져와 처리하도록 수정해 보세요. 이렇게 하면 3명의 `worker`가 5개의 모든 작업을 나눠서 처리하게 됩니다.
        2.  **`main` 함수의 구조**: `WaitGroup`의 카운터를 작업의 총개수(5개)로 설정하고, `worker`들을 먼저 실행시킨 후, 모든 작업을 `jobs` 채널에 보내는 구조로 변경하는 것이 좋습니다. `select` 루프는 이 모든 설정이 끝난 뒤에 한 번만 실행하여 모든 작업이 완료되기를 기다려야 합니다.

**총평**:
이제 Go 동시성의 핵심적인 기술 요소들은 모두 정확하게 사용하고 계십니다. 마지막으로 "워커 풀(Worker Pool)"의 전체적인 구조를 완성하는 단계만 남았습니다. 제안된 개선점을 적용하시면, 여러 작업을 효율적으로 동시에 처리하는 실용적인 동시성 프로그램을 완성하실 수 있을 겁니다. 거의 다 왔습니다!

- **day007: 동시성 웹 스크레이퍼 시뮬레이터**
    - **칭찬할 점**: `sync.Mutex`를 사용해 공유 데이터(크롤링 결과 저장소)를 보호하려는 시도와 `select` 문을 통해 여러 채널을 동시에 기다리려는 구조를 잡으신 점이 좋습니다. 동시성 프로그래밍의 핵심 요소를 잘 파악하고 계십니다.
    - **개선할 점**:
        1.  **`worker` 함수의 `wg.Wait()`**: `worker`는 작업을 마친 후 `wg.Done()`을 호출하여 자신이 끝났음을 알려야 합니다. `wg.Wait()`는 모든 작업이 끝나길 기다리는 함수이므로, `main` 함수에서 사용해야 합니다. 현재 코드는 `worker`들이 서로를 무한정 기다리는 deadlock 상태에 빠지게 됩니다.
        2.  **`main` 함수의 `jobs <- target[i]`**: 버퍼가 없는 채널에 데이터를 보내면, 받는 쪽이 준비될 때까지 코드가 멈춥니다(block). 현재 코드는 `worker`를 실행하기 전에 `jobs` 채널에 데이터를 보내므로 첫 번째 데이터 전송에서 바로 deadlock이 발생합니다. `jobs` 채널을 버퍼 채널(`make(chan string, len(target))`)로 만들면 이 문제를 해결할 수 있습니다.
        3.  **`select` 문의 `case <-worker`**: 이 부분이 "case2" 문제의 핵심입니다. `worker`는 함수이지 채널이 아니므로 `select` 문에서 받을 수 없습니다. 모든 `worker`가 완료되었는지 확인하려면 `WaitGroup`을 사용해야 합니다. 올바른 방법은 다음과 같습니다.
            - 별도의 `done` 채널을 만듭니다: `done := make(chan struct{})`
            - `wg.Wait()`를 실행하고, 완료되면 `done` 채널을 닫는 고루틴을 하나 실행합니다: `go func() { wg.Wait(); close(done) }()`
            - `select` 문에서 `case <-done:` 으로 신호를 기다립니다.
        4.  **기타**:
            - `wg.Add(i)` 대신 `wg.Add(1)`을 사용해야 의도한 대로 동작합니다.
            - `tick` 케이스에서 `len(target)`이 아닌 `len(instance.result)`를 출력해야 현재까지 크롤링된 결과의 개수를 볼 수 있습니다.

**총평**:
동시성 프로그래밍에서 흔히 겪는 deadlock과 `WaitGroup` 사용법에 대한 좋은 학습 경험이 되셨을 겁니다. 특히 `select`와 `WaitGroup`을 함께 사용하는 패턴은 중요하므로, 제안된 수정 방법을 적용하여 코드를 완성해 보시는 것을 추천합니다. 이 부분을 명확히 이해하시면 Go의 동시성 활용 능력이 크게 향상될 것입니다.

- **Part A (기본 goroutines)**
    - **칭찬할 점:** `sync.WaitGroup`을 사용하여 모든 고루틴이 종료될 때까지 기다리는 로직을 정확하게 구현했습니다. 이는 동시성 프로그래밍에서 매우 중요한 부분입니다.
    - **개선할 점:** `say` 함수의 요구사항은 "문자열을 받아 1초 간격으로 3번 출력"하는 것이었지만, 현재 코드는 한 번만 출력하고 있습니다. `for` 루프와 `time.Sleep`을 사용하여 요구사항을 만족시킬 수 있습니다.

- **Part B (channels 활용)**
    - **칭찬할 점:** 피보나치 수열을 생성하여 채널로 보내고, `range`를 이용해 채널이 닫힐 때까지 값을 수신하는 흐름을 완벽하게 구현했습니다. `close` 함수를 사용하여 채널을 닫아주는 부분도 정확합니다.

**총평:**
`goroutine`과 `channel`의 기본적인 사용법과 동시성 제어의 필요성을 잘 이해하고 계십니다. 특히 `WaitGroup`의 사용법과 채널을 활용한 데이터 통신 방법은 매우 훌륭합니다. `say` 함수의 세부 요구사항을 놓친 점이 아쉽지만, 전반적으로 Go의 동시성 모델을 성공적으로 적용하셨습니다.
    