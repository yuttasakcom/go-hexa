BEGIN;

SET TIME ZONE 'Asia/Bangkok';

CREATE TABLE "todos" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR NOT NULL,
  "completed" BOOLEAN DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todos" ("title");

COMMIT;