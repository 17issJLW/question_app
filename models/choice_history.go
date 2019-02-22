package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ChoiceHistory struct {
	Id int `orm:"pk;auto"`
	Question_id int `orm:"index;unique"`
	A int	`orm:"default(0)"`
	B int	`orm:"default(0)"`
	C int	`orm:"default(0)"`
	D int	`orm:"default(0)"`
	E int	`orm:"default(0)"`
	Right float32	`orm:"default(0)"`
}

func init()  {
	orm.RegisterModel(new(ChoiceHistory))
}

func (t *ChoiceHistory) TableName() string {
	return "choice_history"
}

func GetOrCreateChoiceHistory(qid int) (v *ChoiceHistory){
	history := &ChoiceHistory{Question_id:qid}
	o := orm.NewOrm()
	if _, _, err := o.ReadOrCreate(history,"Question_id");err == nil{
		return history
	}else {
		beego.Error(err)
		return nil
	}
}

func ChangeChoiceHistory(qid int,choice string) (sucess bool){
	choice_map := map[string]string{"a":"A","b":"B","c":"C","d":"D"}
	history := &ChoiceHistory{Question_id:qid}
	o := orm.NewOrm()
	//qs := o.QueryTable(new(ChoiceHistory))

	if _, _, err := o.ReadOrCreate(history,"Question_id");err == nil{
		if value,ok := choice_map[choice]; ok{
			if choice=="a"{
				history.A = history.A + 1
			}else if choice=="b"{
				history.B = history.B + 1
			}else if choice=="c"{
				history.C = history.C + 1
			}else if choice=="d"{
				history.D = history.D + 1
			}
			_,err2 := o.Update(history,value)
			if err2 == nil{
				return true
			}
			return false
		}else {
			beego.Error("ChangeChoiceHistory error: param wrong")
			return false
		}
	}
	return false
}