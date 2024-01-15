package service

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-monolithic-demo/app/admin/service/internal/data"

	adminV1 "kratos-monolithic-demo/gen/api/go/admin/service/v1"
	systemV1 "kratos-monolithic-demo/gen/api/go/system/service/v1"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
)

type RouterService struct {
	adminV1.RouterServiceHTTPServer

	log *log.Helper
	uc  *data.MenuRepo
}

func NewRouterService(logger log.Logger, uc *data.MenuRepo) *RouterService {
	l := log.NewHelper(log.With(logger, "module", "router/service/admin-service"))
	return &RouterService{
		log: l,
		uc:  uc,
	}
}

func (s *RouterService) ListPermissionCode(_ context.Context, _ *emptypb.Empty) (*adminV1.ListPermissionCodeResponse, error) {
	return &adminV1.ListPermissionCodeResponse{}, nil
}

func fillRouteItem(menus []*systemV1.Menu) []*adminV1.RouteItem {
	if len(menus) == 0 {
		return nil
	}

	var routers []*adminV1.RouteItem

	for _, v := range menus {
		if !v.GetShow() {
			continue
		}
		if v.GetType() == systemV1.MenuType_BUTTON {
			continue
		}

		item := &adminV1.RouteItem{
			Path:      v.Path,
			Component: v.Component,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Meta: &adminV1.RouteMeta{
				OrderNo:           v.OrderNo,
				Title:             v.Title,
				Icon:              v.Icon,
				HideBreadcrumb:    v.HideBreadcrumb,
				HideTab:           v.HideTab,
				HideMenu:          v.HideMenu,
				CurrentActiveMenu: v.CurrentActiveMenu,
			},
		}

		if len(v.Children) > 0 {
			item.Children = fillRouteItem(v.Children)
		}

		routers = append(routers, item)
	}

	return routers
}

func (s *RouterService) ListRoute(ctx context.Context, _ *emptypb.Empty) (*adminV1.ListRouteResponse, error) {
	menuList, err := s.uc.List(ctx, &pagination.PagingRequest{
		NoPaging: trans.Bool(true),
	})
	if err != nil {
		s.log.Errorf("查询列表发生错误[%s]", err.Error())
		return nil, errors.New("读取列表发生错误")
	}

	resp := &adminV1.ListRouteResponse{Items: fillRouteItem(menuList.Items)}

	return resp, nil
}
