package service

import (
	"GoSpace/internal/common"
	"GoSpace/internal/dao"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"GoSpace/internal/util"
)

type SignInService struct {
}

func (s *SignInService) SignIn(id int64, email, phone, pwd string) (bool, int, error) {
	if len(email) == 0 && len(phone) == 0 {
		return false, -1, errno.NewErrorNo(nil, common.ErrUnknownEmailAndPhone)
	}
	if len(email) > 0 && !util.IsEmailValid(&email) {
		return false, -1, errno.NewErrorNo(nil, common.ErrNotValidEmail)
	}
	if len(phone) > 0 && !util.IsPhoneValid(&phone) {
		return false, -1, errno.NewErrorNo(nil, common.ErrNotValidPhone)
	}
	if len(pwd) > 0 && !util.IsPasswordValid(&pwd) {
		// 如果升级密码要求，这里就不能再校验，避免老密码无法登陆
		return false, -1, errno.NewErrorNo(nil, common.ErrSignInWrongAccountOrPwd)
	}

	user := model.UserBasic{
		Id:    id,
		Email: email,
		Phone: phone,
	}

	if err := dao.PgDAO.GetUserBasic(&user); err != nil {
		return false, -1, errno.NewErrorNo(err, common.ErrNotExistEmailAndPhone)
	}

	tryHashPwd, err := util.HashSalt(pwd, user.Salt)
	if err != nil {
		return false, -1, errno.NewErrorNo(err, common.ErrUtilHashSalt)
	}

	if user.HashPwd == tryHashPwd {
		if user.Status == common.UserStatusNormal {
			return true, common.UserStatusNormal, nil
		}
		return false, user.Status, errno.NewErrorNo(nil, common.ErrSignInUnNormalStatus)
	}

	return false, -1, errno.NewErrorNo(nil, common.ErrSignInWrongPwd)
}
