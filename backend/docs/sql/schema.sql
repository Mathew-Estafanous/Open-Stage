CREATE TABLE IF NOT EXISTS "rooms" (
     "room_code" varchar(15) PRIMARY KEY,
     "host" varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "questions" (
     "question_id" SERIAL PRIMARY KEY,
     "question" VARCHAR(2000) NOT NULL,
     "questioner_name" VARCHAR(45) NOT NULL DEFAULT 'Anonymous',
     "total_likes" INT NOT NULL DEFAULT 0,
     "fk_room_code" VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS "accounts" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS(SELECT 1 FROM pg_constraint WHERE conname = 'fk_room_code') THEN
        ALTER TABLE "questions" ADD FOREIGN KEY ("fk_room_code")
            REFERENCES "rooms" ("room_code") ON DELETE CASCADE;
    END IF;
END;
$$;