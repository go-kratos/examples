package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/role"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	l := log.NewHelper(log.With(logger, "module", "role/repo/admin-service"))
	return &RoleRepo{
		data: data,
		log:  l,
	}
}

func (r *RoleRepo) convertEntToProto(in *ent.Role) *v1.Role {
	if in == nil {
		return nil
	}
	return &v1.Role{
		Id:         in.ID,
		Name:       in.Name,
		Code:       in.Code,
		Remark:     in.Remark,
		OrderNo:    in.OrderNo,
		ParentId:   in.ParentID,
		Status:     (*string)(in.Status),
		CreateTime: util.TimeToTimeString(in.CreateTime),
		UpdateTime: util.TimeToTimeString(in.UpdateTime),
		DeleteTime: util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *RoleRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Role.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *RoleRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListRoleResponse, error) {
	builder := r.data.db.Client().Role.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), role.FieldCreateTime,
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
		return nil, err
	}

	items := make([]*v1.Role, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListRoleResponse{
		Total: int32(count),
		Items: items,
	}, err
}

func (r *RoleRepo) Get(ctx context.Context, req *v1.GetRoleRequest) (*v1.Role, error) {
	ret, err := r.data.db.Client().Role.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *RoleRepo) Create(ctx context.Context, req *v1.CreateRoleRequest) error {
	err := r.data.db.Client().Role.Create().
		SetNillableName(req.Role.Name).
		SetNillableParentID(req.Role.ParentId).
		SetNillableOrderNo(req.Role.OrderNo).
		SetNillableCode(req.Role.Code).
		SetNillableStatus((*role.Status)(req.Role.Status)).
		SetNillableRemark(req.Role.Remark).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *RoleRepo) Update(ctx context.Context, req *v1.UpdateRoleRequest) error {
	builder := r.data.db.Client().Role.UpdateOneID(req.Role.Id).
		SetNillableName(req.Role.Name).
		SetNillableParentID(req.Role.ParentId).
		SetNillableOrderNo(req.Role.OrderNo).
		SetNillableCode(req.Role.Code).
		SetNillableRemark(req.Role.Remark).
		SetNillableStatus((*role.Status)(req.Role.Status)).
		SetUpdateTime(time.Now())

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *RoleRepo) Delete(ctx context.Context, req *v1.DeleteRoleRequest) (bool, error) {
	err := r.data.db.Client().Role.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
