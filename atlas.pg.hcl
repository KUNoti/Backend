table "events" {
    schema = schema.public
    column "id" {
    null = false
    type = serial
  }
  column "title" {
    null = false
    type = character_varying(255)
  }
  column "latitude" {
    type = float
    null = false
  }
  column "longitude" {
    type = float
    null = false
  }
  column "start_date" {
    null = false
    type = timestamp(3)
  }
  column "end_date" {
    null = false
    type = timestamp(3)
  }
  column "price" {
    null = false
    type = float
  }
  column "tag" {
    null = true
    type = character_varying(255)
  }
  column "image" {
    null = true
    type = character_varying(255)
  }
  column "creator" {
    null = false
    type = int
  }
  column "detail" {
    null = false
    type = character_varying(255)
  }
  column "location_name" {
    null = false
    type = character_varying(255)
  }
  column "need_regis" {
    null = false
    type = boolean
  }
  column "created_at" {
    null    = false
    type    = timestamp(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamp(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
}

table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "social_id" {
    null = true
    type = varchar(255)
  }
  column "email" {
    null = false
    type = varchar(255)
  }
  column "profile_image" {
    null = true
    type = varchar(255)
  }
  column "username" {
    null = false
    type = varchar(255)
  }
  column "password" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamp(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
}

schema "public" {
  comment = "standard public schema"
}