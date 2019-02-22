package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type ChoiceQuestion struct {
	Id       int    `orm:"column(id);auto"`
	Question string `orm:"column(question);index"`
	OptionA  string `orm:"column(option_a);size(128);null"`
	OptionB  string `orm:"column(option_b);size(128);null"`
	OptionC  string `orm:"column(option_c);size(128);null"`
	OptionD  string `orm:"column(option_d);size(128);null"`
	OptionE  string `orm:"column(option_e);size(128);null"`
	Answer   string `orm:"column(answer);size(32);null"`
	Reason   string `orm:"column(reason);null"`
	Course   string `orm:"column(course);size(32);null"`
	Chapter  string `orm:"column(chapter);size(32);null"`
	KeyPoint string `orm:"column(key_point);size(64);null"`
	//AddTime orm.DateTimeField `orm:"column(add_time);auto_now_add;type(datetime)"`
}

func (t *ChoiceQuestion) TableName() string {
	return "choice_question"
}

func init() {
	//orm.RegisterDataBase("default", "mysql", "root:jlw123456@tcp(127.0.0.1:3306)/app?charset=utf8")
	orm.RegisterModel(new(ChoiceQuestion))


}

// AddChoiceQuestion insert a new ChoiceQuestion into database and returns
// last inserted Id on success.
func AddChoiceQuestion(m *ChoiceQuestion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChoiceQuestionById retrieves ChoiceQuestion by Id. Returns error if
// Id doesn't exist
func GetChoiceQuestionById(id int) (v *ChoiceQuestion, err error) {
	o := orm.NewOrm()
	v = &ChoiceQuestion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChoiceQuestion retrieves all ChoiceQuestion matches certain condition. Returns empty list if
// no records exist
func GetAllChoiceQuestion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ChoiceQuestion))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ChoiceQuestion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateChoiceQuestion updates ChoiceQuestion by Id and returns error if
// the record to be updated doesn't exist
func UpdateChoiceQuestionById(m *ChoiceQuestion) (err error) {
	o := orm.NewOrm()
	v := ChoiceQuestion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChoiceQuestion deletes ChoiceQuestion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChoiceQuestion(id int) (err error) {
	o := orm.NewOrm()
	v := ChoiceQuestion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ChoiceQuestion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
