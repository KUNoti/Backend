-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "events" table
CREATE TABLE "public"."events" ("id" serial NOT NULL, "start_date" timestamp(3) NOT NULL, "end_date" timestamp(3) NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "title" character varying(255) NOT NULL, "latitude" double precision NOT NULL, "longitude" double precision NOT NULL, "price" double precision NOT NULL, "rating" double precision NOT NULL, "image" character varying(255) NOT NULL, "creator" character varying(255) NOT NULL, "detail" character varying(255) NOT NULL, "location_name" character varying(255) NOT NULL, "need_regis" boolean NOT NULL, PRIMARY KEY ("id"));
-- Create "tags" table
CREATE TABLE "public"."tags" ("id" serial NOT NULL, "color" character varying(20) NOT NULL, "name" character varying(255) NOT NULL, "category" character varying(255) NOT NULL, "is_official" boolean NOT NULL DEFAULT false, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "public"."users" ("id" serial NOT NULL, "name" character varying(255) NOT NULL, "role" character varying(255) NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
