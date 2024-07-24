package usecase

import (
	"context"

	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (s *PublicUsecase) ClientInitialize(ctx context.Context, req *rmsv1.ClientInitializeRequest) (*rmsv1.ClientInitializeResponse, error) {
	projects, err := s.ProjectsList(ctx, &rmsv1.ProjectsListRequest{})
	if err != nil {
		return nil, err
	}

	services, err := s.ServicesList(ctx, &rmsv1.ServicesListRequest{})
	if err != nil {
		return nil, err
	}

	testemonials, err := s.TestemonialsList(ctx, &rmsv1.TestemonialsListRequest{})
	if err != nil {
		return nil, err
	}

	teamMembers, err := s.TeamMembersList(ctx, &rmsv1.TeamMembersListRequest{})
	if err != nil {
		return nil, err
	}
	programs, err := s.ProgramsList(ctx, &rmsv1.ProgramsListRequest{})
	if err != nil {
		return nil, err
	}

	settings, err := s.SettingsFindForUpdate(ctx, &rmsv1.SettingsFindForUpdateRequest{})
	if err != nil {
		return nil, err
	}

	resp := &rmsv1.ClientInitializeResponse{
		Projects:     projects.Records,
		Services:     services.Records,
		Testemonials: testemonials.Records,
		Programs:     programs.Records,
		TeamMembers:  teamMembers.Records,
		Settings:     settings.Settings,
	}

	return resp, nil
}
