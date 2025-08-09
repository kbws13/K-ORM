package session

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // 假设使用 sqlite3
	"korm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	// 需要先创建数据库连接和 dialect
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	dial, _ := dialect.GetDialect("sqlite3")
	s := New(db, dial).Model(&User{}) // New 函数现在需要 dialect 参数

	_ = s.DropTable()
	_ = s.CreateTable()

	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
