package modules

import (
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// M is map
type M map[string]interface{}

// S is slice
type S []interface{}

// Prom ...
type Prom struct {
	ID  int64
	URL string
}

// Rule ...
type Rule struct {
	ID          int64             `json:"id"`
	AlertN      string 			  `json:"alertn"`
	Cluster     string 			  `json:"cluster"`
	Severity    string 			  `json:"severity"`
	Type        string 			  `json:"type"`
	ProjectName string 			  `json:"project_name"`
	AppName     string 			  `json:"app_name"`
	Env         string 			  `json:"env"`
	PromID      int64             `json:"prom_id"`
	Expr        string            `json:"expr"`
	Op          string            `json:"op"`
	Value       string            `json:"value"`
	For         string            `json:"for"`
	Labels      map[string]string `json:"labels"`
	//Labels      Labels
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
}


// Rules ...
type Rules []Rule

// PromRules ...
type PromRules struct {
	Prom  Prom
	Rules Rules
}

// RulesResp ...
type RulesResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Rules  `json:"data"`
}

// PromsResp ...
type PromsResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Prom `json:"data"`
}

// Content get prom rules
func (r Rules) Content() ([]byte, error) {
	rules := S{}
	for _, i := range r {
		rules = append(rules, M{
			"alert":  strconv.FormatInt(i.ID, 10),
			"expr":   strings.Join([]string{i.Expr, i.Op, i.Value}, " "),
			"for":    i.For,
			//"labels": i.Labels,
			"labels": M{
				"alertname":   i.Description,
				"summary":     i.Summary,
				"desp": 	   i.Description,
				"projectName": i.ProjectName,
				"appname": 	   i.AppName,
				"env": 	       i.Env,
				"altern":	   i.AlertN,
				"cluster":	   i.Cluster,
				"type":		   i.Type,
				"severity":    i.Severity,
			},
			"annotations": M{
				"rule_id":     strconv.FormatInt(i.ID, 10),
				"prom_id":     strconv.FormatInt(i.PromID, 10),
				"summary":     i.Summary,
				"description": i.Description,
				"projectName": i.ProjectName,
				"appname": 	   i.AppName,
				"env": 	       i.Env,
				"result":	   "{{ $value }}",
			},
		})
	}
	logs.Info("\"package-Rule---rules---\n %s ", rules) // for debug
	result := M{
		"groups": S{
			M{
				"name":  "ruleengine",
				"rules": rules,
			},
		},
	}

	//logs.Info("package-Rule---result---\n %s ", result) // for debug

	return yaml.Marshal(result)
}

// PromRules cut prom rules
func (r Rules) PromRules() []PromRules {
	tmp := map[int64]Rules{}

	for _, rule := range r {
		if v, ok := tmp[rule.PromID]; ok {
			tmp[rule.PromID] = append(v, rule)
		} else {
			tmp[rule.PromID] = Rules{rule}
		}
	}

	data := []PromRules{}
	for id, rules := range tmp {
		data = append(data, PromRules{
			Prom:  Prom{ID: id},
			Rules: rules,
		})
	}

	return data
}
