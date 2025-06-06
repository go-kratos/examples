package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dict"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

type DictRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictRepo(data *Data, logger log.Logger) *DictRepo {
	l := log.NewHelper(log.With(logger, "module", "dict/repo/admin-service"))
	return &DictRepo{
		data: data,
		log:  l,
	}
}

func (r *DictRepo) convertEntToProto(in *ent.Dict) *v1.Dict {
	if in == nil {
		return nil
	}
	return &v1.Dict{
		Id:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		CreatorId:   in.CreateBy,
		CreateTime:  util.TimeToTimeString(in.CreateTime),
		UpdateTime:  util.TimeToTimeString(in.UpdateTime),
		DeleteTime:  util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *DictRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Dict.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *DictRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListDictResponse, error) {
	builder := r.data.db.Client().Dict.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dict.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
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

	items := make([]*v1.Dict, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListDictResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *DictRepo) Get(ctx context.Context, req *v1.GetDictRequest) (*v1.Dict, error) {
	ret, err := r.data.db.Client().Dict.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *DictRepo) Create(ctx context.Context, req *v1.CreateDictRequest) error {
	err := r.data.db.Client().Dict.Create().
		SetNillableName(req.Dict.Name).
		SetNillableDescription(req.Dict.Description).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *DictRepo) Update(ctx context.Context, req *v1.UpdateDictRequest) error {
	builder := r.data.db.Client().Dict.UpdateOneID(req.Dict.Id).
		SetNillableName(req.Dict.Name).
		SetNillableDescription(req.Dict.Description).
		SetUpdateTime(time.Now())

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *DictRepo) Delete(ctx context.Context, req *v1.DeleteDictRequest) (bool, error) {
	err := r.data.db.Client().Dict.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
