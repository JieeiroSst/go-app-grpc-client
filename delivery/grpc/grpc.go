package grpc

import (
	"context"
	"time"

	"github.com/JIeeiroSst/go-app/domain"
	"github.com/JIeeiroSst/go-app/log"
	"github.com/JIeeiroSst/go-app/proto"
)

type Service struct {
	repo domain.Repository
	grpcClient proto.UserProfileClient
	contextTimeout time.Duration
}

func NewService(repo domain.Repository,grpcClient proto.UserProfileClient,contextTimeout time.Duration) *Service{
	return &Service{
		repo:  repo,
		grpcClient: grpcClient,
		contextTimeout: contextTimeout,
	}
}

func (s *Service) UserAll() (users []domain.User, err error){
	users,err = s.repo.UserAll()
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return nil,err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) UserById(id int) (user domain.User,err error){
	user,err = s.repo.UserById(id)
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return domain.User{},err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) CreateUser(user domain.User) (err error) {
	err = s.repo.CreateUser(user)
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) UpdateUser(id int,user domain.User) (err error) {
	err= s.repo.UpdateUser(id,user)
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) DeleteUser(id int) (err error){
	err = s.repo.DeleteUser(id)
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) ProfileAll() (profiles []domain.Profile,err error) {
	profiles,err = s.repo.ProfileAll()
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return nil,err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) ProfileById(id int) (profile domain.Profile,err error){
	profile,err =s.repo.ProfileById(id)
	if err!=nil {
		log.InitZapLog().Error("server running error")
		return domain.Profile{},err
	}
	log.InitZapLog().Info("server running")
	return 
}

func (s *Service) UpdateEmail(ctx context.Context,id int,profile domain.Profile) (bool,string) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	in:=proto.UpdateEmailRequest{
		Id: int32(id),
		Email: profile.Email,
	}
	res,err:=s.grpcClient.UpdateEmail(ctx,&in)
	if err != nil {
		return false, "update failed"
	}
	if res.Ok {
		if err != nil {
			return false, "update failed"
		}
		return true, "update success"
	}
	return false, "update failed"
}

func (s *Service) CreateEmail(ctx context.Context,profile domain.Profile) (bool,string){
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	in:=proto.CreateEmailRequest{
		Name: profile.Name,
		Email: profile.Email,
		UserId: int32(profile.UserId),
	}
	res,err:=s.grpcClient.CreateEmail(ctx,&in)
	if err != nil {
		return false, "create failed"
	}
	if res.Ok {
		if err != nil {
			return false, "create failed"
		}
		return true,"create success"
	}
	return false,"create failed"
}

func (s *Service) DeleteEmail(ctx context.Context,id int) (bool,string){
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	in:=proto.DeleteEmailRequest{
		Id: int32(id),
	}
	res,err:=s.grpcClient.DeleteEmail(ctx,&in)
	if err != nil {
		return false, "delete failed"
	}
	if res.Ok {
		if err != nil {
			return false,"delete failed"
		}
		return true, "delete success"
	}
	return false, "delete failed"
}