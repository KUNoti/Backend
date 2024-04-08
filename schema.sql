-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "events" table
CREATE TABLE "public"."events" ("id" serial NOT NULL, "start_date" timestamp(3) NOT NULL, "end_date" timestamp(3) NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "title" character varying(255) NOT NULL, "latitude" double precision NOT NULL, "longitude" double precision NOT NULL, "price" double precision NOT NULL, "image" character varying(255) NULL, "detail" character varying(255) NOT NULL, "location_name" character varying(255) NOT NULL, "need_regis" boolean NOT NULL, "tag" character varying(255) NULL, "creator" integer NOT NULL, "regis_amount" integer NULL, "regis_max" integer NULL, PRIMARY KEY ("id"));
-- Create "follow_by_tag" table
CREATE TABLE "public"."follow_by_tag" ("id" serial NOT NULL, "tag" character varying(255) NOT NULL, "user_token" character varying(255) NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "following_events" table
CREATE TABLE "public"."following_events" ("id" serial NOT NULL, "event_id" integer NOT NULL, "user_id" integer NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "regis_events" table
CREATE TABLE "public"."regis_events" ("id" serial NOT NULL, "event_id" integer NOT NULL, "user_id" integer NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "public"."users" ("id" serial NOT NULL, "name" character varying(255) NOT NULL, "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "email" character varying(255) NOT NULL, "profile_image" character varying(255) NULL, "username" character varying(255) NOT NULL, "password" character varying(255) NOT NULL, "social_id" character varying(255) NULL, "token" character varying(255) NOT NULL, PRIMARY KEY ("id"));
