CREATE TABLE "updates" (
  "version" varchar UNIQUE PRIMARY KEY,
  "path" varchar NOT NULL,
  "description" varchar NOT NULL,
  "checksum" varchar NOT NULL,
  "date" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "devices" (
  "device_id" varchar UNIQUE NOT NULL PRIMARY KEY,
  "device_version" varchar NOT NULL DEFAULT '1.0.0',
  "last_update" timestamp NOT NULL DEFAULT 'now()',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "devices" ADD FOREIGN KEY ("device_version") REFERENCES "updates" ("version");

