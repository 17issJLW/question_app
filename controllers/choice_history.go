package controllers

import (
	"github.com/astaxie/beego"
	"test/models"
)

// ChoiceQuestionController operations for ChoiceQuestion
type ChoiceHistoryController struct {
	beego.Controller
}

func (c *ChoiceHistoryController) Get(){
	qid,_ := c.GetInt(":qid")
	beego.Info(qid)
	history := models.GetOrCreateChoiceHistory(qid)
	//c.Data["json"] = map[string]string{"a":"hello","qid":  strconv.Itoa(b)}
	c.Data["json"] = history
	c.ServeJSON()
}


func (c *ChoiceHistoryController) Post(){
	qid,_ := c.GetInt(":qid")
	choice := c.GetString("choice")
	if choice != ""{
		success := models.ChangeChoiceHistory(qid,choice)
		if success{
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]string{"success":"success"}

		}else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = "err"

		}
	}else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"success":"success"}
	}

}