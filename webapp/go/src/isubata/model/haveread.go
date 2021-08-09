package model

import "time"

type HaveRead struct {
	UserID    int64     `db:"user_id"`
	ChannelID int64     `db:"channel_id"`
	MessageID int64     `db:"message_id"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type HaveReads []*HaveRead

func (ms HaveReads) LastMessageID(userID, channelID int64) int64 {
	var lastMessageID int64
	for _, m := range ms {
		if m.UserID == userID && m.ChannelID == channelID {
			lastMessageID = m.MessageID
			break
		}
	}
	return lastMessageID
}
