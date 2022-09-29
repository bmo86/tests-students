package database

import (
	"context"
	"student-proto/models"
)

//questions

func (repo *PostgressRepository) SetQuestions(ctx context.Context, q *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, answer, question, test_id) VALUES($1, $2, $3, $4)", q.Id, q.Answer, q.Question, q.TestId)
	return err
}
