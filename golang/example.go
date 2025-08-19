package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type Email struct {
	Address string
}

func (e Email) Send(message string) error {
	fmt.Printf("'%s' 주소로 이메일 전송: %s\n", e.Address, message)
	return nil
}

type Sms struct {
	PhoneNumber string
}

func (s Sms) Send(message string) error {
	fmt.Printf("'%s' 번호로 문자 전송: %s\n", s.PhoneNumber, message)
	return nil
}

func SendNotification(n Notifier, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Println("전송 실패:", err)
	}
}

func main() {
	emailNotifier := Email{Address: "test@example.com"}
	smsNotifier := Sms{PhoneNumber: "010-1234-5678"}

	fmt.Println("--- 개별 알림 전송 ---")
	SendNotification(emailNotifier, "안녕하세요! 인터페이스 예제입니다. (이메일)")
	SendNotification(smsNotifier, "안녕하세요! 인터페이스 예제입니다. (SMS)")

	fmt.Println("\n--- 여러 알림 한번에 보내기 ---")
	notifiers := []Notifier{
		Email{Address: "user1@example.com"},
		Sms{PhoneNumber: "010-1111-1111"},
		Email{Address: "user2@example.com"},
		Sms{PhoneNumber: "010-2222-2222"},
	}

	for _, notifier := range notifiers {
		SendNotification(notifier, "전체 공지입니다.")
	}
}
