package main

import (
	"fmt"

	v1 "github.com/Asliddin3/exam-api-gateway/api/handlers/v1"
	"github.com/google/uuid"
)

func main() {
	v1.EmailVerification("salom", "adsfasdf", "adsfasdf")
	id := uuid.New()
	fmt.Println(id.String())
}
