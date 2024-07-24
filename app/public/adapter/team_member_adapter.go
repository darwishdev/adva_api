package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) teamMembersListRowGrpcFromSql(resp *db.TeamMembersListRow) *rmsv1.TeamMembersListRow {

	return &rmsv1.TeamMembersListRow{
		TeamMemberId:    resp.TeamMemberID,
		TeamMemberName:  resp.TeamMemberName,
		JobTitle:        resp.JobTitle,
		TeamMemberImage: resp.TeamMemberImage,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:       resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *PublicAdapter) TeamMembersListGrpcFromSql(resp *[]db.TeamMembersListRow) *rmsv1.TeamMembersListResponse {
	records := make([]*rmsv1.TeamMembersListRow, 0)
	deletedRecords := make([]*rmsv1.TeamMembersListRow, 0)
	for _, v := range *resp {
		record := a.teamMembersListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.TeamMembersListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PublicAdapter) TeamMemberCreateSqlFromGrpc(req *rmsv1.TeamMemberCreateRequest) *db.TeamMemberCreateParams {

	return &db.TeamMemberCreateParams{
		TeamMemberName:  req.TeamMemberName,
		JobTitle:        req.JobTitle,
		TeamMemberImage: req.TeamMemberImage,
	}
}
func (a *PublicAdapter) TeamMemberEntityGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMember {

	return &rmsv1.TeamMember{
		TeamMemberId:    int32(resp.TeamMemberID),
		TeamMemberName:  resp.TeamMemberName,
		TeamMemberImage: resp.TeamMemberImage,
		JobTitle:        resp.JobTitle,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt:       resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *PublicAdapter) TeamMemberCreateGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMemberCreateResponse {
	return &rmsv1.TeamMemberCreateResponse{
		TeamMember: a.TeamMemberEntityGrpcFromSql(resp),
	}
}

func (a *PublicAdapter) TeamMemberUpdateSqlFromGrpc(req *rmsv1.TeamMemberUpdateRequest) *db.TeamMemberUpdateParams {

	return &db.TeamMemberUpdateParams{
		TeamMemberID:    req.TeamMemberId,
		JobTitle:        req.JobTitle,
		TeamMemberName:  req.TeamMemberName,
		TeamMemberImage: req.TeamMemberImage,
	}
}
func (a *PublicAdapter) TeamMemberUpdateGrpcFromSql(resp *db.TeamMember) *rmsv1.TeamMemberUpdateResponse {
	return &rmsv1.TeamMemberUpdateResponse{
		TeamMember: a.TeamMemberEntityGrpcFromSql(resp),
	}
}
func (a *PublicAdapter) TeamMemberFindForUpdateGrpcFromSql(resp *db.TeamMemberFindForUpdateRow) *rmsv1.TeamMemberUpdateRequest {
	return &rmsv1.TeamMemberUpdateRequest{
		JobTitle:        resp.JobTitle,
		TeamMemberId:    resp.TeamMemberID,
		TeamMemberName:  resp.TeamMemberName,
		TeamMemberImage: resp.TeamMemberImage,
	}

}

//  TeamMemberFind(ctx context.Context, req *int32) (*db.ProductsSchemaMenuView, error)
