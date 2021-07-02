package hcs

import (
	"github.com/gin-gonic/gin"
	"hcs/internal"
)

func PutAgent(c *gin.Context) {
	p := new(internal.PutAgentRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	res, err := internal.PutAgent(p)
	internal.Return(c, internal.PutAgentResponse{AgentID: res}, err)
}

func Heart(c *gin.Context) {
	p := new(internal.HeartRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	status, err := internal.Heart(p)
	internal.Return(c, internal.HeartResponse{Status: status}, err)
}

func GetTask(c *gin.Context) {
	p := new(internal.GetTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	res, err := internal.GetTask(p)
	internal.Return(c, internal.GetTaskResponse{Task: res}, err)
}

func PutTask(c *gin.Context) {
	p := new(internal.PutTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	err := internal.PutTask(p)
	internal.Return(c, "", err)
}
