package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"github.com/thertype/prom-rule/cmd/alert-gateway/logs"
)

type RuleUnions struct {
	Id                    int64  `orm:"auto" json:"id,omitempty"`
	Plan                  *RuleGroups `orm:"index;rel(fk)" json:"regroup_id"`
//	Plan                  *RuleGroups `orm:"index;rel(fk)" json:"plan_id"`
	StartTime             string `orm:"size(31)" json:"start_time"`
	EndTime               string `orm:"size(31)" json:"end_time"`
	Start                 int    `json:"start"`
	Period                int    `json:"period"`
	Expression            string `orm:"size(1023)" json:"expression"`
	ReversePolishNotation string `orm:"size(1023)" json:"reverse_polish_notation"`
	User                  string `orm:"size(1023)" json:"user"`
	Group                 string `orm:"size(1023)" json:"group"`
	DutyGroup             string `orm:"size(255)" json:"duty_group"`
	Method                string `orm:"size(255)" json:"method"`
}


type Ruler struct {
	Id         int64  `json:"id,omitempty"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Start      int    `json:"start"`
	Period     int    `json:"period"`
	Expression string `json:"expression"`
	User       string `json:"user"`
	Group      string `json:"group"`
	DutyGroup  string `json:"duty_group"`
	Method     string `json:"method"`
}

func (*RuleUnions) TableName() string {
	return "ruleunion"
}

func (p *RuleUnions) GetAllRuleUnion(groupid string) []Ruler {
	ruleunion := []Ruler{}

	Ormer().Raw("SELECT id,start_time,end_time,start,period,expression,user,`group`,duty_group,method FROM ruleunion WHERE regroup_id=?", groupid).QueryRows(&ruleunion)
	return ruleunion
}

func (p *RuleUnions) AddRuleUnion() error {

	var groupId []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()

	_, err := o.Raw("SELECT id FROM rulegroup WHERE id = ? LOCK IN SHARE MODE", p.Plan.Id).QueryRows(&groupId)

	if err != nil {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	} else {
		if len(groupId) > 0 {
			logs.Info("Model---AddRuleUnion--Ruleunion-p: %+v\n ", p)
			logs.Info("Model---AddRuleUnion--Ruleunion-p.Plan: %+v\n ", p.Plan)


			_, err = o.Insert(p)

			logs.Info("Model---AddRuleUnion--Ruleunion-Insert:%s\n ", err)


			if err != nil {
				o.Rollback()

				//logs.Info("Model---AddRuleUnion--Debug--err-p.Plan.Id, groupId- %v,%v\n ",p.Plan.Id, groupId)
				//logs.Info("Model---AddRuleUnion--Debug--err-p-%v\n ",p)
				//logs.Info("Model---AddRuleUnion--Debug--err-p.Plan-%v\n ",p.Plan)


				return errors.Wrap(err, "database insert error-AddRuleUnion")
			}
		} else {
			o.Commit()
			return fmt.Errorf("group id: %v is not exsit", p.Plan.Id)
		}
	}
	o.Commit()
	return errors.Wrap(err, "database insert error")
}

func (p *RuleUnions) UpdateRuleUnion() error {
	_, err := Ormer().Update(p, "start_time", "end_time", "start", "period", "expression", "reverse_polish_notation", "user", "group", "duty_group", "method")
	return errors.Wrap(err, "database update error")
}

func (p *RuleUnions) DeleteRuleUnion(id string) error {
	_, err := Ormer().Raw("DELETE FROM ruleunion WHERE id=?", id).Exec()
	return errors.Wrap(err, "database delete error")
}
