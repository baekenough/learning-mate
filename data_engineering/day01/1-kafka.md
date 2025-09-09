# kafka는 EC2에 생성하였음
43.201.107.204:9092

# 카프카 실행하기(포그라운드)
./bin/kafka-server-start.sh config/server.properties

# 카프카 실행하기(백그라운드)
./bin/kafka-server-start.sh -daemon config/server.properties

# 카프카 종료하기
./bin/kafka-server-stop.sh
# 실행여부 확인
sudo lsof -i:9092

# 토픽 생성
./bin/kafka-topics.sh --bootstrap-server 43.201.107.204:9092 --create --topic email.send

# 토픽 리스트 조회
./bin/kafka-topics.sh --bootstrap-server 43.201.107.204:9092 --list

# email.send 토픽 접근하기
./bin/kafka-console-producer.sh --bootstrap-server 43.201.107.204:9092 --topic email.send
## 입력값
hello1
hello2
hello3

# email.send에 쌓인 메세지 전부 조회하기
./bin/kafka-console-consumer.sh --bootstrap-server 43.201.107.204:9092 --topic email.send --from-beginning
## 출력값
hello1
hello2
hello3

# email.send 토픽에 hello4 추가하기
./bin/kafka-console-producer.sh --bootstrap-server 43.201.107.204:9092 --topic email.send
## 입력값
hello4

# 다시 읽어서 메세지 리스트 확인하기
## 출력
hello1
hello2
hello3
hello4