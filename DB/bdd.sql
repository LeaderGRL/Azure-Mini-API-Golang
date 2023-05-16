CREATE TABLE IF NOT EXISTS users (
  id INTEGER,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE(`username`),
  UNIQUE(`email`)
);

-- INSERT INTO users (username, password, email, created_at, updated_at) 
-- VALUES ('admin', 'admin', 'admin@gmail.com', '2016-01-03', '2016-01-03');