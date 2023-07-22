package dao

import (
	"GoSpace/internal/client"
	"GoSpace/internal/errno"
	"GoSpace/internal/model"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type PgDAO struct {
}

func (p *PgDAO) CreateUser(user *model.UserBasic) error {
	clt := client.ClientManager.GetPgClient()
	_, err := clt.DB.Model(user).Insert(user)
	if err != nil {
		switch err.(pg.Error).Field('C') {
		case "23505":
			return errno.NewErrorNo(err, errno.ErrSignUpEmailOrPhoneDuplicate)
		}
	}
	return nil
}

func (p *PgDAO) GetUser(user *model.UserBasic) error {
	if len(user.Email) > 0 {
		return p.GetUserByEmail(user)
	} else if len(user.Phone) > 0 {
		return p.GetUserByPhone(user)
	}
	return errno.NewErrorNo(nil, errno.ErrSelectUserUnknownEmailOrPhone)
}

func (p *PgDAO) GetUserByEmail(user *model.UserBasic) error {
	clt := client.ClientManager.GetPgClient()
	fmt.Println(user.Email)
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

func (p *PgDAO) GetUserByPhone(user *model.UserBasic) error {
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

// UpdateUserByPk 设置用户状态
// 根据主键找不到也不会报错
func (p *PgDAO) UpdateUserByPk(user *model.UserBasic) error {
	if user.Id <= 0 {
		return errno.NewErrorNo(nil, errno.ErrSelectUserById)
	}
	clt := client.ClientManager.GetPgClient()
	_, err := clt.DB.Model(user).WherePK().Update(user)
	return err
}
