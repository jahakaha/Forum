CREATE TABLE IF NOT EXISTS user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username text NOT NULL unique CHECK(LENGTH(username) <= 40),
  email text NOT NULL unique,
  password_hash text NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(40) NOT NULL CHECK(LENGTH(title) <= 40),
  body TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  amount_likes INTEGER NOT NULL,
  amount_dislikes INTEGER NOT NULL,
  posted_on TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE IF NOT EXISTS liked (
  id integer PRIMARY KEY AUTOINCREMENT,
  liked boolean NOT NULL CHECK(liked IN(0, 1)),
  disliked boolean NOT NULL CHECK(disliked IN(0, 1)),
  post_id integer NOT NULL,
  auth_id integer NOT NULL,
  foreign key(post_id) REFERENCES posts(id),
  foreign key(auth_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS session (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid text NOT NULL,
  auth_id INTEGER,
  FOREIGN KEY (auth_id) REFERENCES user (id)
);

CREATE TABLE IF NOT EXISTS comments (
  id integer PRIMARY KEY,
  comment text,
  auth_id integer NOT NULL,
  post_id integer NOT NULL,
  amount_likes INTEGER NOT NULL DEFAULT 0,
  amount_dislikes INTEGER NOT NULL DEFAULT 0,
  commented_on TEXT NOT NULL,
  FOREIGN KEY (auth_id) REFERENCES user (id) FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS comlikes (
  id integer primary key autoincrement,
  liked boolean NOT NULL CHECK (liked IN (0, 1)),
  disliked boolean NOT NULL CHECK (liked IN (0, 1)),
  auth_id integer not null,
  com_id integer not null,
  foreign key (auth_id) references user (id) foreign key (com_id) references comments (id)
);

CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(30) NOT NULL unique
);

CREATE TABLE IF NOT EXISTS cat_posts (
  id integer primary key autoincrement,
  post_id integer not null,
  cat_id integer not null,
  foreign key (post_id) references posts (id) foreign key (cat_id) references categories (id)
);

CREATE TABLE IF NOT EXISTS image_post (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  pimage TEXT,
  post_id INTEGER NOT NULL UNIQUE,
  FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS notific_liked (
  id integer PRIMARY KEY AUTOINCREMENT,
  valid boolean NOT NULL CHECK(valid IN(0, 1)),
  notific_date text NOT NULL,
  post_id integer NOT NULL,
  from_auth integer NOT NULL,
  to_auth integer NOT NULL,
  foreign key(post_id) REFERENCES posts(id),
  foreign key(from_auth) REFERENCES user(id),
  foreign key(to_auth) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS notific_comment (
  id integer PRIMARY KEY AUTOINCREMENT,
  valid boolean NOT NULL CHECK(valid IN(0, 1)),
  notific_date text NOT NULL,
  post_id integer NOT NULL,
  com_id integer NOT NULL,
  from_auth integer NOT NULL,
  to_auth integer NOT NULL,
  foreign key(post_id) REFERENCES posts(id),
  foreign key(com_id) REFERENCES comments(id),
  foreign key(from_auth) REFERENCES user(id),
  foreign key(to_auth) REFERENCES user(id)
);