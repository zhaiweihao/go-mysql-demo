package model

import(
    "errors"
    _"github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var (
    CustomerNotExistError=errors.New("user not exists")
)

var db *xorm.Engine

func InitDB(dialect,conn string) error {
    var err error
    db,err=xorm.NewEngine(dialect,conn)
    if err !=nil{
        return err
    }
    db.ShowSQL(true);
    return db.Sync2(new(User))
}