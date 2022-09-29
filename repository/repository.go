package repository

import (
	"context"
	"student-proto/models"
)

type Repository interface {
	//student
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	//test
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	//questions
	SetQuestions(ctx context.Context, questions *models.Question) error
	//Enroolment
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error)
	//test-questions-answer
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
	//set answerstudent
	SetAnswerStudent(ctx context.Context, answer *models.Answer) error
	//PerTest
	GetTestScore(ctx context.Context, testId string, studentId string) (*models.TestScore, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

//student

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

//test

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

//questions

func SetQuestions(ctx context.Context, q *models.Question) error {
	return implementation.SetQuestions(ctx, q)
}

//enrollment

func SetEnrollment(ctx context.Context, enroll *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enroll)
}

func GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentPerTest(ctx, testId)
}

//test-questions-answer
func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}

//set answer student
func SetAnswerStudent(ctx context.Context, answer *models.Answer) error {
	return implementation.SetAnswerStudent(ctx, answer)
}

func GetTestScore(ctx context.Context, testId string, studentId string) (*models.TestScore, error) {
	return implementation.GetTestScore(ctx, testId, studentId)
}
