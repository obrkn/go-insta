USE twitter;

INSERT INTO users (nickname, password)
VALUES ('user1@example.com', 'password'),
       ('user2@example.com', 'password'),
       ('user3@example.com', 'password'),
       ('user4@example.com', 'password'),
       ('user5@example.com', 'password'),
       ('user6@example.com', 'password'),
       ('user7@example.com', 'password'),
       ('user8@example.com', 'password'),
       ('user9@example.com', 'password'),
       ('user10@example.com', 'password'),
       ('user11@example.com', 'password'),
       ('user12@example.com', 'password');

INSERT INTO tweets (user_id, message)
VALUES (1, 'hello, world!'),
       (1, 'hello, world!'),
       (1, 'hello, world!'),
       (2, 'hello, world!'),
       (2, 'hello, world!'),
       (2, 'hello, world!'),
       (2, 'hello, world!'),
       (3, 'hello, world!'),
       (3, 'hello, world!'),
       (4, 'hello, world!'),
       (4, 'hello, world!'),
       (4, 'hello, world!');