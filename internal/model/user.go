package model

/*
-- 创建一个名为 user 的表格
CREATE TABLE user_basic (
  -- 用户 id，bigint 类型，主键，不为空，生成方式为 identity
  id bigint PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
  -- 盐值，字符串类型，不为空
  salt VARCHAR NOT NULL,
  -- 加盐哈希后的密码，字符串类型，不为空
  hash_pwd VARCHAR NOT NULL,
  -- 邮箱，字符串类型，可以为空，唯一
  email VARCHAR UNIQUE,
  -- 手机号，字符串类型，可以为空，唯一
  phone VARCHAR UNIQUE,
  -- 状态，整数类型，不为空，默认为0
  status INTEGER NOT NULL DEFAULT 0
);
-- 修改序列的当前值为 10000000
SELECT setval('user_basic_id_seq', 10000000);
*/

type UserBasic struct {
	tableName struct{} `pg:"user_basic"`
	// 用户 id，bigint 类型，主键，不为空，生成方式为 identity
	Id int64 `pg:",pk,notnull"`
	// 盐值，字符串类型，不为空
	Salt string `pg:",notnull"`
	// 加盐哈希后的密码，字符串类型，不为空
	HashPwd string `pg:",notnull"`
	// 邮箱，字符串类型，可以为空，唯一
	Email string `pg:",unique"`
	// 手机号，字符串类型，可以为空，唯一
	Phone string `pg:",unique"`
	// 状态，整数类型，不为空，默认为0
	Status int `pg:",notnull,use_zero"`
}
