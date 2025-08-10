# K-ORM

K-ORM is a simple and lightweight Object-Relational Mapping (ORM) framework for Go, designed to provide basic database operations with an easy-to-use API.

## Features

- Simple ORM operations (Create, Read, Update, Delete)
- Support for raw SQL queries
- Database migration capabilities
- Transaction support
- Hooks for database operations (Before/After Insert, Update, Delete, Query)
- Multiple database dialect support (currently SQLite3, easily extensible)

## Installation

```bash
go get -u github.com/kbws13/K-ORM
```


## Quick Start

### 1. Initialize the Engine

```go
import "korm"

engine, err := korm.NewEngine("sqlite3", "korm.db")
if err != nil {
    log.Fatal(err)
}
defer engine.Close()
```


### 2. Create a Session

```go
s := engine.NewSession()
```


### 3. Define a Model

```go
type User struct {
    Name string `korm:"PRIMARY KEY"`
    Age  int
}
```


### 4. Perform Database Operations

```go
// Create table
s.Model(&User{}).CreateTable()

// Insert records
user1 := &User{Name: "Tom", Age: 18}
s.Insert(user1)

// Query records
var users []User
s.Find(&users)

// Update records
s.Where("Name = ?", "Tom").Update("Age", 20)

// Delete records
s.Where("Name = ?", "Tom").Delete()
```


## API Reference

### Engine

- `NewEngine(driver, source string) (*Engine, error)` - Create a new database engine
- `Close()` - Close the database connection
- `NewSession() *Session` - Create a new session
- `Transaction(f TxFunc) (interface{}, error)` - Execute operations in a transaction
- `Migrate(value interface{}) error` - Migrate database schema

### Session

- `Model(value interface{}) *Session` - Set the model for the session
- `CreateTable() error` - Create table for the model
- `DropTable() error` - Drop the table for the model
- `HasTable() bool` - Check if the table exists
- `Insert(values ...interface{}) (int64, error)` - Insert records
- `Find(values interface{}) error` - Find multiple records
- `First(value interface{}) error` - Find the first record
- `Update(kv ...interface{}) (int64, error)` - Update records
- `Delete() (int64, error)` - Delete records
- `Count() (int64, error)` - Count records
- `Limit(num int) *Session` - Add limit condition
- `Where(desc string, args ...interface{}) *Session` - Add where condition
- `OrderBy(desc string) *Session` - Add order by condition
- `Raw(sql string, values ...interface{}) *Session` - Execute raw SQL
- `Exec() (sql.Result, error)` - Execute the SQL statement
- `QueryRow() *sql.Row` - Query a single row
- `QueryRows() (*sql.Rows, error)` - Query multiple rows

## Hooks

K-ORM supports the following hooks that can be implemented on your models:

- `BeforeInsert(s *session.Session)`
- `AfterInsert(s *session.Session)`
- `BeforeUpdate(s *session.Session)`
- `AfterUpdate(s *session.Session)`
- `BeforeDelete(s *session.Session)`
- `AfterDelete(s *session.Session)`
- `BeforeQuery(s *session.Session)`
- `AfterQuery(s *session.Session)`

Example:
```go
type User struct {
    Name string `korm:"PRIMARY KEY"`
    Age  int
}

func (u *User) BeforeInsert(s *session.Session) {
    log.Println("Before inserting user")
}

func (u *User) AfterQuery(s *session.Session) {
    log.Println("After querying user")
}
```


## Transaction Support

```go
result, err := engine.Transaction(func(s *session.Session) (interface{}, error) {
    // Perform multiple operations
    s.Insert(&User{Name: "Tom", Age: 18})
    s.Insert(&User{Name: "Sam", Age: 25})
    
    // Return result or error
    return "success", nil
})
```

Contributions are welcome! Please feel free to submit a Pull Request.