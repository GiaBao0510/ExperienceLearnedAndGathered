package impl

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
	"github.com/GiaBao0510/Go-Realtime/internal/repositories"
)

type messageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) repositories.IMessageRepository {
	return &messageRepo{db: db}
}

func(r *messageRepo) Create(ctx context.Context, msg *model.DBMessage) error {
	querry := `
		INSERT INTO (conversation_id, sender, content, message_type)
		Values ($1, $2, $3, $4)
		Returning message_id, create_at;
	`

	return r.db.QueryRowContext(ctx, query, 
		msg.ConversationID,
		msg.SenderUID,
		msg.Content,
		msg.MessageType,	
	).Scan(&msg.Message_id, &msg.Create_at)
}

func(r *messageRepo) GetByConversation(ctx context.Context, conversationID int, limit, offset int) ([]model.DBMessage, error) {
	
	querry := `
		SELECT message_id, conversation_id, sender, content, message_type, is_edited, created_at
		FROM message
		WHERE conversation_id = $1
		ORDER BY create_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryRowCotext(ctx, querry, conversationID. limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query messages: %w", err)
	}

	defer rows.Close()

	var messages []model.DBMessage
	for rows.Next() {
		var mssg model.DBMessage
		if err := rows.Scan(
			&msg.MessageID,
			&msg.ConversationID,
			&msg.SenderUID,
			&msg.Content,
			&msg.MessageType,
			&msg.IsEdited,
			&msg.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}

		messages = append(messages, msg)
	}
	return messages, rows.Err()
}

func(r *messageRepo) GetLatestByConversation(ctx context.Context, conversationID int) (*model.DBMessage, error){
	querry := `
		SELECT message_id, conversation_id, sender, content, message_type, is_edited, created_at
		FROM message
		WHERE conversation_id = $1
		ORDER BY create_at DESC
		LIMIT 1
	`
	var msg model.DBMessage
	err := r.db.QueryRowCotext(ctx, querry, conversationID).Scan(
		&msg.MessageID,
		&msg.ConversationID,
		&msg.SenderUID,
		&msg.Content,
		&msg.MessageType,
		&msg.IsEdited,
		&msg.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("query latest message: %w", err)

	}

	return &msg, nil
}