package main

import (
	"./cloudstack_library_code_generator"
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
	return cloudstack_library_code_generator.ApiToObjectNameMap[s]
}

func ParamType(s string) string {
	if s == "string" || s == "tzdata" {
		return "*String"
	} else if s == "boolean" {
		return "*Boolean"
	} else if s == "short" || s == "long" || s == "integer" {
		return "*Integer"
	} else if s == "uuid" {
		return "*ID"
	} else {
		return "*String"
	}
}

func RespType(s string) string {
	if s == "boolean" {
		return "*Boolean"
	} else if s == "short" || s == "long" || s == "integer" {
		return "*Number"
	} else if s == "responseobject" {
		return "json.RawMessage"
	} else {
		return "*String"
	}
}

func IsList(s string) bool    { return s == "list" || s == "set" }
func IsBoolean(s string) bool { return s == "boolean" }
func IsUUID(s string) bool    { return s == "uuid" }
func IsInteger(s string) bool { return s == "short" || s == "long" || s == "integer" }
func IsString(s string) bool  { return s == "string" || s == "tzdata" }

func IsListAPI(s string) bool             { return strings.HasPrefix(strings.ToLower(s), "list") }
func IsQueryAsyncJobResult(s string) bool { return strings.ToLower(s) == "queryasyncjobresult" }

func IsId(s string) bool  { return strings.HasSuffix(strings.ToLower(s), "id") }
func IsIds(s string) bool { return strings.HasSuffix(strings.ToLower(s), "ids") }

func main() {

	listApisJson := os.Args[1]
	input, err := ioutil.ReadFile(listApisJson)
	if err != nil {
		panic(err)
	}

	templateFile := os.Args[2]
	funcMap := template.FuncMap{
		"title":                 strings.Title,
		"toLower":               strings.ToLower,
		"objectName":            GetObjectName,
		"respType":              RespType,
		"paramType":             ParamType,
		"isList":                IsList,
		"isBoolean":             IsBoolean,
		"isUUID":                IsUUID,
		"isInteger":             IsInteger,
		"isString":              IsString,
		"isListAPI":             IsListAPI,
		"isQueryAsyncJobResult": IsQueryAsyncJobResult,
		"isId":                  IsId,
		"isIds":                 IsIds,
	}
	tmpl := template.Must(
		template.New(templateFile).Funcs(funcMap).ParseFiles(templateFile))

	directory := os.Args[3]

	var resp ListApisResponse
	if err := json.Unmarshal(input, &resp); err != nil {
		panic(err)
	}

	for _, api := range resp.ListApisResponse.Api {

		if api.Name == "queryAsyncJobResult" {
			api.Response = append(api.Response, Response{Name: "jobid", Type: "string"})
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
}
