-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "tests" table
CREATE TABLE "public"."tests" ("id" serial NOT NULL, "title" character varying(255) NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
