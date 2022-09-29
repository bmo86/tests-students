package database //test

import (
	"context"
	"log"
	"student-proto/models"
)

//test

func (repo *PostgressRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (repo *PostgressRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var test = models.Test{}

	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
		return &test, nil
	}

	return &test, nil
}

//enrollment

func (repo *PostgressRepository) GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id IN(SELECT student_id FROM enrollment WHERE test_id = $1)", testId)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var students []*models.Student

	for rows.Next() {
		var student = models.Student{}
		if err = rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *PostgressRepository) SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollment (student_id, test_id) VALUES ($1, $2)", enrollment.StudentId, enrollment.TestId)
	return err
}

//test-questions-answer
func (repo *PostgressRepository) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, question FROM questions WHERE test_id = $1", testId)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var qs []*models.Question

	for rows.Next() {
		var question = models.Question{}
		if err = rows.Scan(&question.Id, &question.Question); err == nil {
			qs = append(qs, &question)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return qs, nil
}

//set answer student
func (repo *PostgressRepository) SetAnswerStudent(ctx context.Context, answer *models.Answer) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO studentAnswer (test_id, student_id, question_id, answer_student, correct) VALUES ($1, $2, $3, $4, $5) ", answer.TestId, answer.StudentId, answer.QuestionId, answer.AnswerStudent, answer.Correct)
	return err
}

//get testscore

func (repo *PostgressRepository) GetTestScore(ctx context.Context, testId string, studentId string) (*models.TestScore, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT correct FROM studentAnswer WHERE test_id = $1 AND student_id = $2", testId, studentId)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error in get test score -> ", err)
		}
	}()

	var answer = models.Answer{}
	var testScore = models.TestScore{
		TestId:    testId,
		StudentId: studentId,
	}

	for rows.Next() {
		err := rows.Scan(&answer.Correct)
		if err != nil {
			return nil, err
		}

		testScore.Total += 1

		if answer.Correct {
			testScore.Ok += 1
		} else {
			testScore.Ko += 1
		}

	}

	testScore.Score = testScore.Ok * 10 / testScore.Total

	return &testScore, nil

}
