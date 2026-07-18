-- ────────────────────────────────────────────
-- 1. USERS
-- ────────────────────────────────────────────
CREATE TABLE users (
    uid            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR(100) NOT NULL,
    email          VARCHAR(255) NOT NULL UNIQUE,
    password_hash  VARCHAR(255) NOT NULL,
    avatar_url     VARCHAR(255),
    is_online      BOOLEAN NOT NULL DEFAULT FALSE,
    last_seen_at   TIMESTAMPTZ,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ────────────────────────────────────────────
-- 2. CONVERSATION
-- ────────────────────────────────────────────
CREATE TABLE conversation (
    conversation_id  INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name             VARCHAR(100),  -- NULL nếu là private (tự lấy tên người kia)
    type             VARCHAR(20) NOT NULL DEFAULT 'private'
                     CHECK (type IN ('private', 'group')),
    avatar_url       VARCHAR(255),
    created_by       UUID REFERENCES users(uid) ON DELETE SET NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ────────────────────────────────────────────
-- 3. CONVERSATION_MEMBER
-- ────────────────────────────────────────────
CREATE TABLE conversation_member (
    conversation_id  INT NOT NULL REFERENCES conversation(conversation_id)
                     ON DELETE CASCADE ON UPDATE CASCADE,
    uid              UUID NOT NULL REFERENCES users(uid)
                     ON DELETE CASCADE ON UPDATE CASCADE,
    role             VARCHAR(20) NOT NULL DEFAULT 'member'
                     CHECK (role IN ('admin', 'member')),
    joined_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (conversation_id, uid)
);

-- ────────────────────────────────────────────
-- 4. MESSAGE
-- ────────────────────────────────────────────
CREATE TABLE message (
    message_id       BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    conversation_id  INT NOT NULL REFERENCES conversation(conversation_id)
                     ON DELETE CASCADE ON UPDATE CASCADE,
    sender           UUID REFERENCES users(uid)
                     ON DELETE SET NULL ON UPDATE CASCADE,
    content          VARCHAR(1000) NOT NULL,
    message_type     VARCHAR(20) NOT NULL DEFAULT 'text'
                     CHECK (message_type IN ('text', 'image', 'file', 'system')),
    is_edited        BOOLEAN NOT NULL DEFAULT FALSE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ────────────────────────────────────────────
-- 5. NOTIFICATION
-- ────────────────────────────────────────────
CREATE TABLE notification (
    notification_id  BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    type             VARCHAR(50) NOT NULL,  -- 'new_message','group_invite','system'...
    content          VARCHAR(255) NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ────────────────────────────────────────────
-- 6. USER_NOTIFICATION
-- ────────────────────────────────────────────
CREATE TABLE user_notification (
    uid              UUID NOT NULL REFERENCES users(uid)
                     ON DELETE CASCADE ON UPDATE CASCADE,
    notification_id  BIGINT NOT NULL REFERENCES notification(notification_id)
                     ON DELETE CASCADE ON UPDATE CASCADE,
    is_read          BOOLEAN NOT NULL DEFAULT FALSE,
    read_at          TIMESTAMPTZ,
    PRIMARY KEY (uid, notification_id)
);

-- ============================================================
-- INDEXES — PostgreSQL KHÔNG tự tạo index cho foreign key
-- ============================================================
CREATE INDEX idx_conversation_member_uid ON conversation_member(uid);
CREATE INDEX idx_message_conversation_id ON message(conversation_id);
CREATE INDEX idx_message_sender ON message(sender);
CREATE INDEX idx_message_created_at ON message(created_at DESC); -- cho query "tin nhắn mới nhất"
CREATE INDEX idx_user_notification_uid ON user_notification(uid);
CREATE INDEX idx_user_notification_unread ON user_notification(uid) WHERE is_read = FALSE; -- partial index, tối ưu đếm "chưa đọc"