package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"./vars"
)

type Param struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Length      float64 `json:"length"`
	Required    bool    `json:"required"`
	Type        string  `json:"type"`
	Since       string  `json:"since"`
}

type Response struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Response    []Response `json:"response"`
}

type ListApisResponse struct {
	ListApisResponse struct {
		Count float64 `json:"count"`
		Api   []struct {
			Name        string     `json:"name"`
			Description string     `json:"description"`
			IsAsync     bool       `json:"isasync"`
			Related     string     `json:"related"`
			Since       string     `json:"since"`
			Params      []Param    `json:"params"`
			Response    []Response `json:"response"`
		} `json:"api"`
	} `json:"listapisresponse"`
}

type ParamList []Param

func (l ParamList) Len() int           { return len(l) }
func (l ParamList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ParamList) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l ParamList) Sort()              { sort.Sort(l) }

type ResponseList []Response

func (l ResponseList) Len() int           { return len(l) }
func (l ResponseList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ResponseList) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l ResponseList) Sort() {
	for _, x := range []Response(l) {
		ResponseList(x.Response).Sort()
	}
	sort.Sort(l)
}

func GetObjectName(s string) string {
	return vars.ApiToObjectNameMap[s]
}

func ParamType(s string) string {
	if s == "string" || s == "tzdata" {
		return "NullString"
	} else if s == "boolean" {
		return "NullBool"
	} else if s == "short" || s == "long" || s == "integer" {
		return "NullInt64"
	} else if s == "uuid" {
		return "ID"
	} else if s == "list" {
		return "[]string"
	} else if s == "map" {
		return "map[string]string"
	} else {
		return "NullString"
	}
}

func RespType(s string) string {
	if s == "boolean" {
		return "NullBool"
	} else if s == "short" || s == "long" || s == "integer" {
		return "NullInt64"
	} else if s == "responseobject" {
		return "json.RawMessage"
	} else {
		return "NullString"
	}
}

func IsBoolean(s string) bool { return s == "boolean" }
func IsUUID(s string) bool    { return s == "uuid" }
func IsInteger(s string) bool { return s == "short" || s == "long" || s == "integer" }
func IsString(s string) bool  { return s == "string" || s == "tzdata" }
func IsMap(s string) bool     { return s == "map" }
func IsList(s string) bool    { return s == "list" || s == "set" }
func IsResult(s string) bool  { return strings.ToLower(s) == "result" }

func IsListAPI(s string) bool             { return strings.HasPrefix(strings.ToLower(s), "list") }
func IsQueryAsyncJobResult(s string) bool { return strings.ToLower(s) == "queryasyncjobresult" }
func IsTags(s string) bool                { return strings.ToLower(s) == "tags" }
func IsService(s string) bool             { return strings.ToLower(s) == "service" }
func IsEgressRule(s string) bool          { return strings.ToLower(s) == "egressrule" }
func IsIngressRule(s string) bool         { return strings.ToLower(s) == "ingressrule" }
func IsIpAddressCommand(s string) bool {
	return strings.ToLower(s) == "associateipaddress" ||
		strings.ToLower(s) == "updateipaddress"
}
func IsSecurityGroupCommand(s string) bool {
	return strings.ToLower(s) == "authorizesecuritygroupegress" ||
		strings.ToLower(s) == "authorizesecuritygroupingress"
}

func IsId(s string) bool  { return strings.HasSuffix(strings.ToLower(s), "id") }
func IsIds(s string) bool { return strings.HasSuffix(strings.ToLower(s), "ids") }

func ImportStrings(params []Param) bool {
	b := false
	for _, p := range params {
		if p.Type == "list" {
			b = true
			break
		}
	}
	return b
}

func main() {

	listApisJson := os.Args[1]
	input, err := ioutil.ReadFile(listApisJson)
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"title":                  strings.Title,
		"toLower":                strings.ToLower,
		"objectName":             GetObjectName,
		"respType":               RespType,
		"paramType":              ParamType,
		"isList":                 IsList,
		"isBoolean":              IsBoolean,
		"isUUID":                 IsUUID,
		"isMap":                  IsMap,
		"isInteger":              IsInteger,
		"isString":               IsString,
		"isResult":               IsResult,
		"isListAPI":              IsListAPI,
		"isQueryAsyncJobResult":  IsQueryAsyncJobResult,
		"isIpAddressCommand":     IsIpAddressCommand,
		"isSecurityGroupCommand": IsSecurityGroupCommand,
		"isTags":                 IsTags,
		"isService":              IsService,
		"isEgressRule":           IsEgressRule,
		"isIngressRule":          IsIngressRule,
		"isId":                   IsId,
		"isIds":                  IsIds,
		"importStrings":          ImportStrings,
	}

	var tmpl *template.Template
	var directory string
	var resp ListApisResponse
	if err := json.Unmarshal(input, &resp); err != nil {
		panic(err)
	}

	directory = "apis"
	tmpl = template.Must(
		template.New("api.go.tmpl").Funcs(funcMap).ParseFiles("api.go.tmpl"))

	for _, api := range resp.ListApisResponse.Api {

		if api.Name == "queryAsyncJobResult" {
			continue
		}
		fmt.Println(api.Name)
		ParamList(api.Params).Sort()
		ResponseList(api.Response).Sort()

		var b bytes.Buffer
		if err = tmpl.Execute(&b, api); err != nil {
			panic(err)
		}

		re := regexp.MustCompile(`(?m)^\s+$`)
		s := strings.Replace(re.ReplaceAllString(b.String(), ""), "\n\n", "\n", -1)
		source, err := format.Source([]byte(s))
		if err != nil {
			log.Println(err)
			ioutil.WriteFile(
				directory+"/"+strings.ToLower(api.Name)+".go", []byte(s), 0644)
		} else {
			ioutil.WriteFile(
				directory+"/"+strings.ToLower(api.Name)+".go", source, 0644)
		}
	}

	directory = "structs"
	tmpl = template.Must(
		template.New("struct.go.tmpl").Funcs(funcMap).ParseFiles("struct.go.tmpl"))

	for _, api := range resp.ListApisResponse.Api {

		fmt.Println(api.Name)
		ResponseList(api.Response).Sort()

		var b bytes.Buffer
		if err = tmpl.Execute(&b, api); err != nil {
			panic(err)
		}

		re := regexp.MustCompile(`(?m)^\s+$`)
		s := strings.Replace(re.ReplaceAllString(b.String(), ""), "\n\n", "\n", -1)
		source, err := format.Source([]byte(s))
		if err != nil {
			log.Println(err)
			ioutil.WriteFile(
				directory+"/"+strings.ToLower(GetObjectName(api.Name))+"-"+strings.ToLower(api.Name)+".go", []byte(s), 0644)
		} else {
			ioutil.WriteFile(
				directory+"/"+strings.ToLower(GetObjectName(api.Name))+"-"+strings.ToLower(api.Name)+".go", source, 0644)
		}
	}

	f, err := os.Create("isasync.txt")
	for _, api := range resp.ListApisResponse.Api {
		f.WriteString(fmt.Sprintf("\"%v\": %v,\n", api.Name, api.IsAsync))
	}
}
