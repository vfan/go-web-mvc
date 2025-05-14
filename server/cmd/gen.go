package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"mvc-demo/config"
)

func main() {
	// 连接数据库，读取.env和.env.local
	env := godotenv.Load()
	if env != nil {
		panic(fmt.Errorf("读取.env和.env.local失败: %w", env))
	}

	appConfig := config.GetConfig()
	dbConfig := appConfig.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("连接数据库失败: %w", err))
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		// 相对路径，会自动创建目录
		OutPath: "./dao/query",
		// 模式: QueryMode 只生成查询代码(不需要指定表)
		//      ModelMode 只生成模型代码(需要指定表)
		//      DefaultMode 同时生成模型和查询代码(需要指定表)
		Mode: gen.WithDefaultQuery | gen.WithoutContext,
		// 表字段可为null时，对应字段使用指针类型
		FieldNullable: true,
	})

	// 设置数据库连接
	g.UseDB(db)

	// 从数据库表生成所有模型
	// g.GenerateAllTable()

	// 或者只生成指定表的模型
	// 第一个参数是表名，第二个参数是模型名称
	g.GenerateModel("users")
	g.GenerateModel("universities")
	g.GenerateModel("students")

	// 生成代码
	g.Execute()
}
