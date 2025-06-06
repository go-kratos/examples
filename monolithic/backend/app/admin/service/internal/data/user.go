package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	timeUtil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/user"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/admin-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}

	var authority *v1.UserAuthority
	if in.Authority != nil {
		authority = (*v1.UserAuthority)(trans.Int32(v1.UserAuthority_value[string(*in.Authority)]))
	}

	return &v1.User{
		Id:            in.ID,
		RoleId:        in.RoleID,
		WorkId:        in.WorkID,
		OrgId:         in.OrgID,
		PositionId:    in.PositionID,
		CreatorId:     in.CreateBy,
		UserName:      in.Username,
		NickName:      in.NickName,
		RealName:      in.RealName,
		Email:         in.Email,
		Avatar:        in.Avatar,
		Phone:         in.Phone,
		Gender:        (*string)(in.Gender),
		Address:       in.Address,
		Description:   in.Description,
		Authority:     authority,
		LastLoginTime: in.LastLoginTime,
		LastLoginIp:   in.LastLoginIP,
		Status:        (*string)(in.Status),
		CreateTime:    timeUtil.TimeToTimeString(in.CreateTime),
		UpdateTime:    timeUtil.TimeToTimeString(in.UpdateTime),
		DeleteTime:    timeUtil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	ret, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)

	return u, err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) error {
	builder := r.data.db.Client().User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableNickName(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetNillableRealName(req.User.RealName).
		SetNillablePhone(req.User.Phone).
		SetNillableOrgID(req.User.OrgId).
		SetNillableRoleID(req.User.RoleId).
		SetNillableWorkID(req.User.WorkId).
		SetNillablePositionID(req.User.PositionId).
		SetNillableAvatar(req.User.Avatar).
		SetNillableStatus((*user.Status)(req.User.Status)).
		SetNillableGender((*user.Gender)(req.User.Gender)).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now())

	if len(req.Password) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}
	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) error {
	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.User) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetUser(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().User.UpdateOneID(req.User.Id).
		SetNillableNickName(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetNillableRealName(req.User.RealName).
		SetNillablePhone(req.User.Phone).
		SetNillableOrgID(req.User.OrgId).
		SetNillableRoleID(req.User.RoleId).
		SetNillableWorkID(req.User.WorkId).
		SetNillablePositionID(req.User.PositionId).
		SetNillableAvatar(req.User.Avatar).
		SetNillableStatus((*user.Status)(req.User.Status)).
		SetNillableGender((*user.Gender)(req.User.Gender)).
		SetUpdateTime(time.Now())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}
	if len(req.Password) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.User, req.GetUpdateMask().GetPaths())
		_, nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	ret, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	return u, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.User, error) {
	ret, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, v1.ErrorUserNotExist("用户未找到")
	}

	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *ret.Password)
	if !bMatched {
		return nil, v1.ErrorInvalidPassword("密码错误")
	}

	u := r.convertEntToProto(ret)
	return u, err
}

func (r *UserRepo) UserExists(ctx context.Context, req *v1.UserExistsRequest) (*v1.UserExistsResponse, error) {
	count, _ := r.data.db.Client().User.
		Query().
		Select(user.FieldID).
		Where(user.UsernameEQ(req.GetUserName())).
		Count(ctx)
	return &v1.UserExistsResponse{
		Exist: count > 0,
	}, nil
}
