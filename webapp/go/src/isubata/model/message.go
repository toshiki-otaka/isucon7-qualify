package model

import "time"

type Message struct {
	ID        int64     `db:"id"`
	ChannelID int64     `db:"channel_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type Messages []*Message

func (ms Messages) CountByIDAndChannelID(id, channelID int64) int64 {
	var count int64
	for _, m := range ms {
		if m.ChannelID == channelID && id < m.ID {
			count++
		}
	}
	return count
}

func (ms Messages) CountByChannelID(channelID int64) int64 {
	var count int64
	for _, m := range ms {
		if m.ChannelID == channelID {
			count++
		}
	}
	return count
}
