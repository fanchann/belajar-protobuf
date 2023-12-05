### Id Number
Best practice:
- gunakan id number 1-15 untuk field yang tidak bisa kosong (not null)
- gunakan id number 16-19 untuk field yang bisa kosong (nullable)
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

### Json TagName

```proto3
syntax = "proto3";

package example;

message Response{
	string user_firstname = 1[json_name="user_firstname"];
	repeated string user_hobbies = 2[json_name="user_hobbies"];
}
```
jika tidak diberi tag, maka akan menghasilkan output :
- userFirstname
- userHobbies
### Any
any adalah tipe data yang menerima apapun, `interface{}` adalah representasi tipe data di golang
```proto3

package example;

import "google/protobuf/any.proto"

message Response{
	google.protobuf.Any data = 1;
}

```
### OneOf
Oneof adalah tipe data yang sudah ditentukan tipe datanya di awal.
example:
```proto3
syntax = "proto3";

package example;
option "go_package=protogen/";

message Users{
	string user_id = 1;
	string username = 2;
	oneof allowed_socmed {
		SocialMediaChannel socmed = 16;
		ChattingChannel chat = 17;
	} // one of tidak akan menerima tipe data selain SocialMediaChannel dan ChattingChannel
};

message SocialMediaChannel{
	string instagram_profile = 1;
	string facebook_profile = 2;
}

mesage ChattingChannel{
	string discord_profile = 1;
	string whatsapp_number = 2;
}
```

### Schema evolution
digunakan untuk forward dan backward,
- Forward Compatibility
	- Forward compatibility mengacu pada kemampuan sistem yang telah diupdate atau dimutakhirkan untuk dapat berinteraksi dengan versi yang lebih baru dari skema data.
	- Dalam kasus ini, setiap kali dilakukan pembaruan pada skema data (menambahkan field), sistem masih mampu membaca dan memproses data dari versi sebelumnya tanpa mengalami error.
- Backward Compatibility
	- Backward compatibility mengacu pada kemampuan sistem yang telah diupdate atau dimutakhirkan untuk dapat berinteraksi dengan versi yang lebih lama dari skema data.
	- Dalam kasus ini, setiap kali dilakukan pembaruan pada skema data (menghapus field), sistem masih mampu membaca dan memproses data dari versi yang lebih baru tanpa mengalami error.

Dengan kata lain:
- Forward compatibility menguji apakah sistem yang lebih baru dapat memproses data yang dikirim oleh sistem yang lebih lama.
- Backward compatibility menguji apakah sistem yang lebih lama dapat memproses data yang dikirim oleh sistem yang lebih baru.

example:
user_register_form_v1.proto
```proto3
message UserRegisterForm{
	string user_email = 1;
	string user_password = 2;
}
```
user_register_form_v2.proto
```proto3
message UserRegisterForm{
	string user_email = 1;
	string user_password = 2;
	string user_number = 3;
	string user_birth_date = 4;
	string user_gender = 5;
}
```
user_register_form_v3.proto
```proto3
message UserRegisterForm{
	reserved 4,5;
	string user_email = 1;
	string user_password = 2;
	string user_number = 3;
}
```

`reserved` : adalah fungsi untuk menghapus id number di field, jika id number di field sudah dihapus maka id tersebut sudah tidak bisa digunakan lagi

