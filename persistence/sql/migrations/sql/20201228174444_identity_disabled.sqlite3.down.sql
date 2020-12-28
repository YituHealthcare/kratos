CREATE TABLE "_identities_tmp" (
"id" TEXT PRIMARY KEY,
"schema_id" TEXT NOT NULL,
"traits" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
INSERT INTO "_identities_tmp" (id, schema_id, traits, created_at, updated_at) SELECT id, schema_id, traits, created_at, updated_at FROM "identities";

DROP TABLE "identities";
ALTER TABLE "_identities_tmp" RENAME TO "identities";