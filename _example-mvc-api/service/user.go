package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-iris-sample/_example-mvc-api/model"
)

type UserService struct {}

func (s UserService) GetUserList() ([]model.User, error) {
	// initialize the DbMap
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

	// ユーザーを全取得
	var users []model.User

	_, err := dbMap.Select(&users, `SELECT * FROM users`)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (s UserService) CreateUser(user *model.User) error {
	// initialize the DbMap
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

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

func (s UserService) UpdateUser(id int, user *model.User) error {
	// initialize the DbMap
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

	// id からユーザーが実在するか確認する
	var isExistUser model.User

	err := dbMap.SelectOne(&isExistUser,
		`SELECT
					*
				FROM
					users
				WHERE
					id = :id`,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		fmt.Printf("error! can't find user by id: %v.\n", id)
		return err
	}

	// トランザクションを走らせながらupdate
	tx, _ := dbMap.Begin()

	_, err = tx.Exec(
		`UPDATE
					users
				SET
					name = :name,
					age = :age
				WHERE
					id = :id`,
		map[string]interface{}{
			"name": user.Name,
			"age": user.Age,
			"id": id,
		})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (s UserService) DeleteUser(id int) error {
	// initialize the DbMap
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

	// id から削除するユーザーを取得
	var user model.User

	err := dbMap.SelectOne(&user,
		`SELECT
					*
				FROM
					users
				WHERE
					id = :id`,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		fmt.Printf("error! can't find user by id: %v.\n", id)
		return err
	}

	// トランザクションを走らせながらdelete
	tx, _ := dbMap.Begin()

	_, err = tx.Delete(&user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
