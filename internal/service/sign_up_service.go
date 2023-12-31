package service

import (
	"GoSpace/internal/common"
	"GoSpace/internal/dao"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"GoSpace/internal/util"
)

type SignUpService struct {
}

func (s *SignUpService) SignUp(email, phone, pwd string) (id int64, err error) {
	if len(email) == 0 && len(phone) == 0 {
		return 0, errno.NewErrorNo(nil, common.ErrUnknownEmailAndPhone)
	}

	if len(email) > 0 && !util.IsEmailValid(&email) {
		return 0, errno.NewErrorNo(nil, common.ErrNotValidEmail)
	}
	if len(phone) > 0 && !util.IsPhoneValid(&phone) {
		return 0, errno.NewErrorNo(nil, common.ErrNotValidPhone)
	}
	if len(pwd) > 0 && !util.IsPasswordValid(&pwd) {
		return 0, errno.NewErrorNo(nil, common.ErrNotValidPwd)
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
		return 0, errno.NewErrorNo(err, common.ErrUtilHashSalt)
	}

	if err = dao.PgDAO.CreateUser(&user); err != nil {
		return 0, err
	}

	return user.Id, nil
}
