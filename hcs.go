package hcs

import (
	"github.com/gin-gonic/gin"
	"hcs/controller"
	"hcs/dao"
	"hcs/manager"
)

var (
	agentCtl     *controller.AgentCtl
	taskCtl      *controller.TaskCtl
	agentTaskCtl *controller.AgentTaskCtl
)

func WrapGroup(router *gin.RouterGroup) {
	agentDao := new(dao.AgentDao)
	taskDao := new(dao.TaskDao)
	agentTaskDao := new(dao.AgentTaskDao)

	agentManager := manager.NewAgentManager(agentDao, agentTaskDao)
	taskManager := manager.NewTaskManager(taskDao)
	agentTaskManager := manager.NewAgentTaskManager(agentTaskDao, taskDao)

	agentCtl = controller.NewAgentCtl(agentManager)
	taskCtl = controller.NewTaskCtl(taskManager)
	agentTaskCtl = controller.NewAgentTaskCtl(agentTaskManager)

	{
		t := router.Group("tasks")

		t.POST("/task", taskCtl.AddTask)               // 新建任务
		t.DELETE("/task/:task_id", taskCtl.DeleteTask) // 删除任务
		t.GET("/all", taskCtl.GetAll)                  // 获取任务列表
	}

	{
		a := router.Group("agents")

		a.POST("/agent", agentCtl.AddAgent)       // agent注册
		a.GET("/heart/:agent_id", agentCtl.Heart) // agent心跳
		a.GET("/all", agentCtl.GetAll)            // 获取agent列表
	}

	{
		at := router.Group("relations")

		at.POST("/relation", agentTaskCtl.AddRelation)                       // 给agent增加任务
		at.DELETE("/relation/:relation_id", agentTaskCtl.DeleteRelation)     // 获取任务列表
		at.GET("/all", agentTaskCtl.GetAll)                                  // 获取任务列表
		at.POST("/relation/status", agentTaskCtl.PostStatus)                 // agent回写任务状态
		at.GET("/relation/agent/:agent_id", agentTaskCtl.GetRelationByAgent) // 按照agent查询任务关系列表
		//at.GET("/relation/task/:task_id", agentTaskCtl.GetRelationByTask)    // 按照task查询任务关系列表
	}
}

type MySql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Hostport string `json:"hostport"`
	Database string `json:"database"`
}

func MustInitDB(db MySql) {
	dao.MustInitDB(db.Username, db.Password, db.Hostname, db.Hostport, db.Database)
}
