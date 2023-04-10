package sqlite

import (
	"log"
	"time"
)

type Authuser struct {
	ID          int    `xorm:"id int pk autoincr"`
	Name        string `xorm:"name"`
	Description string `xorm:"description"`
	Dr          int    `xorm:"dr"`
	Secretid    string `xorm:"secretid"`
	Qrcode      string `xorm:"qrcode"`
	CreatedTime string `xorm:"createdtime"`
}

func ListAuthUser() (aus []Authuser, err error) {
	err = authEngine.Table("authuser").Where("dr=?", 0).Find(&aus)
	if err != nil {
		log.Println("cannot list auth user , error: ", err)
		return nil, err
	}
	return aus, nil
}

func CreateAuthUser(au Authuser) error {
	au.CreatedTime = time.Now().Format("2006-01-02 15:30")

	_, err := authEngine.Insert(&au)
	if err != nil {
		log.Println("cannot create auth user, error: ", err)
		return err
	}
	return nil
}

func CountAuthUserByName(name string) (c int64, err error) {
	c, err = authEngine.Table("authuser").Where("authuser.name=?", name).Count()
	if err != nil {
		return 0, err
	}
	return c, nil
}

func GetQrcode(id int) (au Authuser, err error) {
	res, err := authEngine.Table("authuser").Where("id = ?", id).Get(&au)
	if err != nil || res == false {
		log.Println("cannot list auth user , error: ", err)
		return au, err
	}

	return au, nil
}

func SearchUsername(name string) (au Authuser, err error) {
	res, err := authEngine.Table("authuser").Where("name = ?", name).Get(&au)
	if err != nil || res == false {
		log.Println("cannot search auth user , error: ", err)
	}
	return au, nil
}

func DeleteAuthUser(id int) (err error) {
	var au Authuser
	_, err = authEngine.Table("authuser").ID(id).Delete(&au)
	if err != nil {
		log.Println("cannot delete auth user , error: ", err)
		return err
	}
	return nil
}
