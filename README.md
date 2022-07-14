# simpson

## data structure reponse

- Có data trả về:
```
{
    "status":"OK", // OK or FAILED
    "users": [
        {
            "name": "abc",
            "age": 123
        },
        {}
    ]
}
```

- không có data trả về:
    - ok
```
{
    "status":"OK", // OK or FAILED
}
```
    - false
```
{
    "status":"OK", // OK or FAILED
    "code":"-1" // "-1","-2",...
    "message":"database error"
}
```
