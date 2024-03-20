  CREATE TABLE
    IF NOT EXISTS person_role (
      id INTEGER PRIMARY KEY,
      person_id INTEGER,
      role_id INTEGER,
      FOREIGN KEY (role_id) REFERENCES roles (id)
      FOREIGN KEY (person_id) REFERENCES person (id)
    );