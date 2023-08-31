CREATE TABLE users (
  user_id VARCHAR PRIMARY KEY,
  created_at TIMESTAMP not null default now()
);

CREATE TABLE segments (
  slug VARCHAR PRIMARY KEY,
  created_at TIMESTAMP not null default now()
);

CREATE TABLE user_segments (
  user_id VARCHAR REFERENCES users(user_id),
  segment_slug VARCHAR REFERENCES segments(slug) ON DELETE CASCADE,
  PRIMARY KEY (user_id, segment_slug)
);

CREATE TABLE user_ttl (
  task_id SERIAL PRIMARY KEY,
  user_id VARCHAR REFERENCES users(user_id),
  segment_slug VARCHAR,
  ttl TIMESTAMP
);