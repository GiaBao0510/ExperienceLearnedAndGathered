/*────────────────────────────────────────────*/
/* 1. USERS */
/*────────────────────────────────────────────*/
-- name: CreateUser: execresult
INSERT INTO users (name, email, password_hash, avatar_url)
        VALUES ($1, $2, $3, $4)
        RETURNING uid, created_at;

-- name: GetUserByUID :one
SELECT uid, name, email, password_hash, avatar_url,
               is_online, last_seen_at, created_at
        FROM users WHERE uid = $1;

-- name: Update_Put :exec
UPDATE users
SET name = $2, email = $3, password_hash = $4
WHERE uid = $1;

-- name: Update_Patch :exec
UPDATE users
SET name = COALESCE($2, name), email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash)
WHERE uid = $1;

-- name: UpdateOnlineStatus :exec
UPDATE users
    SET is_online = $1, last_seen_at = $2
    WHERE uid = $3;

-- name: CountOnlineUsers :one
SELECT COUNT(*) FROM users WHERE is_online = TRUE;

/*────────────────────────────────────────────*/
/* 2. CONVERSATION */
/*────────────────────────────────────────────*/
-- name: CreateConversation :execresult
Insert into conversation (name, type, avatar_url, created_by)
        VALUES ($1, $2, $3, $4)
        RETURNING conversation_id, created_at;

-- name: GetConversationByUserID :one
SELECT c.name, c.type, c.avatar_url, c.created_by, c.created_at
FROM conversation c
JOIN conversation_member cm ON c.conversation_id = cm.conversation_id
WHERE cm.uid = $1;

-- name: GetConversationByID :one
SELECT c.name, c.type, c.avatar_url, c.created_by, c.created_at
FROM conversation c
WHERE c.conversation_id = $1;

/*────────────────────────────────────────────*/
/* 3. CONVERSATION_MEMBER */
/*────────────────────────────────────────────*/

-- name: AddMember :exec
INSERT INTO conversation_member (conversation_id, uid, role)
        VALUES ($1, $2, $3);

-- name: RemoveMember :exec
DELETE FROM conversation_member
WHERE conversation_id = $1 AND uid = $2;

-- name: GetMembers: many
SELECT 


/*────────────────────────────────────────────*/
/* 4. MESSAGE */
/*────────────────────────────────────────────*/
-- name: CreateMessage :execresult
INSERT INTO message (conversation_id, sender, content, message_type)
        VALUES ($1, $2, $3, $4)
        RETURNING message_id, created_at;

-- name: GetByConversation :many
SELECT message_id, conversation_id, sender, content,
               message_type, is_edited, created_at
        FROM message
        WHERE conversation_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3;

-- name: GetLatestByConversation :one
SELECT message_id, conversation_id, sender, content,
               message_type, is_edited, created_at
        FROM message
        WHERE conversation_id = $1
        ORDER BY created_at DESC
        LIMIT 1;

/* ────────────────────────────────────────────*/
/* 5. NOTIFICATION */
/*────────────────────────────────────────────*/
-- name: CreateNotification :execresult
INSERT INTO notification (type, content)
        VALUES ($1, $2)
        RETURNING notification_id, created_at;


/*────────────────────────────────────────────*/
/* 6. USER_NOTIFICATION */
/*────────────────────────────────────────────*/
-- name: CreateUserNotification :exec
INSERT INTO user_notification (uid, notification_id)
        VALUES ($1, $2);

-- name: GetUnreadByUser :many
SELECT n.notification_id, n.type, n.content, n.created_at
FROM notification n
INNER JOIN user_notification un ON n.notification_id = un.notification_id
WHERE un.uid = $1 AND un.is_read = FALSE
ORDER BY n.created_at DESC;

-- name: MarkAsRead :exec
UPDATE user_notification
SET is_read = TRUE, read_at = $1
WHERE uid = $2 AND notification_id = $3;

-- name: CountUnreadByUser :one
SELECT COUNT(*) FROM user_notification WHERE uid = $1 AND is_read = FALSE;