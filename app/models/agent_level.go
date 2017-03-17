package models

import "github.com/revel/revel"

type Agentlevel struct {
	Id         int    `xorm:"pk"`
	Level_name string `xorm:"varchar(11)"`
	Level_info string `xorm:"text"`
	Divided    string `xorm:"DECIMAL(11,0)"`
}

func (al *Agentlevel) Save(agentlevel Agentlevel) bool {
	has, err := DB_Write.Table("dl_agent_level").Insert(agentlevel)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (al *Agentlevel) Delete(id int) bool {
	agentlevel := new(Agentlevel)
	has, err := DB_Write.Table("dl_agent_level").Id(id).Delete(agentlevel)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (al *Agentlevel) Update(id int, agentlevel Agentlevel) bool {
	has, err := DB_Write.Table("dl_agent_level").Id(id).Update(agentlevel)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (al *Agentlevel) Getall() []Agentlevel {
	var agentlevels []Agentlevel
	err := DB_Write.Table("dl_agent_level").Find(&agentlevels)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return agentlevels
}
