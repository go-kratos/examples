package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/organization"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type OrganizationRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrganizationRepo(data *Data, logger log.Logger) *OrganizationRepo {
	l := log.NewHelper(log.With(logger, "module", "organization/repo/admin-service"))
	return &OrganizationRepo{
		data: data,
		log:  l,
	}
}

func (r *OrganizationRepo) convertEntToProto(in *ent.Organization) *v1.Organization {
	if in == nil {
		return nil
	}
	return &v1.Organization{
		Id:         in.ID,
		Name:       in.Name,
		Remark:     in.Remark,
		OrderNo:    in.OrderNo,
		ParentId:   in.ParentID,
		Status:     (*string)(in.Status),
		CreateTime: util.TimeToTimeString(in.CreateTime),
		UpdateTime: util.TimeToTimeString(in.UpdateTime),
		DeleteTime: util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *OrganizationRepo) travelChild(nodes []*v1.Organization, node *v1.Organization) bool {
	if nodes == nil {
		return false
	}

	if node.ParentId == nil {
		nodes = append(nodes, node)
		return true
	}

	for _, n := range nodes {
		if node.ParentId == nil {
			continue
		}

		if n.Id == *node.ParentId {
			n.Children = append(n.Children, node)
			return true
		} else {
			if r.travelChild(n.Children, node) {
				return true
			}
		}
	}
	return false
}

func (r *OrganizationRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Organization.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *OrganizationRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListOrganizationResponse, error) {
	builder := r.data.db.Client().Organization.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), organization.FieldCreateTime,
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
		return nil, err
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[j].ParentID == nil {
			return true
		}
		if results[i].ParentID == nil {
			return true
		}
		return *results[i].ParentID < *results[j].ParentID
	})

	items := make([]*v1.Organization, 0, len(results))
	for _, m := range results {
		if m.ParentID == nil {
			item := r.convertEntToProto(m)
			items = append(items, item)
		}
	}
	for _, m := range results {
		if m.ParentID != nil {
			item := r.convertEntToProto(m)
			r.travelChild(items, item)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	ret := v1.ListOrganizationResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *OrganizationRepo) Get(ctx context.Context, req *v1.GetOrganizationRequest) (*v1.Organization, error) {
	ret, err := r.data.db.Client().Organization.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *OrganizationRepo) Create(ctx context.Context, req *v1.CreateOrganizationRequest) error {
	builder := r.data.db.Client().Organization.Create().
		SetNillableName(req.Org.Name).
		SetNillableParentID(req.Org.ParentId).
		SetNillableOrderNo(req.Org.OrderNo).
		SetNillableRemark(req.Org.Remark).
		SetNillableStatus((*organization.Status)(req.Org.Status)).
		SetCreateTime(time.Now()).
		SetCreateBy(req.GetOperatorId())

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *OrganizationRepo) Update(ctx context.Context, req *v1.UpdateOrganizationRequest) error {

	builder := r.data.db.Client().Organization.UpdateOneID(req.Org.Id).
		SetNillableName(req.Org.Name).
		SetNillableParentID(req.Org.ParentId).
		SetNillableOrderNo(req.Org.OrderNo).
		SetNillableRemark(req.Org.Remark).
		SetNillableStatus((*organization.Status)(req.Org.Status)).
		SetUpdateTime(time.Now())

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *OrganizationRepo) Delete(ctx context.Context, req *v1.DeleteOrganizationRequest) (bool, error) {
	err := r.data.db.Client().Organization.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
