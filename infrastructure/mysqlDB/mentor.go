package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type MentorRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewMentorRepoImpl(conn *sql.DB) repository.MentorRepository {
	return &MentorRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *MentorRepositoryImpliment) InsertMentorData(name, email, password string) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO mentor(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "read")
	}
	_, err = stmt.Exec(name, email, password)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	return nil
}

func (repo *MentorRepositoryImpliment) SelectMentorById(mentorId int) (*model.Mentor, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM mentor WHERE id = ?", mentorId)
	return convertToMentor(row)
}

func (repo *MentorRepositoryImpliment) SelectMentorsByMentorIDs(mentorIDs []int) ([]*model.Mentor, error) {
	var queryStr string

	for i, _ := range mentorIDs {
		if i+1 == len(mentorIDs) {
			queryStr += fmt.Sprintf("%v", mentorIDs[i])
			break
		}
		queryStr += fmt.Sprintf("%v,", mentorIDs[i])
	}

	rows, err := repo.DbConn.Query(fmt.Sprintf("SELECT * FROM mentor WHERE id IN (%v)", queryStr))
	if err != nil {
		return nil, errors.Wrap(err, "DB Error")
	}

	defer rows.Close()
	return convertToMentors(rows)
}

// convertToMentor rowデータをMentorデータへ変換する
func convertToMentor(row *sql.Row) (*model.Mentor, error) {
	mentor := model.Mentor{}
	err := row.Scan(&mentor.Id, &mentor.Name, &mentor.Email, &mentor.Password, &mentor.AuthToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error")
		}
		return nil, errors.Wrap(err, "DB Error")
	}

	log.Println(mentor)
	return &mentor, nil
}

func convertToMentors(rows *sql.Rows) ([]*model.Mentor, error) {
	var results []*model.Mentor
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email, password, auth_token string

		if err := rows.Scan(&id, &name, &email, &password, &auth_token); err != nil {
			return nil, errors.Wrap(err, "DB Error")
		}

		row := model.Mentor{Id: id, Name: name, Email: email, Password: password, AuthToken: auth_token}

		results = append(results, &row)
	}

	for i, _ := range results {
		log.Println(results[i])
	}

	return results, nil
}
