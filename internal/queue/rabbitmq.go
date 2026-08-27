package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient interface {
	Publish(queueName string, payload interface{}) error
	Consume(queueName string, handler func(body []byte) error) error
	IsAvailable() bool
	Close()
}

type rabbitMQClient struct {
	url       string
	conn      *amqp.Connection
	ch        *amqp.Channel
	available bool
	mu        sync.RWMutex
}

func NewRabbitMQClient(rawURL string) RabbitMQClient {
	r := &rabbitMQClient{url: rawURL}
	r.connect()
	return r
}

func (r *rabbitMQClient) connect() {
	r.mu.Lock()
	defer r.mu.Unlock()

	conn, err := amqp.Dial(r.url)
	if err != nil {
		log.Printf("⚠️ [RabbitMQ Notice]: Không thể đăng nhập RabbitMQ (%v). Chuyển sang Fallback Synchronous Processing.", err)
		r.available = false
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		log.Printf("⚠️ [RabbitMQ Notice]: Không thể mở channel RabbitMQ. Chuyển sang Fallback Synchronous Processing.")
		r.available = false
		return
	}

	r.conn = conn
	r.ch = ch
	r.available = true
	log.Println("⚡ Kết nối RabbitMQ Message Queue thành công!")
}

func (r *rabbitMQClient) IsAvailable() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.available
}

func (r *rabbitMQClient) Publish(queueName string, payload interface{}) error {
	if !r.IsAvailable() {
		return fmt.Errorf("rabbitmq offline")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", queueName, err)
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = r.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         bytesData,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish to queue %s: %w", queueName, err)
	}

	log.Printf("📥 [RabbitMQ Engine]: Đã ném tin nhắn vào Queue '%s' thành công!", queueName)
	return nil
}

func (r *rabbitMQClient) Consume(queueName string, handler func(body []byte) error) error {
	if !r.IsAvailable() {
		return fmt.Errorf("rabbitmq offline")
	}

	r.mu.Lock()
	_, err := r.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.mu.Unlock()
		return fmt.Errorf("failed to declare queue %s: %w", queueName, err)
	}

	msgs, err := r.ch.Consume(
		queueName,
		"",    // consumer
		false, // autoAck (manual ack for safety)
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)
	r.mu.Unlock()

	if err != nil {
		return fmt.Errorf("failed to consume from queue %s: %w", queueName, err)
	}

	go func() {
		log.Printf("👷 [RabbitMQ Worker]: Đăng ký lắng nghe Queue '%s' thành công!", queueName)
		for d := range msgs {
			if err := handler(d.Body); err == nil {
				_ = d.Ack(false)
			} else {
				log.Printf("❌ [RabbitMQ Worker Error]: Xử lý message từ '%s' thất bại: %v", queueName, err)
				_ = d.Nack(false, true) // requeue
			}
		}
	}()

	return nil
}

func (r *rabbitMQClient) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
	r.available = false
}
