package model

import(
    "errors"
)

type User struct{
    Id int64
    Name string 
}
func (User) TableName() string {
    return "user"
}
func (c *User) Create() error{
    if affected,err :=db.InsertOne(c);err !=nil{
        return err
    }else if affected==0{
        return errors.New("Affeced rows:0")
    }
    return nil
}
func (User) Retrive() (*User,error) {
    var c User
    _,err :=db.Get(&c)
    if err !=nil{
        return nil,err
    }
    return &c,nil
}
func (User) Get(id int64) (*User,error){
    var c User
    _,err :=db.Id(id).Get(&c)
    if(err !=nil){
        return nil,err
    }
    return &c,nil
}