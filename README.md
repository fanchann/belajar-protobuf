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

## Import
- directory app
```sh
├── proto
│   ├── hello.proto
│   └── world.proto
```

- hello.proto
```proto3
syntax = "proto3";

package hello;
option go_package="gen/"; // file generated will save in proto/gen/

message Hello{
	string name = 1;
};
```

- world.proto
```proto3
syntax = "proto3";

package hello;
option go_package="gen/"; // file generated will save in proto/gen/

import "hello.proto"

message World{
	int32 id = 1;
	Hello hello = 2;
};
```

## Nested message
```proto3
syntax = "proto3";

package basic

message User{
	int32 id_user = 1[json_name="id_user"]; // default is IdUser,
	string name = 2;	
}

message Address{
	string city = 1;
	string country = 2;
	Coordinate coordinate = 16; // use 1 - 15 for importants fields,
	message Coordinate{
		double latitude = 1; // double == float64
		double longtitude = 2;
	}
}
```