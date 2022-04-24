package service

import (
	"fmt"
	"go-iris-sample/_example-mvc-api/model"

	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {}

func (UserService) GetUserList() ([]model.User, error) {
	dbMap := initDb()

	fmt.Println(456)

	var users []model.User

	// ユーザーを全取得
	_, err := dbMap.Select(&users, `SELECT * FROM users`)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (UserService) CreateUser(user *model.User) error {
	dbMap := initDb()

	// トランザクションを走らせながらinsert
	tx, _ := dbMap.Begin()

	err := tx.Insert(user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (UserService) UpdateUser(user *model.User) error {
	dbMap := initDb()

	// トランザクションを走らせながらupdate
	tx, _ := dbMap.Begin()

	_, err := tx.Update(user)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (UserService) DeleteUser(id string) error {
	dbMap := initDb()

	// トランザクションを走らせながらdelete
	tx, _ := dbMap.Begin()

	_, err := tx.Delete(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
