package dao

import (
	"GoSpace/internal/client"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"GoSpace/internal/util"
	"github.com/go-pg/pg/v10"
)

type PgDAO struct {
}

// CreateUser 创建用户
// 需要至少传递邮箱/手机号其中之一；入参HashPwd是明文密码，在该函数生成盐，并将明文密码加盐，哈希加密后替换
func (p *PgDAO) CreateUser(user *model.UserBasic) (err error) {
	if len(user.Email) == 0 && len(user.Phone) == 0 {
		return errno.NewErrorNo(nil, errno.ErrSignUpUnknownEmailOrPhone)
	}
	user.Salt = util.RandStringBytesMaskImprSrcUnsafe(8)
	user.HashPwd, err = util.HashSalt(user.HashPwd, user.Salt)
	if err != nil {
		return errno.NewErrorNo(err, errno.ErrUtilHashSalt)
	}

	clt := client.ClientManager.GetPgClient()
	_, err = clt.DB.Model(user).Insert(user)
	if err != nil {
		switch err.(pg.Error).Field('C') {
		case "23505":
			return errno.NewErrorNo(err, errno.ErrSignUpEmailOrPhoneDuplicate)
		}
	}
	return nil
}

func (p *PgDAO) GetUserBasic(user *model.UserBasic) error {
	if len(user.Email) > 0 {
		return p.GetUserBasicByEmail(user)
	} else if len(user.Phone) > 0 {
		return p.GetUserBasicByPhone(user)
	}
	return errno.NewErrorNo(nil, errno.ErrSelectUserUnknownEmailOrPhone)
}

func (p *PgDAO) GetUserBasicByEmail(user *model.UserBasic) error {
	if len(user.Email) == 0 {
		return errno.NewErrorNo(nil, errno.ErrSelectUserUnknownEmailOrPhone)
	}
	clt := client.ClientManager.GetPgClient()
	err := clt.DB.Model(user).
		Where("email = ?", user.Email).
		Select()
	if err != nil {
		switch err.Error() {
		case "pg: no rows in result set":
			return errno.NewErrorNo(err, errno.ErrSelectUserByEmailEmpty)
		default:
			return err
		}
	}
	return nil
}

func (p *PgDAO) GetUserBasicByPhone(user *model.UserBasic) error {
	if len(user.Phone) == 0 {
		return errno.NewErrorNo(nil, errno.ErrSelectUserUnknownEmailOrPhone)
	}
	clt := client.ClientManager.GetPgClient()
	err := clt.DB.Model(user).
		Where("phone = ?", user.Phone).
		Select()
	if err != nil {
		switch err.Error() {
		case "pg: no rows in result set":
			return errno.NewErrorNo(err, errno.ErrSelectUserByPhoneEmpty)
		default:
			return err
		}
	}
	return nil
}

// UpdateUserBasicByPk 更新用户数据
// 根据主键找不到也不会报错
func (p *PgDAO) UpdateUserBasicByPk(user *model.UserBasic) error {
	if user.Id <= 0 {
		return errno.NewErrorNo(nil, errno.ErrSelectUserById)
	}
	clt := client.ClientManager.GetPgClient()
	_, err := clt.DB.Model(user).WherePK().Update(user)
	return err
}

// SetUserStatus 设置用户状态
func (p *PgDAO) SetUserStatus(user *model.UserBasic, status int) error {
	user.Status = status
	return p.UpdateUserBasicByPk(user)
}

// NormalizeUser 恢复用户正常状态
func (p *PgDAO) NormalizeUser(user *model.UserBasic) error {
	if user.Status != model.UserStatusNormal {
		return p.SetUserStatus(user, model.UserStatusNormal)
	}
	return nil
}

// DisableUser 封禁用户
func (p *PgDAO) DisableUser(user *model.UserBasic) error {
	if user.Status != model.UserStatusDisabled {
		return p.SetUserStatus(user, model.UserStatusDisabled)
	}
	return nil
}

// LogicDeleteUser 逻辑删除用户
func (p *PgDAO) LogicDeleteUser(user *model.UserBasic) error {
	if user.Status != model.UserStatusDeleted {
		return p.SetUserStatus(user, model.UserStatusDeleted)
	}
	return nil
}
