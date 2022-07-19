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

```
{
    "status":"OK", // OK or FAILED
}
```

```
{
    "status":"OK", // OK or FAILED
    "code":"-1" // "-1","-2",...
    "message":"database error"
}
```

## gen key

```
ssh-keygen -t rsa -b 4096 -E SHA512 -m PEM -P "" -f RS512.key
openssl rsa -in RS512.key -pubout -outform PEM -out RS512.key.pub
```