## 1. 학습목표
- 아래 link의 JD에 기재된 기술 스택을 학습
  - https://www.wanted.co.kr/wd/297237
- AWS를 적극적으로 활용한다
- 주 언어는 python 3.12.9(conda activate data-engineering)
- 부 언어로 golang 1.25를 사용한다

## 2. 참고자료
- **Apache Kafka 공식 문서**: https://kafka.apache.org/documentation/
- **Apache Airflow 공식 문서**: https://airflow.apache.org/docs/
- **AWS Big Data 서비스 가이드**: https://aws.amazon.com/big-data/
- **PySpark 공식 문서**: https://spark.apache.org/docs/latest/api/python/
- **Kafka Python Client**: https://kafka-python.readthedocs.io/
- **AWS Glue 가이드**: https://docs.aws.amazon.com/glue/
- **데이터 엔지니어링 기본서**: https://www.oreilly.com/library/view/designing-data-intensive/9781449373320/

## 3. 학습내용
### 1단계: Apache Kafka & Airflow 기초
**Apache Kafka**는 고성능 분산 스트리밍 플랫폼으로, 실시간 데이터 파이프라인의 핵심입니다. LinkedIn에서 개발된 이 기술은 초당 수백만 개의 메시지를 처리할 수 있어 대규모 시스템의 필수 기술입니다.

**Apache Airflow**는 데이터 워크플로우를 코드로 작성하고 스케줄링할 수 있는 도구입니다. Python으로 DAG(Directed Acyclic Graph)를 정의하여 복잡한 데이터 파이프라인을 관리할 수 있습니다.

이 두 기술을 먼저 학습하는 이유:
1. 데이터 엔지니어링의 핵심인 데이터 이동과 처리 워크플로우를 이해할 수 있습니다
2. 실무에서 가장 많이 사용되는 오픈소스 기술입니다
3. 다른 도구들(Spark, AWS 서비스)과 연동할 때 기반이 되는 기술입니다

## 4. 과제(write by Teacher)
### Day 1: Apache Kafka 기초
1. **Kafka 로컬 환경 구축**: Docker를 사용해 Kafka 클러스터 구축하고 topic 생성/메시지 송수신 테스트
2. **Python Kafka Producer/Consumer 작성**: kafka-python 라이브러리를 사용해 간단한 데이터 스트림 구현

### Day 1 과제 피드백 (write by Homework Coach)

#### 과제 1: Kafka 로컬 환경 구축 ✅ **완료**
- EC2에 Kafka 서버 구축 완료 (43.201.107.204:9092)
- `email.send` 토픽 생성 및 메시지 송수신 테스트 성공
- Console producer/consumer로 동작 확인 완료

#### 과제 2: Python Kafka Producer/Consumer 작성 ✅ **완료**
- `kafka-data-stream.py`에서 Producer와 Consumer 모두 구현
- Producer: JSON 형태의 이메일 데이터 전송 성공
- Consumer: 메시지 수신 및 처리 로직 구현 완료
- Consumer Group(`email-processor`) 및 offset 관리 이해 완료

**우수한 점:**
1. 예외 처리가 잘 구현되어 있음
2. 실제 데이터 구조(이메일)를 사용한 실용적 예시
3. Producer와 Consumer가 하나의 파일에 통합되어 테스트 용이
4. Consumer Group과 offset 개념 정확히 이해

**학습 성과:**
- Kafka의 핵심 개념(Topic, Producer, Consumer, Offset) 습득
- 실시간 데이터 스트리밍 파이프라인 구현 경험

### Day 2: Apache Airflow 기초
3. **Airflow 설치 및 첫 DAG 작성**: 로컬에 Airflow 설치하고 Hello World DAG 작성하여 실행
4. **Kafka + Airflow 통합**: Airflow DAG에서 Kafka topic의 메시지를 읽어 처리하는 파이프라인 구현