package models

import "github.com/revel/revel"

type Agent struct {
	Id           int    `xorm:"pk"`
	Name         string `xorm:"varchar(11)"`
	Cellphone    int    `xorm:"int(11)"`
	City         int    `xorm:"int(11)"`
	Address      string `xorm:"varchar(11)"`
	Bankno       int    `xorm:"int(25)"`
	Bank         string `xorm:"varchar(11)"`
	Signing_date int    `xorm:"int(11)"`
	Expire_date  int    `xorm:"int(11)"`
	Level        int    `xorm:"int(11)"`
}

func (a *Agent) Save(agent Agent) bool {
	has, err := DB_Write.Table("dl_agent").Insert(agent)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (a *Agent) Delete(id int) bool {
	agent := new(Agent)
	has, err := DB_Write.Table("dl_agent").Id(id).Delete(agent)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (a *Agent) Update(id int, agent Agent) bool {
	has, err := DB_Write.Table("dl_agent").Id(id).Update(agent)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (a *Agent) Getall() []Agent {
	var agents []Agent
	err := DB_Write.Table("dl_agent").Find(&agents)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return agents
}
