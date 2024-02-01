CREATE TABLE follows (
    created_at TEXT,
    updated_at TEXT,
    follower_id INTEGER,
    followee_id INTEGER,
    PRIMARY KEY (follower_id, followee_id)
);

CREATE INDEX idx_follows_follower_id ON follows(follower_id);
CREATE INDEX idx_follows_followee_id ON follows(followee_id);
