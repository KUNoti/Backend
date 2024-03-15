Deploy เลยก็ได้นะ

# Events
url = ""
### Request
Just Get
### Response

# (Create) Event 
url = ""
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
url = ""
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
url = ""
### Request
```json
{
  "username" : "username",
  "password" : "password"
}
```
### Response

# Create User (Post)
url = ""
### Request
```json
  "username" : "username",
  "password" : "password",
  "email" : "email"
  "image" : "TODO!!!!!!!!!!!!!!!!!!!!!!!!!! can be null"
```

# สมัทรเข้า event 
### Request
```json
{
  "userId" : "userId",
  "reason" : "reason"
}
```
# My Event
### Creator      
Todo show list ของคนที่สมัทรเข้าevent ข้อม ให้ กด approve ไรงี้

### คนเข้าร่วม หรือ following      
Todo

