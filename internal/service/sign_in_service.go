package service

import (
	"GoSpace/internal/dao"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"GoSpace/internal/util"
)

type SignInService struct {
}

func (s *SignInService) SignIn(id int64, email, phone, pwd string) (bool, error) {
	if len(email) == 0 && len(phone) == 0 {
		return false, errno.NewErrorNo(nil, errno.ErrSelectUserUnknownEmailAndPhone)
	}
	user := model.UserBasic{
		Id:    id,
		Email: email,
		Phone: phone,
	}

	if err := dao.PgDAO.GetUserBasic(&user); err != nil {
		return false, err
	}

	tryHashPwd, err := util.HashSalt(pwd, user.Salt)
	if err != nil {
		return false, errno.NewErrorNo(err, errno.ErrUtilHashSalt)
	}

	if user.HashPwd == tryHashPwd {
		return true, nil
	}

	return false, errno.NewErrorNo(nil, errno.ErrSignInWrongPwd)
}
