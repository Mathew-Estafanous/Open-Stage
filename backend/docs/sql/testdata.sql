-- Creating Accounts (Used Bcrypt hashing for passwords)
INSERT INTO accounts (name, username, password, email)
    VALUES ('Mathew', 'MatMat',
            '$2y$10$PlK/Si1y4zwR/4lp9WqELubyLbHiIWm5Xj9Q6Jf9.GNkJg.gmBaVK', 'mathew@gmail.com');
INSERT INTO accounts (name, username, password, email)
    VALUES ('Elijah', 'JarJarBinx',
            '$2y$10$5V80wTKIXcx31evy0bWnyOUADoeHdmGMooBMHUn7PFwQEcJEg5/n2', 'jaja@gmail.com');

-- Creating Rooms
INSERT INTO rooms (host, room_code, fk_account_id) VALUES ('Mathew', 'cppcGroup', 1);
INSERT INTO rooms (host, room_code, fk_account_id) VALUES ('Elijah', 'goto', 2);

-- Adding Questions to 'cppcGroup' room.
INSERT INTO questions (question, fk_room_code) VALUES ('How is everything?', 'cppcGroup');
INSERT INTO questions (question, fk_room_code) VALUES ('Are we allowed to ask this?', 'cppcGroup');
INSERT INTO questions (question, fk_room_code, total_likes) VALUES ('Is anime dumb?', 'cppcGroup', 2);

-- Adding Questions to 'goto' room.
INSERT INTO questions (question, fk_room_code) VALUES ('How big is the earth?', 'goto');
INSERT INTO questions (question, fk_room_code) VALUES ('how much water should I drink?', 'goto');
INSERT INTO questions (question, fk_room_code, total_likes) VALUES ('Is the new iPhone good?', 'goto', 1);
