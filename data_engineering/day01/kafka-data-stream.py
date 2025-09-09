from kafka import KafkaProducer, KafkaConsumer
import json

try:
    producer = KafkaProducer(
        bootstrap_servers=['43.201.107.204:9092'],
        value_serializer=lambda x: json.dumps(x).encode('utf-8'),
        request_timeout_ms=10000,
        retries=3
    )
    print("Kafka producer created successfully")
except Exception as e:
    print(f"Failed to create producer: {e}")

test_data = {
    "email_id": "email_001",
    "to": "user@example.com",
    "subject": "Welcome to our service",
    "content": "Thank you for signing up!",
    "timestamp": "2025-01-09T10:30:00"}

def send_email():
    try:
        future = producer.send('email.send', test_data)
        record_metadata = future.get(timeout=10)
        print(f"Message sent successfully to {record_metadata.topic} partition {record_metadata.partition}")
        producer.flush()
    except Exception as e:
        print(f"Failed to send message: {e}")
    # ubuntu@ip-172-31-36-85:~/kafka_2.13-4.0.0$ ./bin/kafka-console-consumer.sh --bootstrap-server 43.201.107.204:9092 --topic email.send --from-beginning
    # {"email_id": "email_001", "to": "user@example.com", "subject": "Welcome to our service", "content": "Thank you for signing up!", "timestamp": "2025-01-09T10:30:00"}
    # {"email_id": "email_001", "to": "user@example.com", "subject": "Welcome to our service", "content": "Thank you for signing up!", "timestamp": "2025-01-09T10:30:00"}
    # {"email_id": "email_001", "to": "user@example.com", "subject": "Welcome to our service", "content": "Thank you for signing up!", "timestamp": "2025-01-09T10:30:00"}

# if 'producer' in locals():
#     send_email()
# else:
#     print("Producer not available")

consumer = KafkaConsumer(
    'email.send',
    bootstrap_servers=['43.201.107.204:9092'],
    value_deserializer=lambda x: json.loads(x.decode('utf-8')),
    auto_offset_reset='earliest',
    group_id='email-processor')

print("Consumer started...")
for message in consumer:
    email_data = message.value
    print(f"Received email: {email_data['subject']} to {email_data['to']}")
    # Received email: Welcome to our service to user@example.com
    # Received email: Welcome to our service to user@example.com
    # Received email: Welcome to our service to user@example.com