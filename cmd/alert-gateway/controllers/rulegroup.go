package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/thertype/prom-rule/cmd/alert-gateway/common"
	"github.com/thertype/prom-rule/cmd/alert-gateway/logs"
	//"github.com/thertype/prom-rule/cmd/alert-gateway/models"
	"github.com/thertype/prom-rule/cmd/alert-gateway/models"
)

type RuleGroupController struct {
	beego.Controller
}

func (c *RuleGroupController) URLMapping() {
	c.Mapping("GetAllRuleUnion", c.GetAllRuleUnion)
	c.Mapping("AddRuleUnion", c.AddRuleUnion)
	c.Mapping("GetAllRuleGroup", c.GetAllRuleGroup)
	c.Mapping("AddRuleGroup", c.AddRuleGroup)
	c.Mapping("UpdateRuleGroup", c.UpdateRuleGroup)
	c.Mapping("DeleteRuleGroup", c.DeleteRuleGroup)
}

// @router / [get]
func (c *RuleGroupController) GetAllRuleGroup() {
	logs.Info("RuleGroupController---GetAllGroup---\n %s ")
	var Ruleunion *models.RuleGroups
	groups := Ruleunion.GetAllRuleGroups()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: groups,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *RuleGroupController) AddRuleGroup() {
	var group models.RuleGroups
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = group.AddRuleGroups()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, group)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:groupid/reunion/ [get]
func (c *RuleGroupController) GetAllRuleUnion() {
	groupId := c.Ctx.Input.Param(":groupid")
	var Ruleunion *models.RuleUnions
	ruleunion := Ruleunion.GetAllRuleUnion(groupId)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: ruleunion,
	}
	c.ServeJSON()
}

// @router /:groupid/reunion/ [post]
func (c *RuleGroupController) AddRuleUnion() {
//	groupId := c.Ctx.Input.Param(":groupid")
	groupId := c.Ctx.Input.Param(":groupid")

	var Ruleunion *models.RuleUnions
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Ruleunion)

	logs.Info("RuleGroupController---AddRuleUnion--c.Ctx.Input.RequestBody- %s\n ", c.Ctx.Input.RequestBody)
	logs.Info("RuleGroupController---AddRuleUnion--Ruleunion- %+v\n ", Ruleunion)


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
				Ruleunion.Plan = &models.RuleGroups{Id: id}
				err = Ruleunion.AddRuleUnion()
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(groupId, 10, 64)
			Ruleunion.Plan = &models.RuleGroups{Id: id}
			//logs.Info("RuleGroupController---AddRuleUnion--Ruleunion- %v\n ", Ruleunion)

			//logs.Info("RuleGroupController---AddRuleUnion--else- %v\n ", Ruleunion.Plan)
			err = Ruleunion.AddRuleUnion()
			if err != nil {
				//logs.Info("RuleGroupController---AddRuleUnion--ERROR- %v\n ", id)
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
func (c *RuleGroupController) UpdateRuleGroup() {
	var group models.RuleGroups
	groupId := c.Ctx.Input.Param(":groupid")
	id, _ := strconv.ParseInt(groupId, 10, 64)
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err == nil {
		group.Id = id
		err = group.UpdateRuleGroups()
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
func (c *RuleGroupController) DeleteRuleGroup() {
	groupId := c.Ctx.Input.Param(":groupid")
	id, _ := strconv.ParseInt(groupId, 10, 64)
	var Ruleunion *models.RuleGroups
	var ans common.Res
	err := Ruleunion.DeleteRuleGroups(id)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, groupId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
