insert into users (name, nick, mail, password)
values
("User 1", "user 1", "user1@gmail.com", "$2a$10$wlMeCePmF6.EhtTbAobEse..6rovcrJ11aovTT2vbX4I4ScjSo6oS"),
("User 2", "user 2", "user2@gmail.com", "$2a$10$wlMeCePmF6.EhtTbAobEse..6rovcrJ11aovTT2vbX4I4ScjSo6oS"),
("User 3", "user 3", "user3@gmail.com", "$2a$10$wlMeCePmF6.EhtTbAobEse..6rovcrJ11aovTT2vbX4I4ScjSo6oS");

insert into followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);