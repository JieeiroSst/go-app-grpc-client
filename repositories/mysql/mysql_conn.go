package mysql

import (
	"sync"

	"github.com/JIeeiroSst/go-app/domain"
	"github.com/JIeeiroSst/go-app/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mutex    sync.Mutex
	instance *MysqlConn
)

type MysqlConn struct {
	db *gorm.DB
}

type Config struct {
	DSN string
}

func GetMysqlConnInstance(c *Config) *MysqlConn {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := c.DSN
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.InitZapLog().Sugar().Errorf("running server mysql serror %s", err)
			}
			log.InitZapLog().Info("server running get mysql success")
			instance = &MysqlConn{db: db}
			err = db.AutoMigrate(&domain.User{},&domain.Profile{})
			if err != nil {
				log.InitZapLog().Error("error connect db")
			}
			log.InitZapLog().Info("conecting db success")
		}
	}
	return instance
}

func NewMysqlConn(c *Config) *MysqlConn {
	return &MysqlConn{
		db: GetMysqlConnInstance(c).db,
	}
}

func (mysql *MysqlConn) UserAll() (users []domain.User, err error){
	err=mysql.db.Find(&users).Error
	if err!=nil{
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return nil,err
	}
	log.InitZapLog().Info("server query sucess")
	return
}

func (mysql *MysqlConn) UserById(id int) (user domain.User,err error) {
	err = mysql.db.Where("id = ?",id).Find(&user).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return domain.User{},err
	}
	log.InitZapLog().Info("server query sucess")
	return
}

func (mysql *MysqlConn) CreateUser(user domain.User) (err error){
	err=mysql.db.Create(&user).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn) UpdateUser(id int,user domain.User) (err error){
	err = mysql.db.Model(&domain.User{}).Where("id = ?",id).Updates(&user).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn) DeleteUser(id int) (err error) {
	err =mysql.db.Where("id = ?",id).Delete(&domain.User{}).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn) ProfileAll() (profiles []domain.Profile,err error){
	err = mysql.db.Find(&profiles).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return nil,err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn) ProfileById(id int) (profile domain.Profile,err error){
	err =mysql.db.Where("id = ?",id).Find(&profile).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return domain.Profile{},err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn)UpdateProfile(id int,profile domain.Profile) (err error){
	err = mysql.db.Model(&domain.User{}).Where("id = ?",id).Updates(&profile).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn) DeleteProfile(id int) (err error) {
	err =mysql.db.Where("id = ?",id).Delete(&domain.Profile{}).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}

func (mysql *MysqlConn)CreateProfile(profile domain.Profile) (err error){
	err=mysql.db.Create(&profile).Error
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("server query error %s",err)
		return err
	}
	log.InitZapLog().Info("server query sucess")
	return 
}