package main

import (
	"flag"
	"gozero_demo/cmd/gorm_gen/common"
	"gozero_demo/internal/config"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasSuffix(ct, "unsigned") {
			return "uint8"
		}
		return "int8"
	},
	"smallint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasSuffix(ct, "unsigned") {
			return "uint16"
		}
		return "int16"
	},
	"mediumint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasSuffix(ct, "unsigned") {
			return "uint32"
		}
		return "int32"
	},
	"int": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasSuffix(ct, "unsigned") {
			return "uint32"
		}
		return "int32"
	},
	"bigint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasSuffix(ct, "unsigned") {
			return "uint64"
		}
		return "int64"
	},
	"json": func(columnType gorm.ColumnType) (dataType string) {
		return "json.RawMessage"
	},
}

var configFile = flag.String("f", "etc/gozero_demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	g := gen.NewGenerator(gen.Config{
		OutPath:      c.Mysql.OutPath,      // 查询类代码的输出路径
		OutFile:      "gen.go",             // 代码输出文件名，默认: gen.go
		ModelPkgPath: c.Mysql.ModelPkgPath, // 生成的 model 包名

		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式

		// 如果你希望为可为null的字段生成属性为指针类型, 设置 FieldNullable 为 true
		FieldNullable: true,
		// 如果你希望在 `Create` API 中为字段分配默认值, 设置 FieldCoverable 为 true, 参考: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: false,
		// 如果你希望生成无符号整数类型字段, 设置 FieldSignable 为 true
		FieldSignable: true,
		// 如果你希望从数据库生成索引标记, 设置 FieldWithIndexTag 为 true
		FieldWithIndexTag: true,
		// 如果你希望从数据库生成类型标记, 设置 FieldWithTypeTag 为 true
		FieldWithTypeTag: true,
		// 如果你需要对查询代码进行单元测试, 设置 WithUnitTest 为 true
		WithUnitTest: false,
	})

	gormDB, err := gorm.Open(mysql.Open(c.Mysql.DataSource))
	if err != nil {
		panic(err)
	}
	// 连接数据库
	g.UseDB(gormDB) // reuse your gorm db

	// 数据类型映射
	g.WithDataTypeMap(dataMap)

	// 生成模型
	g.ApplyBasic(
		g.GenerateAllTable(
			// gen.FieldNew("gorm.Model", "gorm.Model", gen.FieldTag),

			// gen.FieldIgnore("id", "created_at", "updated_at", "deleted_at"),

			gen.WithMethod(common.Method{}),
		)...,
	)

	g.Execute()
}
