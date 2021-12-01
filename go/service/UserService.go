package service

import (
	dao2 "gindemo/go/dao"
	entity2 "gindemo/go/entity"
)

func Create(user *entity2.User) (err error) {
	cerr := dao2.SqlSession.Create(user).Error
	if cerr != nil {
		return cerr
	}
	return
}

func GetAllUser() (userList []*entity2.User, err error) {
	err = dao2.SqlSession.Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserById(id string) (user *entity2.User, err error) {
	err = dao2.SqlSession.Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}
