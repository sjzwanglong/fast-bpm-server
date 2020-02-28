package main

import (
	"bufio"
	"fast-bpm/utils"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Column struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG"`
	TableSchema            string `gorm:"column:TABLE_SCHEMA"`
	TabName                string `gorm:"column:TABLE_NAME"`
	ColumnName             string `gorm:"column:COLUMN_NAME"`
	OrdinalPosition        int    `gorm:"column:ORDINAL_POSITION"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT"`
	IsNullable             string `gorm:"column:IS_NULLABLE"`
	DataType               string `gorm:"column:DATA_TYPE"`
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   int    `gorm:"column:CHARACTER_OCTET_LENGTH"`
	NumericPrecision       int    `gorm:"column:NUMERIC_PRECISION"`
	NumericScale           int    `gorm:"column:NUMERIC_SCALE"`
	DatetimePrecision      int    `gorm:"column:DATETIME_PRECISION"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME"`
	CollationName          string `gorm:"column:COLLATION_NAME"`
	ColumnType             string `gorm:"column:COLUMN_TYPE"`
	ColumnKey              string `gorm:"column:COLUMN_KEY"`
	Extra                  string `gorm:"column:EXTRA"`
	Privileges             string `gorm:"column:PRIVILEGES"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT"`
	GenerationExpression   string `gorm:"column:GENERATION_EXPRESSION"`
	SrsId                  int    `gorm:"column:SRS_ID"`
}

func (c *Column) TableName() string {
	return "COLUMNS"
}

var wg sync.WaitGroup

// 数据库字段类型映射
var columnTypeMap = map[string]string{
	"varchar":   "string",
	"tinyint":   "int",
	"mediumint": "int",
	"datetime":  "time.Time",
}

// 系统字段，无需在具体model中声明
var specialColumn = "id,createdAt,updatedAt,deletedAt,createdBy,sysState"

// flag 参数表
var (
	tableName        string
	mode             string
	to               string
	database         string
	ip               string
	username         string
	password         string
	replacePrefix    string
	controllerFile   string
	controllerStruct string
	serviceFile      string
	serviceStruct    string
	daoFile          string
	daoStruct        string
	modelFile        string
	modelStruct      string
	routerFile       string
	routerStruct     string
	filePrefix       string
	structPrefix     string
)

func main() {
	// 基本
	flag.StringVar(&mode, "mode", "0", "Use -mode <0:all, 1:controller, 2:service, 3:dao, 4:model, 5:router 允许多选，以','分隔>")
	flag.StringVar(&tableName, "table", "", "Use -table <table name>")
	flag.StringVar(&to, "to", "dist", "Use -to <生成文件到指定目录. dist:当前dist目录, pro:生成到项目目录对应目录>")
	flag.StringVar(&ip, "ip", "127.0.0.1", "Use -ip <database ip address>")
	flag.StringVar(&database, "db", "information_schema", "Use -db <对应数据库>")
	flag.StringVar(&username, "u", "root", "Use -u <数据库用户名>")
	flag.StringVar(&password, "p", "root", "Use -p <数据库密码>")
	flag.StringVar(&replacePrefix, "prefix", "bpm", "Use -prefix <转化结构体和文件时会被剔除此前缀>")
	// 高级
	flag.StringVar(&controllerFile, "cf", "", "Use -cf <controller文件名>")
	flag.StringVar(&controllerStruct, "cs", "", "Use -cs <controller结构体名称>")
	flag.StringVar(&serviceFile, "sf", "", "Use -sf <service文件名>")
	flag.StringVar(&serviceStruct, "ss", "", "Use -ss <service结构体名称>")
	flag.StringVar(&daoFile, "df", "", "Use -df <dao文件名>")
	flag.StringVar(&daoStruct, "ds", "", "Use -ds <dao结构体名称>")
	flag.StringVar(&modelFile, "mf", "", "Use -mf <model文件名>")
	flag.StringVar(&modelStruct, "ms", "", "Use -ms <model结构体名称>")
	flag.StringVar(&routerFile, "rf", "", "Use -rf <router文件名>")
	flag.StringVar(&routerStruct, "rs", "", "Use -rs <router结构体名称>")
	flag.Parse()
	if len(tableName) <= 0 {
		fmt.Println("请使用 '-table' 添加表名！！！")
		return
	}

	filePrefix = getFilePrefix(strings.Replace(tableName, replacePrefix+"_", "", 1))
	structPrefix = getStructPrefix(filePrefix)

	fmt.Println("开始执行...")
	modes := strings.Split(mode, ",")
	for _, v := range modes {
		switch v {
		case "0":
			wg.Add(5)
			go genRouter()
			go genController()
			go genService()
			go genDao()
			go genModel()
			goto end
		case "1":
			wg.Add(1)
			go genController()
		case "2":
			wg.Add(1)
			go genService()
		case "3":
			wg.Add(1)
			go genDao()
		case "4":
			wg.Add(1)
			go genModel()
		case "5":
			wg.Add(1)
			go genRouter()
		}
	}
end:
	wg.Wait()
	fmt.Println("生成完毕...")
}

// 生成controller
func genController() {
	defer wg.Done()
	cf := filePrefix
	if len(controllerFile) > 0 {
		cf = controllerFile
	}
	cs := structPrefix
	if len(controllerStruct) > 0 {
		cs = controllerStruct
	}
	fmt.Println("开始生成controller...")
	tmplFile, err := os.Open("controllerTmpl.tmpl")
	if err != nil {
		panic(err)
	}
	defer tmplFile.Close()
	filePath := utils.GetAbsPath("./dist/" + cf + "Controller.go")
	if to == "pro" {
		filePath = utils.GetAbsPath("../controller/" + cf + "Controller.go")
	}
	targetFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	rd := bufio.NewReader(tmplFile)
	wd := bufio.NewWriter(targetFile)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "{controllerName}") {
			line = strings.ReplaceAll(line, "{controllerName}", cs)
		}
		if strings.Contains(line, "{servName}") {
			line = strings.ReplaceAll(line, "{servName}", cf)
		}
		if strings.Contains(line, "{servType}") {
			line = strings.ReplaceAll(line, "{servType}", cs)
		}
		wd.WriteString(line)
	}
	wd.Flush() //将缓存中的内容写入文件
	fmt.Println("controller生成完毕！")
}

// 生成service
func genService() {
	defer wg.Done()
	fmt.Println("开始生成service...")
	sf := filePrefix
	if len(serviceFile) > 0 {
		sf = serviceFile
	}
	ss := structPrefix
	if len(serviceStruct) > 0 {
		ss = serviceStruct
	}
	tmplFile, err := os.Open("serviceTmpl.tmpl")
	if err != nil {
		panic(err)
	}
	defer tmplFile.Close()
	filePath := utils.GetAbsPath("./dist/" + sf + "Serv.go")
	if to == "pro" {
		filePath = utils.GetAbsPath("../service/" + sf + "Serv.go")
	}
	targetFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	rd := bufio.NewReader(tmplFile)
	wd := bufio.NewWriter(targetFile)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "{servName}") {
			line = strings.ReplaceAll(line, "{servName}", ss)
		}
		if strings.Contains(line, "{daoName}") {
			line = strings.ReplaceAll(line, "{daoName}", sf)
		}
		if strings.Contains(line, "{daoType}") {
			line = strings.ReplaceAll(line, "{daoType}", ss)
		}
		wd.WriteString(line)
	}
	wd.Flush() //将缓存中的内容写入文件
	fmt.Println("service生成完毕！")
}

// 生成dao
func genDao() {
	defer wg.Done()
	fmt.Println("开始生成dao...")
	df := filePrefix
	if len(daoFile) > 0 {
		df = daoFile
	}
	ds := structPrefix
	if len(daoStruct) > 0 {
		ds = daoStruct
	}
	tmplFile, err := os.Open("daoTmpl.tmpl")
	if err != nil {
		panic(err)
	}
	defer tmplFile.Close()
	filePath := utils.GetAbsPath("./dist/" + df + "Dao.go")
	if to == "pro" {
		filePath = utils.GetAbsPath("../dao/" + df + "Dao.go")
	}
	targetFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	rd := bufio.NewReader(tmplFile)
	wd := bufio.NewWriter(targetFile)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "{daoName}") {
			line = strings.ReplaceAll(line, "{daoName}", ds)
		}
		wd.WriteString(line)
	}
	wd.Flush() //将缓存中的内容写入文件
	fmt.Println("dao生成完毕！")
}

// 生成model
func genModel() {
	defer wg.Done()
	fmt.Println("开始生成model...")
	mf := filePrefix
	if len(modelFile) > 0 {
		mf = modelFile
	}
	ms := structPrefix
	if len(modelStruct) > 0 {
		ms = modelStruct
	}
	tmplFile, err := os.Open("modelTmpl.tmpl")
	if err != nil {
		panic(err)
	}
	defer tmplFile.Close()
	filePath := utils.GetAbsPath("./dist/" + mf + "Model.go")
	if to == "pro" {
		filePath = utils.GetAbsPath("../model/" + mf + "Model.go")
	}
	targetFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	rd := bufio.NewReader(tmplFile)
	wd := bufio.NewWriter(targetFile)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "{modelName}") {
			line = strings.ReplaceAll(line, "{modelName}", ms)
		}
		if strings.Contains(line, "{tableName}") {
			line = strings.ReplaceAll(line, "{tableName}", tableName)
		}
		if strings.Contains(line, "{columns}") {
			line = strings.ReplaceAll(line, "{columns}", "")
			genColumns(wd)
		}
		wd.WriteString(line)
	}
	wd.Flush() //将缓存中的内容写入文件
	fmt.Println("model生成完毕！")
}

// 生成model结构体具体属性
func genColumns(wd *bufio.Writer) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ip, database)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		fmt.Println("数据库连接失败！ err=", err)
	}
	defer db.Close()
	columns := &[]*Column{}
	db.Where(map[string]interface{}{"TABLE_NAME": tableName}).Find(columns)
	for _, v := range *columns {
		var columnName = v.ColumnName
		if strings.Contains(specialColumn, columnName) {
			continue
		}
		s1 := firstCharUpper(columnName)
		column := fmt.Sprintf("%s %s `gorm:\"column:%s\" form:\"%s\" uri:\"%s\" json:\"%s\"`\n", s1, columnTypeMap[v.DataType], columnName, columnName, columnName, columnName)
		wd.WriteString(column)
	}
}

// 生成router
func genRouter() {
	defer wg.Done()
	fmt.Println("开始生成router...")
	rf := filePrefix
	if len(daoFile) > 0 {
		rf = daoFile
	}
	rs := structPrefix
	if len(daoStruct) > 0 {
		rs = daoStruct
	}
	router := getRouterName(rf)
	tmplFile, err := os.Open("routerTmpl.tmpl")
	if err != nil {
		panic(err)
	}
	defer tmplFile.Close()
	filePath := utils.GetAbsPath("./dist/" + rf + "Router.go")
	if to == "pro" {
		filePath = utils.GetAbsPath("../server/" + rf + "Router.go")
	}
	targetFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	rd := bufio.NewReader(tmplFile)
	wd := bufio.NewWriter(targetFile)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "{routerName}") {
			line = strings.ReplaceAll(line, "{routerName}", rs)
		}
		if strings.Contains(line, "{dualRouter}") {
			line = strings.ReplaceAll(line, "{dualRouter}", router+"s")
		}
		if strings.Contains(line, "{singularRouter}") {
			line = strings.ReplaceAll(line, "{singularRouter}", router)
		}
		wd.WriteString(line)
	}
	wd.Flush() //将缓存中的内容写入文件
	fmt.Println("router生成完毕！")
}

// 转化字符串 将'_'去掉，改为驼峰
func converStr(str string) string {
	return getFilePrefix(str)
}

// 将第一个字符大写
func firstCharUpper(str string) string {
	return getStructPrefix(str)
}

// 得到文件名前缀
func getFilePrefix(tableName string) string {
	temp := strings.Split(tableName, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

// 得到结构体前缀
func getStructPrefix(filePrefix string) string {
	chars := []rune(filePrefix)

	if unicode.IsLower(chars[0]) {
		chars[0] = chars[0] - 32
	}

	return string(chars)
}

// 得到实际路由的名字
func getRouterName(filePrefix string) string {
	chars := []rune(filePrefix)
	for i, v := range chars {
		if i != 0 && unicode.IsUpper(v) {
			chars[i] = chars[i] + 32
			return string(chars[i:])
		}
	}
	return filePrefix
}
