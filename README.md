# Simple GO REST API

<img src="https://1.bp.blogspot.com/-J_7Dq7ovbwQ/XmFfXD2wAmI/AAAAAAAAr6I/6_sUNbgsrsofNg7xoIrJndN1xida3XdDwCLcBGAsYHQ/s1600/Golang.png"
     alt="Go"
     width="80%" height="50%"
     class="aligncenter"
     />

> Simple RESTful API to create, read, update and delete books. No database implementation yet
## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```



## Endpoints

### Get All Persons
``` bash
GET /people
```
### Get Single person
``` bash
GET /people/{id}
```

### Delete Person
``` bash
DELETE /people/{id}
```

### Create Person
``` bash
POST /people/{id}
# Request sample
#    {
#        "id": "1",
#        "firstname": "Juan",
#        "lastname": "Lopez",
#        "address": {
#            "city": "Cali",
#            "state": "Valle del cauca"
#    }
```

### Update Person
``` bash
PUT /people/{id}
# Request sample
#    {
#        "id": "1",
#        "firstname": "Update firstname",
#        "lastname": "Update lastname",
#        "address": {
#            "city": "Cali",
#            "state": "Valle del cauca"
#    }
#    
```
```

### License
This project is licensed under the MIT License
