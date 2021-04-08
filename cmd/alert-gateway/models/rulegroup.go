package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"github.com/thertype/prom-rule/cmd/alert-gateway/logs"
)

type RuleGroups struct {
	Id          int64  `orm:"auto" json:"id,omitempty"`
	GroupName  string `orm:"column(groupname);size(255)" json:"groupname"`
	Description string `orm:"column(description);size(1023)" json:"description"`
}

func (*RuleGroups) TableName() string {
	return "rulegroup"
}

func (group *RuleGroups) GetAllRuleGroups() []RuleGroups {
	logs.Info("Models---GetAllRuleGroups---\n %s ")

	var groups []RuleGroups
	Ormer().QueryTable(new(RuleGroups)).Limit(-1).All(&groups)

	return groups
}

func (group *RuleGroups) AddRuleGroups() error {
	_, err := Ormer().Insert(group)
	logs.Info("Models---AddRuleGroups---\n %s ")
	return errors.Wrap(err, "database insert error")
}

func (group *RuleGroups) UpdateRuleGroups() error {
	_, err := Ormer().Update(group)
	return errors.Wrap(err, "database update error")
}

func (group *RuleGroups) DeleteRuleGroups(id int64) error {
	var rules []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM rule WHERE group_id = ? LOCK IN SHARE MODE", id).QueryRows(&rules) //rule绑定规则组
	if err == nil {
		if len(rules) > 0 {
			o.Commit()
			return fmt.Errorf("cannot delete this group,it is associated with following rules:%v", rules)
		} else {
			_, err = o.Raw("DELETE FROM rulegroup WHERE id = ?", id).Exec()
			if err == nil {
				_, err = o.Raw("DELETE FROM ruleunion WHERE id = ?", id).Exec()
				if err != nil {
					o.Rollback()
					return errors.Wrap(err, "database delete error")
				}
			} else {
				o.Rollback()
				return errors.Wrap(err, "database delete error")
			}
		}
	} else {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	}
	o.Commit()
	return errors.Wrap(err, "database delete error")
}

/*   //backup
func (group *RuleGroups) DeleteRuleGroups(id int64) error {
	var rules []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM rule WHERE plan_id = ? LOCK IN SHARE MODE", id).QueryRows(&rules)
	if err == nil {
		if len(rules) > 0 {
			o.Commit()
			return fmt.Errorf("cannot delete this plan,it is associated with following rules:%v", rules)
		} else {
			_, err = o.Raw("DELETE FROM plan WHERE id = ?", id).Exec()
			if err == nil {
				_, err = o.Raw("DELETE FROM plan_receiver WHERE plan_id = ?", id).Exec()
				if err != nil {
					o.Rollback()
					return errors.Wrap(err, "database delete error")
				}
			} else {
				o.Rollback()
				return errors.Wrap(err, "database delete error")
			}
		}
	} else {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	}
	o.Commit()
	return errors.Wrap(err, "database delete error")
}
 */
