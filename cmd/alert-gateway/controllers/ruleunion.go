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

type RuleUnionController struct {
	beego.Controller
}

func (c *RuleUnionController) URLMapping() {
	c.Mapping("UpdateReceiver", c.UpdateRuleUnion)
	c.Mapping("DeleteReceiver", c.DeleteRuleUnion)
}

// @router /:ruleunionid [put]
func (c *RuleUnionController) UpdateRuleUnion() {
	var ruleunion models.RuleUnions
	var ans common.Res
	receiverId := c.Ctx.Input.Param(":ruleunionid")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ruleunion)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		if ruleunion.Expression != "" {
			root, err := common.BuildTree(ruleunion.Expression)
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			} else {
				ReversePolishNotation := common.Converse2ReversePolishNotation(root)
				ruleunion.ReversePolishNotation = ReversePolishNotation
				id, _ := strconv.ParseInt(receiverId, 10, 64)
				ruleunion.Id = id
				err = ruleunion.UpdateRuleUnion()
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(receiverId, 10, 64)
			ruleunion.Id = id
			err = ruleunion.UpdateRuleUnion()
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			}
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, ruleunion)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:ruleunionid [delete]
func (c *RuleUnionController) DeleteRuleUnion() {
	ruleunionId := c.Ctx.Input.Param(":ruleunionid")
	var Ruleunion *models.RuleUnions
	var ans common.Res
	err := Ruleunion.DeleteRuleUnion(ruleunionId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, ruleunionId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
