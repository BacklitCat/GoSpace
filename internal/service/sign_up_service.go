package service

import (
	"GoSpace/internal/dao"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"GoSpace/internal/util"
)

type SignUpService struct {
}

func (s *SignUpService) SignUp(email, phone, pwd string) (id int64, err error) {
	if len(email) == 0 && len(phone) == 0 {
		return 0, errno.NewErrorNo(nil, errno.ErrSignUpUnknownEmailAndPhone)
	}

	user := model.UserBasic{
		HashPwd: pwd,
		Email:   email,
		Phone:   phone,
		Status:  0,
	}

	user.Salt = util.RandStringBytesMaskImprSrcUnsafe(8)
	user.HashPwd, err = util.HashSalt(user.HashPwd, user.Salt)
	if err != nil {
		return 0, errno.NewErrorNo(err, errno.ErrUtilHashSalt)
	}

	if err = dao.PgDAO.CreateUser(&user); err != nil {
		return 0, err
	}

	return user.Id, nil
}