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
		return false, errno.NewErrorNo(nil, errno.ErrUnknownEmailAndPhone)
	}
	if len(email) > 0 && !util.IsEmailValid(&email) {
		return false, errno.NewErrorNo(nil, errno.ErrNotValidEmail)
	}
	if len(phone) > 0 && !util.IsPhoneValid(&phone) {
		return false, errno.NewErrorNo(nil, errno.ErrNotValidPhone)
	}
	if len(pwd) > 0 && !util.IsPasswordValid(&pwd) {
		// 如果升级密码要求，这里就不能再校验，避免老密码无法登陆
		return false, errno.NewErrorNo(nil, errno.ErrSignInWrongAccountOrPwd)
	}

	user := model.UserBasic{
		Id:    id,
		Email: email,
		Phone: phone,
	}

	if err := dao.PgDAO.GetUserBasic(&user); err != nil {
		return false, errno.NewErrorNo(err, errno.ErrNotExistEmailAndPhone)
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
