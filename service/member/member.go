package member

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	common "github.com/paper-trade-chatbot/be-common"
	"github.com/paper-trade-chatbot/be-member/dao/memberDao"
	"github.com/paper-trade-chatbot/be-member/database"
	"github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-proto/member"
	"golang.org/x/crypto/bcrypt"
)

type MemberIntf interface {
	CreateMember(ctx context.Context, in *member.CreateMemberReq) (*member.CreateMemberRes, error)
	GetMember(ctx context.Context, in *member.GetMemberReq) (*member.GetMemberRes, error)
	GetMembers(ctx context.Context, in *member.GetMembersReq) (*member.GetMembersRes, error)
	ModifyMember(ctx context.Context, in *member.ModifyMemberReq) (*member.ModifyMemberRes, error)
	ResetPassword(ctx context.Context, in *member.ResetPasswordReq) (*member.ResetPasswordRes, error)
	DeleteMember(ctx context.Context, in *member.DeleteMemberReq) (*member.DeleteMemberRes, error)
}

type MemberImpl struct {
	MemberClient member.MemberServiceClient
}

func New() MemberIntf {
	return &MemberImpl{}
}

func (impl *MemberImpl) CreateMember(ctx context.Context, in *member.CreateMemberReq) (*member.CreateMemberRes, error) {
	db := database.GetDB()

	checkAccountForm := struct {
		Account string `valid:"account~account:6-12" json:"account"`
	}{
		Account: in.Account,
	}

	if _, err := govalidator.ValidateStruct(checkAccountForm); err != nil {
		return nil, common.ErrWrongAccountFormat
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	id, err := memberDao.New(db, databaseModels.MemberModel{
		Account:      in.Account,
		PasswordHash: string(passwordHash),
		Mail:         in.Mail,
		LineID:       in.LineID,
		RoleCode:     int32(in.RoleCode),
		Status:       int32(in.Status),
		VerifyStatus: int32(in.VerifyStatus),
		GroupID:      uint64(in.GroupID),
		CreatedAt:    time.Unix(in.CreatedAtUnix, 0),
	})
	if err != nil {
		return nil, err
	}

	return &member.CreateMemberRes{
		Id: int64(id),
	}, nil
}

func (impl *MemberImpl) GetMember(ctx context.Context, in *member.GetMemberReq) (*member.GetMemberRes, error) {
	db := database.GetDB()

	queryModel := &memberDao.QueryModel{}

	switch query := in.Member.(type) {
	case *member.GetMemberReq_Id:
		queryModel.ID = uint64(query.Id)
	case *member.GetMemberReq_Account:
		queryModel.Account = query.Account
	case *member.GetMemberReq_Mail:
		queryModel.Mail = query.Mail
	case *member.GetMemberReq_LineID:
		queryModel.LineID = query.LineID
	default:
		return nil, common.ErrNoQueryCondition
	}

	model, err := memberDao.Get(db, queryModel)
	if err != nil {
		return nil, err
	}

	if model == nil {
		return &member.GetMemberRes{}, nil
	}

	return &member.GetMemberRes{
		Member: &member.Member{
			Account:       model.Account,
			Mail:          model.Mail,
			LineID:        model.LineID,
			CountryCode:   model.CountryCode,
			Phone:         model.Phone,
			RoleCode:      member.RoleCodeType(model.RoleCode),
			Status:        member.StatusType(model.Status),
			VerifyStatus:  member.VerifyStatus(model.VerifyStatus),
			GroupID:       int64(model.GroupID),
			CreatedAtUnix: model.CreatedAt.Unix(),
		},
	}, nil

}

func (impl *MemberImpl) GetMembers(ctx context.Context, in *member.GetMembersReq) (*member.GetMembersRes, error) {
	return nil, common.ErrNotImplemented
}

func (impl *MemberImpl) ModifyMember(ctx context.Context, in *member.ModifyMemberReq) (*member.ModifyMemberRes, error) {
	return nil, common.ErrNotImplemented
}

func (impl *MemberImpl) ResetPassword(ctx context.Context, in *member.ResetPasswordReq) (*member.ResetPasswordRes, error) {
	return nil, common.ErrNotImplemented
}

func (impl *MemberImpl) DeleteMember(ctx context.Context, in *member.DeleteMemberReq) (*member.DeleteMemberRes, error) {
	return nil, common.ErrNotImplemented
}
