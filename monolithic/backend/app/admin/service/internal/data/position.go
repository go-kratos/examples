package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/position"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type PositionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPositionRepo(data *Data, logger log.Logger) *PositionRepo {
	l := log.NewHelper(log.With(logger, "module", "position/repo/admin-service"))
	return &PositionRepo{
		data: data,
		log:  l,
	}
}

func (r *PositionRepo) convertEntToProto(in *ent.Position) *v1.Position {
	if in == nil {
		return nil
	}
	return &v1.Position{
		Id:         in.ID,
		Name:       &in.Name,
		Code:       &in.Code,
		Remark:     in.Remark,
		OrderNo:    &in.OrderNo,
		ParentId:   &in.ParentID,
		Status:     (*string)(in.Status),
		CreateTime: util.TimeToTimeString(in.CreateTime),
		UpdateTime: util.TimeToTimeString(in.UpdateTime),
		DeleteTime: util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *PositionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Position.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *PositionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPositionResponse, error) {
	builder := r.data.db.Client().Position.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), position.FieldCreateTime,
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

	items := make([]*v1.Position, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListPositionResponse{
		Total: int32(count),
		Items: items,
	}, err
}

func (r *PositionRepo) Get(ctx context.Context, req *v1.GetPositionRequest) (*v1.Position, error) {
	ret, err := r.data.db.Client().Position.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *PositionRepo) Create(ctx context.Context, req *v1.CreatePositionRequest) error {
	err := r.data.db.Client().Position.Create().
		SetNillableName(req.Position.Name).
		SetNillableParentID(req.Position.ParentId).
		SetNillableOrderNo(req.Position.OrderNo).
		SetNillableCode(req.Position.Code).
		SetNillableStatus((*position.Status)(req.Position.Status)).
		SetNillableRemark(req.Position.Remark).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *PositionRepo) Update(ctx context.Context, req *v1.UpdatePositionRequest) error {
	builder := r.data.db.Client().Position.UpdateOneID(req.Position.Id).
		SetNillableName(req.Position.Name).
		SetNillableParentID(req.Position.ParentId).
		SetNillableOrderNo(req.Position.OrderNo).
		SetNillableCode(req.Position.Code).
		SetNillableRemark(req.Position.Remark).
		SetNillableStatus((*position.Status)(req.Position.Status)).
		SetUpdateTime(time.Now())

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *PositionRepo) Delete(ctx context.Context, req *v1.DeletePositionRequest) (bool, error) {
	err := r.data.db.Client().Position.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
