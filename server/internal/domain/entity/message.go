package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	ReceiverID uuid.UUID
	SenderID   uuid.UUID
	Content    string
	SentAt     time.Time
	ReadAt     time.Time
}

func NewMessage(content string, senderID, receiverID uuid.UUID) *Message {
	return &Message{
		ID:         uuid.New(),
		Content:    content,
		SenderID:   senderID,
		ReceiverID: receiverID,
		SentAt:     time.Now(),
	}
}

func (msg *Message) Read() {
	msg.ReadAt = time.Now()
}
