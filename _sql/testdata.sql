INSERT INTO rooms (host, room_code) VALUES ('Mathew', 'cppcGroup');
INSERT INTO rooms (host, room_code) VALUES ('Elijah', 'goto');

INSERT INTO questions (question, fk_room_code) VALUES ('How is everything?', 'cppcGroup');
INSERT INTO questions (question, fk_room_code) VALUES ('Are we allowed to ask this?', 'cppcGroup');
INSERT INTO questions (question, fk_room_code, total_likes) VALUES ('Is anime dumb?', 'cppcGroup', 2);

INSERT INTO questions (question, fk_room_code) VALUES ('How big is the earth?', 'goto');
INSERT INTO questions (question, fk_room_code) VALUES ('how much water should I drink?', 'goto');
INSERT INTO questions (question, fk_room_code, total_likes) VALUES ('Is the new iPhone good?', 'goto', 1);
