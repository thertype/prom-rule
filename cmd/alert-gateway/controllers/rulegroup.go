package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type RuleGroupController struct {
	beego.Controller
}

func (c *RuleGroupController) URLMapping() {
	c.Mapping("GetAllRuel", c.GetAllRule)
	c.Mapping("AddRule", c.AddRule)
	c.Mapping("GetAllGroup", c.GetAllGroup)
	c.Mapping("AddGroup", c.AddGroup)
	c.Mapping("UpdateGroup", c.UpdateGroup)
	c.Mapping("DeleteGroup", c.DeleteGroup)
}

// @router / [get]
func (c *RuleGroupController) GetAllGroup() {
	var Ruleunion *models.RuleGroups
	groups := Ruleunion.GetAllGroups()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: groups,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *RuleGroupController) AddGroup() {
	var group models.RuleGroups
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = group.AddGroups()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, group)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:groupid/ruleunion/ [get]
func (c *RuleGroupController) GetAllRule() {
	groupId := c.Ctx.Input.Param(":groupid")
	var Ruleunion *models.Ruleunion
	ruleunion := Ruleunion.GetAllRuleUnion(groupId)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: ruleunion,
	}
	c.ServeJSON()
}

// @router /:groupid/ruleunion/ [post]
func (c *RuleGroupController) AddRule() {
	groupId := c.Ctx.Input.Param(":groupid")
	var Ruleunion *models.Ruleunion
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Ruleunion)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		if Ruleunion.Expression != "" {
			root, err := common.BuildTree(Ruleunion.Expression)
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			} else {
				ReversePolishNotation := common.Converse2ReversePolishNotation(root)
				Ruleunion.ReversePolishNotation = ReversePolishNotation
				id, _ := strconv.ParseInt(groupId, 10, 64)
				Ruleunion.Plan = &models.RuleGroups{Id: id}   //need edit
				err = Ruleunion.AddRuleUnion()
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(groupId, 10, 64)
			Ruleunion.Plan = &models.RuleGroups{Id: id}
			err = Ruleunion.AddRuleUnion()
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			}
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, Ruleunion)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:groupid [put]
func (c *RuleGroupController) UpdateGroup() {
	var group models.RuleGroups
	groupId := c.Ctx.Input.Param(":groupid")
	id, _ := strconv.ParseInt(groupId, 10, 64)
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err == nil {
		group.Id = id
		err = group.UpdateGroups()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, group)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:groupid [delete]
func (c *RuleGroupController) DeleteGroup() {
	groupId := c.Ctx.Input.Param(":groupid")
	id, _ := strconv.ParseInt(groupId, 10, 64)
	var Ruleunion *models.RuleGroups
	var ans common.Res
	err := Ruleunion.DeleteGroups(id)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, groupId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
