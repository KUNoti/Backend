Deploy เลยก็ได้นะ

# Events
### Request
Just Get
### Response

# (Post) Event 
### Request
```json
{
  "latitude" : 12.312,
  "longitude" : 1200.123,
  "title" : "title",
  "image" : TODO!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!,
  "creator" : "prem" -> User ต้อง getจากDB,
  "detail" : "detail",
  "tag" : "String",
  "locationName" : "locationName",
  "startDateTime" : DateTime,
  "endDateTime" : DateTime,
  ""
}
```

# Search 
### Request
```json
{
  "keyword" : "string"
}
```

```json
{
  "tag" : "string"
}
```

### Response
[Event]

# Auth (login)
### Request
```json
{
  "username" : "username",
  "password" : "password"
}
```
### Response

# Create User (Post)
### Request
```json
  "username" : "username",
  "password" : "password",
  "email" : "email"
  "image" : "TODO!!!!!!!!!!!!!!!!!!!!!!!!!! can be null"
```

# สมัทรเข้า event 
```json
{
  "userId" : "userId",
  "reason" : "reason"
}
```
# My Event
Creator
Todo show list ของคนที่สมัทรเข้าevent ข้อม ให้ กด approve ไรงี้

คนเข้าร่วม หรือ following
Todo

