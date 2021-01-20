package tgormv1

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	if db, err = gorm.Open("mysql", "root:1234qwer@tcp(localhost:3306)/admin"); err != nil {
		panic(err.Error())
	}
	beforeQueryCallback := func(scope *gorm.Scope) {
		fmt.Printf("--->sql 1:[%+v]\r\n", scope.CombinedConditionSql())
		fmt.Printf("--->sql 2:[%+v]\r\n", scope.SQL)
		fmt.Printf("--->sql 3:[%+v]\r\n", scope.SQLVars)
		for _, f := range scope.Fields() {
			fmt.Printf("--->sql select f:[%+v]\r\n", f.DBName)
		}
		// scope.SkipLeft()
	}
	db.Callback().Query().Before("gorm:query").Register("gorm:before_query", beforeQueryCallback)
}

func Test_Query(t *testing.T) {
	var cuser struct {
		ID   uint64 `grom:"cloumn:id"`
		Name string `grom:"cloumn:name"`
		Age  int    `grom:"cloumn:age"`
	}

	// if err := db.Table("customers").Where("id=?", 15889337792).Find(&cuser).Error; err != nil {
	// if err := db.Table("customers").Find(&cuser, "id=?", 15889337792).Error; err != nil {
	if err := db.Table("users").Find(&cuser, "name=? and age=?", "ssy", 18).Error; err != nil {
		t.Fatal("--->err: ", err.Error())
	}
	t.Log("--->result:", cuser)
}

func Test_Count(t *testing.T) {
	count := 0
	if err := db.Table("customers").Count(&count).Error; err != nil {
		t.Fatal("--->err: ", err.Error())
	}
	t.Log("--->result:", count)
}
