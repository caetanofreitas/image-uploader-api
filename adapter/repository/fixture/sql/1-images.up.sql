CREATE TABLE images
(
  id TEXT NOT NULL,
  name TEXT NOT NULL,
  size REAL NOT NULL,
  extension TEXT NOT NULL,
  status TEXT NOT NULL,
  error_message TEXT,
  created_at    TEXT NOT NULL,
  updated_at    TEXT NOT NULL
);
