truncate table client,appointment,appointment_type

insert into client values (1,'{
    "name": "Dukler",
    "active": true,
    "telephones": [{
        "number": 1556960610,
        "type": "movil",
        "default": true
    }, {
        "number": 46390945,
        "type": "casa",
        "default": false
    }],
    "addresses": [{
        "street": "marcos sastre",
        "number": 5790,
        "location": "devoto",
        "default": true
    }],
    "email": "8amartin@gmail.com",
    "description": "god"
}')

insert into appointment_type values (1,'{
    "name": "Masaje",
    "active": true
}')

insert into appointment values (1,1,1,'{
  "name": "masaje",
  "active": true,
  "date": "2018-03-13T00:17:41+00:00"
}')