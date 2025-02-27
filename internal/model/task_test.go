package model_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/greencoda/tasq/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type errorReader int

func (errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

type TaskTestSuite struct {
	suite.Suite
}

func (s *TaskTestSuite) SetupTest() {
	uuid.SetRand(nil)
}

func (s *TaskTestSuite) TestNewTask() {
	// Create task successfully
	task := model.NewTask("testTask", true, "testQueue", 0, 5)
	assert.NotNil(s.T(), task)

	// Fail by creating task with nil args
	nilTask := model.NewTask("testTask", nil, "testQueue", 0, 5)
	assert.Nil(s.T(), nilTask)

	// Fail by causing uuid generation to return error
	uuid.SetRand(new(errorReader))
	invalidUUIDTask := model.NewTask("testTask", false, "testQueue", 0, 5)
	assert.Nil(s.T(), invalidUUIDTask)
}

func (s *TaskTestSuite) TestTaskGetDetails() {
	// Create task successfully
	task := model.NewTask("testTask", true, "testQueue", 0, 5)
	assert.NotNil(s.T(), task)

	// Get task details
	taskModel := task.GetDetails()
	assert.IsType(s.T(), &model.Task{}, taskModel)
}

func (s *TaskTestSuite) TestTaskUnmarshalArgs() {
	// Create task successfully
	task := model.NewTask("testTask", true, "testQueue", 0, 5)
	assert.NotNil(s.T(), task)

	// Unmarshal task args successfully
	var args bool
	err := task.UnmarshalArgs(&args)
	assert.Nil(s.T(), err)
	assert.True(s.T(), args)

	// Fail by unmarshaling args to incorrect type
	var incorrectTypeArgs string
	err = task.UnmarshalArgs(&incorrectTypeArgs)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), incorrectTypeArgs)
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}
