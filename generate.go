package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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

type Api struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsAsync     bool       `json:"isasync"`
	Related     string     `json:"related"`
	Since       string     `json:"since"`
	Params      []Param    `json:"params"`
	Response    []Response `json:"response"`
}

type ListApisResponse struct {
	ListApisResponse struct {
		Count float64 `json:"count"`
		Api   []Api   `json:"api"`
	} `json:"listapisresponse"`
}

type ApiList []Api

func (l ApiList) Len() int           { return len(l) }
func (l ApiList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ApiList) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l ApiList) Sort() {
	sort.Sort(l)
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

func Unique(xs interface{}, fn func(interface{}, interface{}) bool) interface{} {
	vs := reflect.ValueOf(xs)
	ys := reflect.MakeSlice(vs.Type(), 0, vs.Len())

	for i := 0; i < vs.Len(); i++ {
		if i == 0 {
			ys = reflect.Append(ys, vs.Index(i))
		} else {
			// fmt.Println(vs.Index(i-1).Interface().(Param).Name, vs.Index(i).Interface().(Param).Name)
			if !fn(vs.Index(i-1).Interface(), vs.Index(i).Interface()) {
				ys = reflect.Append(ys, vs.Index(i))
			} else {
				// fmt.Println("REDUNDANNT", vs.Index(i-1).Interface())
			}
		}
	}
	return ys.Interface()
}

func GetObjectName(s string) string {
	return vars.ApiToObjectNameMap[s]
}

func GetCamelCase(s string) string {
	camel, ok := vars.CamelMap[strings.ToLower(s)]
	if s != "" && !ok {
		log.Fatalf("%s is not found in CamelMap", strings.ToLower(s))
	}
	return camel
}

func GetFileFromApi(s string) string {
	for file := range vars.FileApiMap {
		for _, api := range vars.FileApiMap[file] {
			if s == api {
				return file
			}
		}
	}
	return ""
}

func ModifyReservedName(s string) string {
	switch s {
	case "type":
		return "typ"
	default:
		return s
	}
}

func ParamType(s string) string {
	if s == "string" || s == "tzdata" {
		return "string"
	} else if s == "boolean" {
		return "bool"
	} else if s == "short" || s == "long" || s == "integer" {
		return "int"
	} else if s == "list" {
		return "[]string"
	} else if s == "map" {
		return "map[string]string"
	} else {
		return "string"
	}
}

func ParamNullType(s string) string {
	if s == "string" || s == "tzdata" {
		return "NullString"
	} else if s == "boolean" {
		return "NullBool"
	} else if s == "short" || s == "long" || s == "integer" {
		return "NullNumber"
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
		return "NullNumber"
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

func Comment(s string) string {
	maxlen := 80
	words := strings.Split(s, " ")
	lines := []string{}

	// initialize line
	line := []string{"//"}
	lineLen := 2

	if strings.TrimSpace(s) == "" {
		return ""
	}

	for _, word := range words {
		if lineLen+1+len(word) > maxlen {
			lines = append(lines, strings.Join(line, " "))
			line = []string{"//"}
			lineLen = 2
		}
		line = append(line, word)
		lineLen = lineLen + 1 + len(word)
	}
	if lineLen > 2 {
		lines = append(lines, strings.Join(line, " "))
	}
	// fmt.Println(s)
	// fmt.Print(strings.Join(lines, "\n"))
	return strings.Join(lines, "\n")
}

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
		"camel":                  GetCamelCase,
		"modReservedName":        ModifyReservedName,
		"comment":                Comment,
		"respType":               RespType,
		"paramType":              ParamType,
		"paramNullType":          ParamNullType,
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
	var resp ListApisResponse
	if err := json.Unmarshal(input, &resp); err != nil {
		panic(err)
	}

	tmpl = template.Must(template.New(vars.ApiTemplateFile).Funcs(funcMap).ParseFiles(vars.ApiTemplateFile))

	if err := os.MkdirAll(vars.ApiOutputDir, 0755); err != nil {
		log.Fatal(err)
	}

	for _, api := range resp.ListApisResponse.Api {

		if api.Name == "queryAsyncJobResult" {
			continue
		}
		fmt.Println(api.Name)
		ParamList(api.Params).Sort()
		api.Params = Unique(api.Params,
			func(x interface{}, y interface{}) bool {
				return x.(Param).Name == y.(Param).Name
			}).([]Param)
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
			source = []byte(s)
		}

		fname := GetFileFromApi(api.Name)
		if fname == "" {
			log.Printf("Not Found file for API: %s\n", api.Name)
			fname = strings.ToLower(api.Name)
		}
		fpath := vars.ApiOutputDir + "/" + fname + "Api.go"

		var f *os.File
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			f, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			f.WriteString("package cloudstack\n\n")
		} else {
			f, err = os.OpenFile(fpath, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}
		f.Write(source)
		f.WriteString("\n")
		f.Close()
	}

	tmpl = template.Must(template.New(vars.StructTemplateFile).Funcs(funcMap).ParseFiles(vars.StructTemplateFile))

	if err := os.MkdirAll(vars.StructOutputDir, 0755); err != nil {
		log.Fatal(err)
	}

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
			source = []byte(s)
		}

		objName := GetCamelCase(strings.ToLower(GetObjectName(api.Name)))
		fname := vars.StructOutputDir + "/" + objName + "Struct-" + strings.ToLower(api.Name) + ".go"
		ioutil.WriteFile(fname, source, 0644)
	}

	isAsyncFile, err := os.Create("isasync.txt")
	for _, api := range resp.ListApisResponse.Api {
		isAsyncFile.WriteString(fmt.Sprintf("\"%v\": %v,\n", api.Name, api.IsAsync))
	}
	isAsyncFile.Close()

	names := make([]string, 0)
	for _, api := range resp.ListApisResponse.Api {
		for i := range api.Params {
			names = append(names, strings.ToLower(api.Params[i].Name))
		}
		for i := range api.Response {
			names = append(names, strings.ToLower(api.Response[i].Name))
		}
	}
	sort.Strings(names)
	names = Unique(names,
		func(s, t interface{}) bool { return s.(string) == t.(string) }).([]string)

	namesFile, err := os.Create("names.txt")
	for i := range names {
		namesFile.WriteString(fmt.Sprintf("%s\n", names[i]))
	}
	namesFile.Close()

	apis := resp.ListApisResponse.Api
	ApiList(apis).Sort()
	commandsFile, err := os.Create("commands.txt")
	commandsFile.WriteString("func getCommand(name string) *Command {\n")
	commandsFile.WriteString("switch strings.ToLower(name) {\n")
	for _, api := range apis {
		commandsFile.WriteString(fmt.Sprintf(
			`case "%s": 
					     return &Command {
					         Name: "%s",
				             IsAsync: %v,
						     IsList: %v,
						     ObjectType : "%s",
				         }
				 `,
			strings.ToLower(api.Name), api.Name, api.IsAsync,
			IsListAPI(api.Name), GetObjectName(api.Name)))
	}
	commandsFile.WriteString(`
	    default:
			return nil
	}
    `)
	commandsFile.WriteString("}")
	commandsFile.Close()
}
