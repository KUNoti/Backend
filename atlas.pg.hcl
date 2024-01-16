table "tests" {
    schema = schema.public
    column "id" {
    null = false
    type = serial
  }
  column "title" {
    null = true
    type = character_varying(255)
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