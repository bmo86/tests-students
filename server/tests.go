package server

import (
	"context"
	"io"
	"log"
	"student-proto/models"
	"student-proto/repository"
	"student-proto/studentpb"
	"student-proto/testpb"
	"time"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	err := s.repo.SetTest(ctx, test)

	if err != nil {
		return nil, err
	}

	return &testpb.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		question := models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}

		err = s.repo.SetQuestions(context.Background(), &question)

		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}

		err = s.repo.SetEnrollment(context.Background(), enrollment)

		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentPerTest(context.Background(), req.GetTestId())

	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}

		err := stream.Send(student)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error receiving test - id: %v", err)
			return err
		}

		questions, err := s.repo.GetQuestionsPerTest(context.Background(), msg.TestId)
		if err != nil {
			return err
		}

		i := 0

		var currentQuestion = &models.Question{}
		for {

			if i < len(questions) {
				currentQuestion = questions[i]
				questionToSend := &testpb.QuestionsPerTest{
					Id:       currentQuestion.Id,
					Question: currentQuestion.Question,
					Ok:       false,
					Current:  int32(i + 1),
					Total:    int32(len(questions)),
				}

				err := stream.Send(questionToSend)

				if err != nil {
					return err
				}

				i++

				answer, err := stream.Recv()

				if err == io.EOF {
					return nil
				}

				if err != nil {
					log.Printf("Error receiving answer: %v", err)
					return err
				}

				log.Println("Answer for question:", currentQuestion.Question, "is", answer.GetAnswer())

				//set answer student

				modelAnswerStudent := models.Answer{
					TestId:        answer.GetTestId(),
					StudentId:     answer.GetStudentId(),
					QuestionId:    currentQuestion.Id,
					AnswerStudent: answer.GetAnswer(),
					Correct:       (currentQuestion.Answer == answer.GetAnswer()),
				}

				err = s.repo.SetAnswerStudent(context.Background(), &modelAnswerStudent)

				if err != nil {
					log.Println("Error in saving answer-student", err)
					return err
				}

			} else {
				questionToSend := &testpb.QuestionsPerTest{
					Id:       "",
					Question: "",
					Ok:       true,
					Current:  int32(0),
					Total:    int32(0),
				}

				err := stream.Send(questionToSend)

				if err == io.EOF {
					return nil
				}

				if err != nil {
					return err
				}

				break

			}

		}
	}
}

func (s *TestServer) GetTestScore(ctx context.Context, req *testpb.GetScoreTestRequest) (*testpb.TestScore, error) {
	ts, err := s.repo.GetTestScore(ctx, req.TestId, req.StudentId)

	if err != nil {
		log.Printf("Error in GetTestScore -> %v", err)
		return nil, err
	}

	return &testpb.TestScore{
		TestId:    ts.TestId,
		StudentId: ts.StudentId,
		Ok:        ts.Ok,
		Ko:        ts.Ko,
		Total:     ts.Total,
		Score:     ts.Score,
	}, nil

}
