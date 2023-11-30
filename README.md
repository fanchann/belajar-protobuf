## Scallar

| ProtoBuf | Golang | Notes |
|-----------|-----------|-----------|
|double|double|
|float|float|
|int32|int32| 32 bit, signed, variant length bytes|
|int64|int64| 64 bit, signed, variant length bytes|
|uint32|uint32|32 bit, unsigned|
|uint64|uint64| 64 bit, unsigned|
|sint32|int32|32 bit, signed|
|sint64|int64| 64 bit,signed|
|bool|bool| true or false|
|string|string|string|
|bytes|[]byte|

## Repeated Field, Enum, Comments
 - repeated ( like array / list)
```proto3
syntax = "proto3";

message User{
	repeated string emails = 5;
}
```

- enum  ( nilai yang sudah ditentukan ) & Comments
```proto3
/*
This is comment
*/
// this also comment
syntax = "proto3";

enum Gender{
	Gender_UNDEFINED = 0;
	Gender_MALE = 1;
	Gender_FEMALE = 2;
}

message User{
	string name = 1;
	Gender gender = 5;
}
```


