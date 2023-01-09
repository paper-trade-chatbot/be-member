package member

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	common "github.com/paper-trade-chatbot/be-common"
	"github.com/paper-trade-chatbot/be-common/database"
	"github.com/paper-trade-chatbot/be-common/logging"
	"github.com/paper-trade-chatbot/be-member/dao/memberDao"
	"github.com/paper-trade-chatbot/be-member/dao/memberGroupDao"
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
	CreateMemberGroup(ctx context.Context, in *member.CreateMemberGroupReq) (*member.CreateMemberGroupRes, error)
	GetMemberGroups(ctx context.Context, in *member.GetMemberGroupsReq) (*member.GetMemberGroupsRes, error)
	DeleteMemberGroup(ctx context.Context, in *member.DeleteMemberGroupReq) (*member.DeleteMemberGroupRes, error)
}

type MemberImpl struct {
	MemberClient member.MemberServiceClient
}

func New() MemberIntf {
	return &MemberImpl{}
}

func (impl *MemberImpl) CreateMember(ctx context.Context, in *member.CreateMemberReq) (*member.CreateMemberRes, error) {
	db := database.GetDB()

	logging.Info(ctx, "CreateMember: %s %s %s", in.Account, in.Password, in.Mail, in.GroupID)

	checkAccountForm := struct {
		Account string `valid:"stringlength(6|12)" json:"account"`
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

	if in.GroupID == 0 {
		in.GroupID = 1
	}

	_, err = memberDao.New(db, databaseModels.MemberModel{
		Account:      in.Account,
		PasswordHash: string(passwordHash),
		Mail:         in.Mail,
		LineID:       in.LineID,
		Status:       int32(in.Status),
		VerifyStatus: int32(in.VerifyStatus),
		GroupID:      uint64(in.GroupID),
		CreatedAt:    time.Unix(in.CreatedAtUnix, 0),
	})
	if err != nil {
		return nil, err
	}

	return &member.CreateMemberRes{}, nil
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
			Id:            model.ID,
			Account:       model.Account,
			Mail:          model.Mail,
			LineID:        model.LineID,
			CountryCode:   model.CountryCode,
			Phone:         model.Phone,
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

func (impl *MemberImpl) CreateMemberGroup(ctx context.Context, in *member.CreateMemberGroupReq) (*member.CreateMemberGroupRes, error) {
	return nil, common.ErrNotImplemented
}

func (impl *MemberImpl) GetMemberGroups(ctx context.Context, in *member.GetMemberGroupsReq) (*member.GetMemberGroupsRes, error) {
	db := database.GetDB()

	queryModel := &memberGroupDao.QueryModel{}

	if in.Id != nil {
		queryModel.ID = *in.Id
	}
	if in.Name != nil {
		queryModel.Name = *in.Name
	}

	models, err := memberGroupDao.Gets(db, queryModel)
	if err != nil {
		return nil, err
	}

	res := &member.GetMemberGroupsRes{}

	for _, m := range models {
		memberGroup := &member.MemberGroup{
			Id:        m.ID,
			Name:      m.Name,
			Memo:      m.Memo,
			CreatedAt: m.CreatedAt.Unix(),
		}
		res.MemberGroups = append(res.MemberGroups, memberGroup)
	}

	return res, nil
}

func (impl *MemberImpl) DeleteMemberGroup(ctx context.Context, in *member.DeleteMemberGroupReq) (*member.DeleteMemberGroupRes, error) {
	return nil, common.ErrNotImplemented
}
