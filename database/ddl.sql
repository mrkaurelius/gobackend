DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    username text,
    password text
);

-- password sha256 hex digest 
-- 1234 sha256 => 03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4

CREATE TABLE posts (
    user text,
    date text,
    title text,
    post text
);

-- datetime format YYYY-MM-DD

INSERT INTO posts (user, date, title, post) 
    VALUES("mrk","2020-11-11","title","lorem ipsum dolor sit amet");
INSERT INTO posts (user, date, title, post)
    VALUES("mrk","2020-11-12","title","lorem ipsum dolor sit amet");
INSERT INTO posts (user, date, title, post)
    VALUES("mrk","2020-11-13","title","lorem ipsum dolor sit amet");


INSERT INTO users (username, password) 
    VALUES("mrk","03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4");
INSERT INTO users (username, password) 
    VALUES("ayse","03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4");

select * from posts;