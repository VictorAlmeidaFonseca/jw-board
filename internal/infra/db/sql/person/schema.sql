CREATE TABLE
  IF NOT EXISTS person (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    role_id INTEGER,
    FOREIGN KEY (role_id) REFERENCES roles (id)
  );