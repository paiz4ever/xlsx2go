package cfgs

import (
	"log"
	"strconv"
	"github.com/xuri/excelize/v2"
)

func init() {
	f, err := excelize.OpenFile("xlsx/test.xlsx")
	if err != nil {
		log.Println("xlsx/test.xlsx异常: ", err)
		return
	}
	defer f.Close()
	
	initUser(f)
	initCompany(f)
	
}



var testUserCfgs = make(map[string]*User)

func initUser(f *excelize.File) {
	rows, err := f.GetRows("user")
	if err != nil {
		log.Println("xlsx/test.xlsx获取sheet<user>异常: ", err)
		return
	}
	for _, row := range rows[2:] {
		data := &User{}
		
		data.Name = row[0]
		
		
		v1, err := strconv.Atoi(row[1])
		if err != nil {
			log.Println("类型转换错误: ", err)
			return
		}
		data.Age = v1
		
		v2, err := strconv.ParseBool(row[2])
		if err != nil {
			log.Println("类型转换错误: ", err)
			return
		}
		data.IsAlive = v2
		testUserCfgs[data.Name] = data
	}
}

type User struct {
	Name string
	Age int
	IsAlive bool
	
}

func (c *User) GetData(keys ...string) []*User {
	datas := make([]*User, 0)
	for _, key := range keys {
		datas = append(datas, testUserCfgs[key])
	}
	return datas
}


var testCompanyCfgs = make(map[string]*Company)

func initCompany(f *excelize.File) {
	rows, err := f.GetRows("company")
	if err != nil {
		log.Println("xlsx/test.xlsx获取sheet<company>异常: ", err)
		return
	}
	for _, row := range rows[2:] {
		data := &Company{}
		
		data.Name = row[0]
		
		data.Boss = row[1]
		
		data.Address = row[2]
		
		testCompanyCfgs[data.Name] = data
	}
}

type Company struct {
	Name string
	Boss string
	Address string
	
}

func (c *Company) GetData(keys ...string) []*Company {
	datas := make([]*Company, 0)
	for _, key := range keys {
		datas = append(datas, testCompanyCfgs[key])
	}
	return datas
}

	