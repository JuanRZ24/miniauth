ALTER TABLE users
ADD COLUMN email_verified_at TIMESTAMP NULL;

CREATE TABLE email_verification_tokens (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  token_hash TEXT NOT NULL UNIQUE,
  expires_at TIMESTAMP NOT NULL,
  used_at TIMESTAMP NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),

  CONSTRAINT fk_email_verification_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_email_verification_token_hash
ON email_verification_tokens(token_hash);

CREATE UNIQUE INDEX uniq_active_email_verification_token
ON email_verification_tokens(user_id)
WHERE used_at IS NULL;
