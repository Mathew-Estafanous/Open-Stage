CREATE TABLE "rooms" (
     "room_code" varchar(15) PRIMARY KEY,
     "host" varchar(100) NOT NULL
);

CREATE TABLE "questions" (
     "question_id" SERIAL PRIMARY KEY,
     "question" varchar(2000) NOT NULL,
     "questioner_name" varchar(45) NOT NULL DEFAULT 'Anonymous',
     "total_likes" int NOT NULL DEFAULT 0,
     "fk_room_code" varchar(15) NOT NULL
);

ALTER TABLE "questions" ADD FOREIGN KEY ("fk_room_code") REFERENCES "rooms" ("room_code") ON DELETE CASCADE;