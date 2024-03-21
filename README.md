Deploy เลยก็ได้นะ

# Events
url = "/event/events"
### Request
Just Get
### Response
```json
  {
    [
      {
        "Id": 1,
        "Title": "Music Concert",
        "Lat": 13.8517,
        "Lon": 100.5678,
        "StartDate": "2024-02-10T18:00:00Z",
        "EndDate": "2024-02-12T22:00:00Z",
        "created_at": "2024-03-19T06:23:20.652Z",
        "updated_at": "2024-03-19T06:23:20.652Z",
        "price": 10.99,
        "rating": 4.5,
        "image": "https://news.airbnb.com/wp-content/uploads/sites/4/2019/06/PJM020719Q202_Luxe_WanakaNZ_LivingRoom_0264-LightOn_R1.jpg?fit=2500%2C1666",
        "creator": "InnovateTech",
        "detail": "Explore the latest in technology",
        "location_name": "Concert Hall",
        "need_regis": true
      },
      ...
    ]
  }
```
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

