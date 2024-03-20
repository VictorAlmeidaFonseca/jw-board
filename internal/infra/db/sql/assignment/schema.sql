CREATE TABLE
  IF NOT EXISTS assignments (
    id INTEGER PRIMARY KEY,
    date_assignment TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id INTEGER,
    person_id INTEGER,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (person_id) REFERENCES person (id)
  );