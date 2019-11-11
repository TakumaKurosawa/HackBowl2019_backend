package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type UserRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewUserRepoImpl(conn *sql.DB) repository.UserRepository {
	return &UserRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *UserRepositoryImpliment) InsertUserData(name string, email string, password string, authToken string) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO user(name, email, password, auth_token) VALUES(?, ?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "read")
	}
	_, err = stmt.Exec(name, email, password, authToken)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	return nil
}

func (repo *UserRepositoryImpliment) SelectByAuthToken(authToken string) (*model.User, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM user WHERE auth_token = ?", authToken)
	return convertToUser(row)
}

func (repo *UserRepositoryImpliment) SelectUserByEmail(email string) (*model.User, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM user WHERE email = ?", email)
	return convertToUser(row)
}

func (repo *UserRepositoryImpliment) SelectUserByUserId(userId string) (*model.User, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM user WHERE user_id = ?", userId)
	return convertToUser(row)
}

func (repo *UserRepositoryImpliment) SelectUsersByUserIDs(userIDs []int) ([]*model.User, error) {
	var queryStr string

	for i, _ := range userIDs {
		if i+1 == len(userIDs) {
			queryStr += fmt.Sprintf("%v", userIDs[i])
			break
		}
		queryStr += fmt.Sprintf("%v,", userIDs[i])
	}

	rows, err := repo.DbConn.Query(fmt.Sprintf("SELECT * FROM user WHERE id IN (%v)", queryStr))
	if err != nil {
		return nil, errors.Wrap(err, "DB Error")
	}

	defer rows.Close()
	return convertToUsers(rows)
}

func (repo *UserRepositoryImpliment) UpdateUserName(token, name string) error {
	stmt, err := repo.DbConn.Prepare("UPDATE user SET name = ? WHERE auth_token = ?")

	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(name, token)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	return nil
}

func (repo *UserRepositoryImpliment) UpdateUserEmail(token, email string) error {
	stmt, err := repo.DbConn.Prepare("UPDATE user SET email = ? WHERE auth_token = ?")

	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(email, token)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	return nil
}

func (repo *UserRepositoryImpliment) UpdateUserData(token, name, email string) error {
	stmt, err := repo.DbConn.Prepare("UPDATE user SET name = ?, email = ? WHERE auth_token = ?")

	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(name, email, token)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	return nil
}

func convertToUsers(rows *sql.Rows) ([]*model.User, error) {
	var results []*model.User
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email, password, authToken string

		if err := rows.Scan(&id, &name, &email, &password, &authToken); err != nil {
			return nil, errors.Wrap(err, "DB Error")
		}

		row := model.User{Id: id, Name: name, Email: email, Password: password, AuthToken: authToken}

		results = append(results, &row)
	}

	return results, nil
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*model.User, error) {
	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.AuthToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error")
		}
		return nil, errors.Wrap(err, "DB Error")
	}
	return &user, nil
}
