CREATE TABLE IF NOT EXISTS "rooms" (
     "room_code" varchar(15) PRIMARY KEY,
     "host" varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "questions" (
     "question_id" SERIAL PRIMARY KEY,
     "question" varchar(2000) NOT NULL,
     "questioner_name" varchar(45) NOT NULL DEFAULT 'Anonymous',
     "total_likes" int NOT NULL DEFAULT 0,
     "fk_room_code" varchar(15) NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS(SELECT 1 FROM pg_constraint WHERE conname = 'fk_room_code') THEN
        ALTER TABLE "questions" ADD FOREIGN KEY ("fk_room_code")
            REFERENCES "rooms" ("room_code") ON DELETE CASCADE;
    END IF;
END;
$$;