package tests

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/fanchann/belajar-protobuf/protogen"
	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHelloProto(t *testing.T) {
	var hello protogen.Hello

	hello.Name = "Tifanni"

	fmt.Printf("Original  : %v\n", hello)
	fmt.Printf("As String : %v\n", hello.String())
}

func selectInformationChannel(types string) *anypb.Any {
	var anyPb anypb.Any

	paperMail := protogen.PaperMail{
		PaperMail: "This is papermail",
	}
	socialMedia := protogen.SocialMedia{
		SocialMedia: "This is social media",
	}
	instantMessage := protogen.InstantMessaging{
		InstantMessaging: "This is instant messaging",
	}

	switch types {
	case "papermail":
		anypb.MarshalFrom(&anyPb, &paperMail, proto.MarshalOptions{})
	case "socmed", "socialmedia":
		anypb.MarshalFrom(&anyPb, &socialMedia, proto.MarshalOptions{})
	case "instantmsg":
		anypb.MarshalFrom(&anyPb, &instantMessage, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&anyPb, &instantMessage, proto.MarshalOptions{})

	}
	return &anyPb
}

func TestPrintProtoUser(t *testing.T) {
	var user protogen.User

	user.Id = 1
	user.Username = "kaori miyazono"
	user.Emails = []string{
		"kaorimiyazono@gmail.com",
	}
	user.Gender = protogen.Gender_Gender_FEMALE
	user.Password = []byte("kaorimiyazono")
	user.IsActive = true
	user.Addresses = &protogen.Address{
		City:    "Semarang",
		Country: "Indonesia",
		Street:  "Jl Chinchin",
		Coordinate: &protogen.Address_Coordinate{
			Latitude:  5.2,
			Longitude: 5.44,
		},
	}
	user.SocialMedia = selectInformationChannel("")
	user.SkillUser = map[string]*protogen.Skill{
		"basic_skill": &protogen.Skill{
			NameSkill:   "Nangis",
			RatingSkill: 3,
		},
	}
	user.BirthDate = &protogen.Date{
		Year:  2003,
		Month: 12,
		Day:   26,
	}
	user.LastKnowLocation = &protogen.LatLng{
		Latitude:  113.31212,
		Longitude: 12.131213,
	}
	user.LastLoginUser = &timestamppb.Timestamp{
		Seconds: 30,
		Nanos:   12,
	}

	fmt.Printf("Original  : %v\n", &user)
	fmt.Printf("As String : %v\n", user.String())

	// json
	userJson, err := protojson.Marshal(&user)
	assert.Nil(t, err)

	fmt.Printf("As Json : %v\n", string(userJson))
	fmt.Printf("As Byte : %v\n", userJson)

}

// jsonpb as 3party
func TestMarshalProtoToJsonJsonPb(t *testing.T) {
	userDataExample := protogen.User{
		Id:       2,
		Username: "Tiffani Aulia",
		Password: []byte("tifanni"),
		Emails:   []string{"tiffani@gmail.com"},
		IsActive: true,
	}

	var buff bytes.Buffer

	err := (&jsonpb.Marshaler{}).Marshal(&buff, &userDataExample)
	assert.Nil(t, err)

	jsonStr := buff.String()
	fmt.Println(jsonStr)
}

func TestUnmarshalJsonToProtoJsonPb(t *testing.T) {
	jsonStr := `{"id":2,"username":"Tiffani Aulia","is_active":true,"password":"dGlmYW5uaQ==","emails":["tiffani@gmail.com"]}`

	buff := strings.NewReader(jsonStr)
	protoObj := new(protogen.User)

	err := (&jsonpb.Unmarshaler{}).Unmarshal(buff, protoObj)
	assert.Nil(t, err)

	fmt.Println(protoObj)
}

// protojson also 3party
func TestMarshalProtoToJsonProtoJson(t *testing.T) {
	userDataExample := protogen.User{
		Id:       2,
		Username: "Tiffani Aulia",
		Password: []byte("tifanni"),
		Emails:   []string{"tiffani@gmail.com"},
		IsActive: true,
	}

	jsonResult, err := protojson.Marshal(&userDataExample)
	assert.Nil(t, err)

	fmt.Println(string(jsonResult))
}

func TestUnmarshalJsonToProtoProtoJson(t *testing.T) {
	jsonStr := `{"id":2,"username":"Tiffani Aulia","is_active":true,"password":"dGlmYW5uaQ==","emails":["tiffani@gmail.com"]}`

	protoObj := new(protogen.User)

	err := protojson.Unmarshal([]byte(jsonStr), protoObj)
	assert.Nil(t, err)

	fmt.Println(protoObj)
}

func TestPrintJsonGroupUsers(t *testing.T) {
	groupUsers := protogen.UserGroup{
		GroupId:     1,
		GroupName:   "Tifanni Fan's",
		Roles:       []string{"admin", "moderator", "women in tech", "member"},
		Description: "This group for tifanni fan",
	}

	jsonResult, err := protojson.Marshal(&groupUsers)
	assert.Nil(t, err)
	fmt.Println(string(jsonResult))
}

func TestPrintApplication(t *testing.T) {
	appService := protogen.ApplicationServices{
		AppService: &protogen.AppsService{
			AppId:   "com.x",
			AppName: "X",
		},
		Description: "This is app service",
	}
	jobService := protogen.JobServices{
		JobService: &protogen.JobService{
			JobId:   "xxx",
			JobName: "Romusha",
		},
		Description: "This is job service",
	}

	fmt.Println(appService)
	fmt.Println(jobService)

	// to json
	appServiceJson, err := protojson.Marshal(&appService)
	assert.Nil(t, err)

	fmt.Println(string(appServiceJson))
}

func TestWriteUserProtoToFile(t *testing.T) {
	var user protogen.User

	user.Id = 1
	user.Username = "kaori miyazono"
	user.Emails = []string{
		"kaorimiyazono@gmail.com",
	}
	user.Gender = protogen.Gender_Gender_FEMALE
	user.Password = []byte("kaorimiyazono")
	user.IsActive = true
	user.Addresses = &protogen.Address{
		City:    "Semarang",
		Country: "Indonesia",
		Street:  "Jl Chinchin",
		Coordinate: &protogen.Address_Coordinate{
			Latitude:  5.2,
			Longitude: 5.44,
		},
	}
	user.SocialMedia = selectInformationChannel("")
	user.SkillUser = map[string]*protogen.Skill{
		"basic_skill": &protogen.Skill{
			NameSkill:   "Nangis",
			RatingSkill: 3,
		},
	}

	protoBytes, err := proto.Marshal(&user)
	assert.Nil(t, err)
	err = ioutil.WriteFile("../kaori.bin", protoBytes, fs.ModePerm)
	assert.Nil(t, err)

}

func TestReadBinFileGeneratedFromProtobuf(t *testing.T) {
	var user protogen.User

	bytes, err := ioutil.ReadFile("../kaori.bin")
	assert.Nil(t, err)

	err = proto.Unmarshal(bytes, &user)
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestWriteProtoToJsonFile(t *testing.T) {
	var user protogen.User

	user.Id = 1
	user.Username = "kaori miyazono"
	user.Emails = []string{
		"kaorimiyazono@gmail.com",
	}
	user.Gender = protogen.Gender_Gender_FEMALE
	user.Password = []byte("kaorimiyazono")
	user.IsActive = true
	user.Addresses = &protogen.Address{
		City:    "Semarang",
		Country: "Indonesia",
		Street:  "Jl Chinchin",
		Coordinate: &protogen.Address_Coordinate{
			Latitude:  5.2,
			Longitude: 5.44,
		},
	}
	user.SocialMedia = selectInformationChannel("")
	user.SkillUser = map[string]*protogen.Skill{
		"basic_skill": &protogen.Skill{
			NameSkill:   "Nangis",
			RatingSkill: 3,
		},
	}

	byteJson, err := protojson.Marshal(&user)
	assert.Nil(t, err)

	err = ioutil.WriteFile("../user.json", byteJson, fs.ModePerm)
	assert.Nil(t, err)
}

func TestReadJsonFileProto(t *testing.T) {
	var user protogen.User

	fileByte, err := ioutil.ReadFile("../user.json")
	assert.Nil(t, err)

	err = protojson.Unmarshal(fileByte, &user)
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestWriteFileBinUserContentV1(t *testing.T) {
	userContent := protogen.UserContent{
		UserContentId: 1,
		Slug:          "Slug",
		// Title:         "This is title",
		HtmlContent: "example",
		// AuthorId:      10,
	}

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	ioutil.WriteFile("../userContentv1.bin", byteJson, fs.ModePerm)
}

func TestReadFileBinUserContentV1(t *testing.T) {
	var userContent protogen.UserContent

	bytesData, err := ioutil.ReadFile("../userContentv1.bin")
	assert.Nil(t, err)

	err2 := protojson.Unmarshal(bytesData, &userContent)
	assert.Nil(t, err2)

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	fmt.Println(string(byteJson))
}

func TestWriteFileBinUserContentV2(t *testing.T) {
	userContent := protogen.UserContent{
		UserContentId: 1,
		Slug:          "Slug",
		// Title:         "This is title",
		HtmlContent: "example",
		// AuthorId:      10,
		Category: "example",
	}

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	ioutil.WriteFile("../userContentv2.bin", byteJson, fs.ModePerm)
}

func TestReadFileBinUserContentV2(t *testing.T) {
	var userContent protogen.UserContent

	bytesData, err := ioutil.ReadFile("../userContentv2.bin")
	assert.Nil(t, err)

	protojson.Unmarshal(bytesData, &userContent)
	// assert.Nil(t, err2)

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	fmt.Println(string(byteJson))
}

func TestWriteFileBinUserContentV3(t *testing.T) {
	userContent := protogen.UserContent{
		UserContentId: 1,
		Slug:          "Slug",
		// Title:         "This is title",
		HtmlContent: "example",
		// AuthorId:      10,
		Category: "example",
		Bio:      "This is bio",
	}

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	ioutil.WriteFile("../userContentv3.bin", byteJson, fs.ModePerm)
}

func TestReadFileBinUserContentV3(t *testing.T) {
	var userContent protogen.UserContent

	bytesData, err := ioutil.ReadFile("../userContentv3.bin")
	assert.Nil(t, err)

	err2 := protojson.Unmarshal(bytesData, &userContent)
	assert.Nil(t, err2)

	byteJson, err := protojson.Marshal(&userContent)
	assert.Nil(t, err)

	fmt.Println(string(byteJson))
}
