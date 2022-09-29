package main

import (
	"context"
	"io"
	"log"
	"student-proto/testpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	DoUnary(c)
	//DoClientStreaming(c)
	//DoServerStreaming(c)
	//DoBidirectionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Getest: %v", err)
	}

	log.Printf("reponse from server: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	listQ := []*testpb.Question{
		{
			Id:       "q11t1",
			Answer:   "azul",
			Question: "Color favorito?",
			TestId:   "t1",
		},
		{
			Id:       "q12t1",
			Answer:   "Google",
			Question: "Empresa que creo golang? ",
			TestId:   "t1",
		},
		{
			Id:       "q13t1",
			Answer:   "backend",
			Question: "especialidad de golang?",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())

	if err != nil {
		log.Fatalf("error while calling SetQuestions: %v", err)
	}

	for _, question := range listQ {
		log.Println("sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	log.Printf("response from server: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GetStudentPerTes: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("response from server: %v", msg)

	}
}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "42",
	}

	numberOfCuestions := 4

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())

	if err != nil {
		log.Fatalf("error while calling takeTest : %v", err)
	}

	go func() {
		for i := 0; i < numberOfCuestions; i++ {
			stream.Send(&answer)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
			}

			log.Printf("response from server: %v", res)

		}

		close(waitChannel)
	}()
	<-waitChannel

}
