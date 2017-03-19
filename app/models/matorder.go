package models

<<<<<<< HEAD
import "github.com/revel/revel"
import "drinking_life/app/utils"
=======
import _ "github.com/revel/revel"
>>>>>>> 25137ba484fa05a5661709767b3e2c5095cb42c9

type Matorder struct {
	Id             int64  `xomr:"pk"`
	Cart_id        int64  `xorm:"int"`
	Pubdate        int64  `xorm:"int"`
	Total_amount   int64  `xorm:"int"`
	Conact_name    string `xorm:"varchar(11)"`
	Conact_phone   string `xorm:"VARCHAR(11)"`
	Courier_number string `xorm:"VARCHAR(80)"`
	IsProcessing   int    `xorm:"int"`
	Product_info   string `xorm:"TEXT"`
}

func (m *Matorder) Save(order Matorder) bool {
	has, err := DB_Write.Table("dl_material_order").Insert(order)
	if err != nil {
		revel.WARN.Printf(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

func (m *Matorder) Payment(orderId int64) bool {
	has, err := DB_Write.Table("dl_order").Exec("update dl_material_order set isPayment = ?, Paytime = ? , where id = ? ", utils.ALERD_PAYID, utils.GetTimeNow, orderId)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
