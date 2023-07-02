package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-iris-sample/_example-mvc-api/model"
)

// UserService ユーザーサービス
type UserService struct{}

// GetUserList ユーザー一覧取得
func (s UserService) GetUserList() ([]model.User, error) {
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

	var users []model.User

	_, err := dbMap.Select(&users, `SELECT * FROM users`)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

// CreateUser ユーザー作成
func (s UserService) CreateUser(user *model.User) error {
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

	// トランザクション開始
	tx, _ := dbMap.Begin()

	err := tx.Insert(user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// UpdateUser ユーザー更新
func (s UserService) UpdateUser(id int, user *model.User) error {
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

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

	// トランザクション開始
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
			"age":  user.Age,
			"id":   id,
		})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// DeleteUser ユーザー削除
func (s UserService) DeleteUser(id int) error {
	dbMap := s.InitDb()
	defer dbMap.Db.Close()

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

	// トランザクション開始
	tx, _ := dbMap.Begin()

	_, err = tx.Delete(&user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
