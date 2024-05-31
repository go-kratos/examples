package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/menu"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) *MenuRepo {
	l := log.NewHelper(log.With(logger, "module", "menu/repo/admin-service"))
	return &MenuRepo{
		data: data,
		log:  l,
	}
}

func (r *MenuRepo) convertEntToProto(in *ent.Menu) *v1.Menu {
	if in == nil {
		return nil
	}

	var menuType *v1.MenuType
	if in.Type != nil {
		menuType = (*v1.MenuType)(trans.Int32(v1.MenuType_value[string(*in.Type)]))
	}

	return &v1.Menu{
		Id:                in.ID,
		ParentId:          in.ParentID,
		OrderNo:           in.OrderNo,
		Name:              in.Name,
		Title:             in.Title,
		Path:              in.Path,
		Component:         in.Component,
		Icon:              in.Icon,
		KeepAlive:         in.KeepAlive,
		Show:              in.Show,
		IsExt:             in.IsExt,
		ExtUrl:            in.ExtURL,
		Permissions:       in.Permissions,
		HideTab:           in.HideTab,
		HideMenu:          in.HideMenu,
		HideBreadcrumb:    in.HideBreadcrumb,
		CurrentActiveMenu: in.CurrentActiveMenu,
		Redirect:          in.Redirect,
		Type:              menuType,
		Status:            (*string)(in.Status),
		CreateTime:        util.TimeToTimeString(in.CreateTime),
		UpdateTime:        util.TimeToTimeString(in.UpdateTime),
		DeleteTime:        util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *MenuRepo) travelChild(nodes []*v1.Menu, node *v1.Menu) bool {
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

func (r *MenuRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Menu.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *MenuRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListMenuResponse, error) {
	builder := r.data.db.Client().Menu.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), menu.FieldCreateTime,
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

	items := make([]*v1.Menu, 0, len(results))
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

	return &v1.ListMenuResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *MenuRepo) Get(ctx context.Context, req *v1.GetMenuRequest) (*v1.Menu, error) {
	ret, err := r.data.db.Client().Menu.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *MenuRepo) Create(ctx context.Context, req *v1.CreateMenuRequest) error {
	builder := r.data.db.Client().Menu.Create().
		SetNillableName(req.Menu.Name).
		SetNillableStatus((*menu.Status)(req.Menu.Status)).
		SetNillableParentID(req.Menu.ParentId).
		SetNillablePath(req.Menu.Path).
		SetNillableOrderNo(req.Menu.OrderNo).
		SetNillableComponent(req.Menu.Component).
		SetNillableIcon(req.Menu.Icon).
		SetNillableKeepAlive(req.Menu.KeepAlive).
		SetNillableShow(req.Menu.Show).
		SetNillableIsExt(req.Menu.IsExt).
		SetNillableExtURL(req.Menu.ExtUrl).
		SetPermissions(req.Menu.Permissions).
		SetCreateTime(time.Now())

	if req.Menu.Type != nil {
		builder.SetType((menu.Type)(req.Menu.Type.String()))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *MenuRepo) Update(ctx context.Context, req *v1.UpdateMenuRequest) error {
	builder := r.data.db.Client().Menu.UpdateOneID(req.Menu.Id).
		SetNillableName(req.Menu.Name).
		SetNillableStatus((*menu.Status)(req.Menu.Status)).
		SetNillableParentID(req.Menu.ParentId).
		SetNillablePath(req.Menu.Path).
		SetNillableOrderNo(req.Menu.OrderNo).
		SetNillableComponent(req.Menu.Component).
		SetNillableIcon(req.Menu.Icon).
		SetNillableKeepAlive(req.Menu.KeepAlive).
		SetNillableShow(req.Menu.Show).
		SetNillableIsExt(req.Menu.IsExt).
		SetNillableExtURL(req.Menu.ExtUrl).
		SetUpdateTime(time.Now())

	if req.Menu.Permissions != nil {
		builder.SetPermissions(req.Menu.Permissions)
	}
	if req.Menu.Type != nil {
		builder.SetType((menu.Type)(req.Menu.Type.String()))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *MenuRepo) Delete(ctx context.Context, req *v1.DeleteMenuRequest) (bool, error) {
	err := r.data.db.Client().Menu.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
