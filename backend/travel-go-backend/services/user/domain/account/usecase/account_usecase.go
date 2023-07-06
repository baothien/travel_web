package usecase

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/security"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/model"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AccountUseCase struct {
	AccountRepository repository.AccountRepository
}

func NewAccountUseCase(accountRepository repository.AccountRepository) AccountUseCase {
	return AccountUseCase{
		AccountRepository: accountRepository,
	}
}

func (a AccountUseCase) CreateAccount(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = randutil.UUIDRand()
	user.UserName = user.Email
	user.UserType = model.CUSTOMER.Name()
	hash := security.HashAndSalt([]byte(user.Password))
	user.Password = hash

	err := a.AccountRepository.CreateAccount(ctx, user)
	if err != nil {
		return user, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	// gen token for user
	tokenObj, errToken := security.GenTokenObj(user)
	if errToken != nil {
		return user, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	user.Token = &tokenObj

	return user, nil
}

func (a AccountUseCase) Login(ctx context.Context, req request.LoginReq) (entity.User, error) {
	user, err := a.AccountRepository.GetUserByUseName(ctx, req.UserName)
	if err != nil {
		return user, errors.New("Người dùng không tồn tại")
	}

	//compare password.
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return entity.User{}, errors.New("Mật khẩu không chính xác")
	}

	// gen token for user
	tokenObj, errToken := security.GenTokenObj(user)
	if errToken != nil {
		return user, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	user.Token = &tokenObj

	return user, nil
}

func (a AccountUseCase) Profile(ctx context.Context, id string) (entity.User, error) {
	user, err := a.AccountRepository.GetUserByUserId(ctx, id)
	user.Password = ""
	if err != nil {
		return user, errors.New("Người dùng không tồn tại")
	}

	return user, nil
}

func (u AccountUseCase) RefreshToken(ctx context.Context, user entity.User) (entity.User, error) {
	userInfo, err := u.AccountRepository.GetUserByUserId(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New(constant.ERROR_INTERNAL_SERVER)
		}
		return entity.User{}, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	//generate token
	var token entity.Token

	jwtExpConfig := viper.Get("JWT_EXPIRED").(string)
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)
	jwtExpDuration := time.Hour * time.Duration(jwtExpValue) / time.Second
	token.ExpiredTime = jwtExpDuration

	token.AccessToken, err = security.GenToken(userInfo)
	//token.RefreshToken, err = security.GenRefreshToken(userInfo)

	userInfo.Token = &token
	userInfo.Password = ""

	return userInfo, nil
}

func (a AccountUseCase) DetailUser(ctx context.Context, id string) (entity.User, error) {
	user, err := a.AccountRepository.GetUserByUserId(ctx, id)
	user.Password = ""
	if err != nil {
		return user, errors.New("Người dùng không tồn tại")
	}

	return user, nil
}

func (a AccountUseCase) ListUser(ctx context.Context, param request.ListUser) (api_response.Pagination, error) {
	paging, err := a.AccountRepository.ListUser(ctx, param)
	if err != nil {
		return paging, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return paging, nil
}

func (a AccountUseCase) UpdateAccount(ctx context.Context, user entity.User) (entity.User, error) {
	err := a.AccountRepository.UpdateAccount(ctx, user)
	if err != nil {
		return user, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return user, nil
}

func (a AccountUseCase) ChangePass(ctx context.Context, param request.ChangePassReq) error {
	user, err := a.AccountRepository.GetUserByUserId(ctx, param.ID)
	if err != nil {
		return errors.New("Người dùng không tồn tại")
	}

	//compare password.
	isTheSame := security.ComparePasswords(user.Password, []byte(param.OldPassword))
	if !isTheSame {
		return errors.New("Mật khẩu cũ không chính xác")
	}

	user.Password = security.HashAndSalt([]byte(param.NewPassword))
	err = a.AccountRepository.UpdateAccount(ctx, user)
	if err != nil {
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return nil
}
