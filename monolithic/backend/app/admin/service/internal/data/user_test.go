package data

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"

	userV1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

var reSpaces = regexp.MustCompile(`\s+`)

func TestUserFieldMask(t *testing.T) {
	u := &userV1.User{
		UserName: trans.String("UserName"),
		RealName: trans.String("RealName"),
		//Avatar:   trans.String("Avatar"),
		Address: trans.String("Address"),
	}

	updateUserReq := &userV1.UpdateUserRequest{
		User: &userV1.User{
			UserName: trans.String("UserName1"),
			RealName: trans.String("RealName1"),
			//Avatar:   trans.String("Avatar1"),
			Address: trans.String("Address1"),
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"userName", "realName", "avatar", "roleId"},
		},
	}
	updateUserReq.UpdateMask.Normalize()
	if !updateUserReq.UpdateMask.IsValid(u) {
		// Return an error.
		panic("invalid field mask")
	}

	fieldmaskutil.Filter(updateUserReq.GetUser(), updateUserReq.UpdateMask.GetPaths())
	proto.Merge(u, updateUserReq.GetUser())

	fmt.Println(reSpaces.ReplaceAllString(u.String(), " "))
}

func TestFilterReuseMask(t *testing.T) {
	users := []*userV1.User{
		{
			Id:       1,
			UserName: trans.String("name 1"),
		},
		{
			Id:       2,
			UserName: trans.String("name 2"),
		},
	}
	// Create a mask only once and reuse it.
	mask := fieldmaskutil.NestedMaskFromPaths([]string{"userName", "realName", "positionId"})
	for _, user := range users {
		mask.Filter(user)
	}
	fmt.Println(users)
	assert.Equal(t, len(users), 2)
	// Output: [userName:"name 1" userName:"name 2"]
}

func TestNilValuePaths(t *testing.T) {
	u := &userV1.User{
		Id:       2,
		UserName: trans.String("name 2"),
		//RealName: trans.String(""),
	}
	paths := []string{"userName", "realName", "positionId"}
	nilPaths := fieldmaskutil.NilValuePaths(u, paths)
	fmt.Println(nilPaths)
	fmt.Println(u.PositionId)
}

func TestMessageNil(t *testing.T) {
	u := &userV1.User{
		Id:       2,
		UserName: trans.String("name 2"),
	}

	pr := u.ProtoReflect()
	md := pr.Descriptor()
	fd := md.Fields().ByName("userName")
	if fd == nil {

	} else {
		fmt.Println(fd, fd.Name())
	}

	v := pr.Get(fd)
	fmt.Println(v)
}
