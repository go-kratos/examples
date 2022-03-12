package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-realtimemap/api/admin/v1"
	"kratos-realtimemap/app/admin/internal/pkg/data"
)

func (s *AdminService) GetOrganizations(_ context.Context, _ *emptypb.Empty) (*v1.GetOrganizationsReply, error) {
	reply := &v1.GetOrganizationsReply{
		Organizations: data.AllOrganizations.MapToBaseInfoArray(),
	}

	return reply, nil
}

func (s *AdminService) GetGeofences(_ context.Context, req *v1.GetGeofencesReq) (*v1.GetGetGeofencesReply, error) {
	if org, ok := data.AllOrganizations[req.OrgId]; ok {
		return &v1.GetGetGeofencesReply{
			Id:        org.Id,
			Name:      org.Name,
			Geofences: org.MapToGeofenceArray(),
		}, nil
	} else {
		return nil, v1.ErrorResourceNotFound(fmt.Sprintf("Organization %s not found", req.OrgId))
	}
}

func (s *AdminService) GetPositionsHistory(_ context.Context, req *v1.GetPositionsHistoryReq) (*v1.GetPositionsHistoryReply, error) {
	his := s.positionHistory.GetPositionsHistory(req.Id)
	if his == nil {
		return nil, v1.ErrorResourceNotFound(fmt.Sprintf("%s positions history not found", req.Id))
	}
	return &v1.GetPositionsHistoryReply{Positions: his}, nil
}
