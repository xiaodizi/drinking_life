// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


<<<<<<< HEAD
type tRole struct {}
var Role tRole
=======
<<<<<<< HEAD
type tAdvertisement struct {}
var Advertisement tAdvertisement


func (_ tAdvertisement) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Save", args).Url
}

func (_ tAdvertisement) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Update", args).Url
}

func (_ tAdvertisement) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Getall", args).Url
=======
type tAdClass struct {}
var AdClass tAdClass
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f


func (_ tRole) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Save", args).Url
}

func (_ tRole) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Update", args).Url
}

func (_ tRole) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Getall", args).Url
}


type tAdvertisement struct {}
var Advertisement tAdvertisement


func (_ tAdvertisement) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Save", args).Url
}

func (_ tAdvertisement) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Delete", args).Url
}

func (_ tAdvertisement) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Update", args).Url
}

func (_ tAdvertisement) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Getall", args).Url
}

func (_ tAdvertisement) Sort(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.Sort", args).Url
}

func (_ tAdvertisement) IsDisplay(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.IsDisplay", args).Url
}

func (_ tAdvertisement) IsnotDisplay(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Advertisement.IsnotDisplay", args).Url
}


type tAgent struct {}
var Agent tAgent


func (_ tAgent) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Save", args).Url
}

func (_ tAgent) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Delete", args).Url
}

func (_ tAgent) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Update", args).Url
}

func (_ tAgent) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Getall", args).Url
}


type tCart struct {}
var Cart tCart


func (_ tCart) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cart.Save", args).Url
}


type tCoupon struct {}
var Coupon tCoupon


func (_ tCoupon) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Save", args).Url
}

func (_ tCoupon) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Delete", args).Url
}

func (_ tCoupon) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Update", args).Url
}

func (_ tCoupon) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Getall", args).Url
}

func (_ tCoupon) Isuse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Isuse", args).Url
}

func (_ tCoupon) Isnotuse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Coupon.Isnotuse", args).Url
}


type tEquipment struct {}
var Equipment tEquipment


func (_ tEquipment) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Equipment.Save", args).Url
}

func (_ tEquipment) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Equipment.Delete", args).Url
}

func (_ tEquipment) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Equipment.Update", args).Url
}

func (_ tEquipment) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Equipment.Getall", args).Url
}


type tModular struct {}
var Modular tModular


func (_ tModular) Save(
		) string {
	args := make(map[string]string)
	
<<<<<<< HEAD
	return revel.MainRouter.Reverse("Modular.Save", args).Url
=======
	return revel.MainRouter.Reverse("Product.Getall", args).Url
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f
}

func (_ tModular) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Modular.Update", args).Url
}

<<<<<<< HEAD
func (_ tModular) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Modular.Getall", args).Url
=======
type tMember struct {}
var Member tMember


func (_ tMember) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Save", args).Url
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f
}

func (_ tMember) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Update", args).Url
}

<<<<<<< HEAD
type tAdClass struct {}
var AdClass tAdClass
=======
func (_ tMember) Hasbyname(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Hasbyname", args).Url
}
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f

func (_ tMember) Hascode(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Hascode", args).Url
}

<<<<<<< HEAD
func (_ tAdClass) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdClass.Save", args).Url
}

func (_ tAdClass) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdClass.Delete", args).Url
}

func (_ tAdClass) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdClass.Update", args).Url
}

func (_ tAdClass) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdClass.Getall", args).Url
}


type tAgentlevel struct {}
var Agentlevel tAgentlevel


func (_ tAgentlevel) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agentlevel.Save", args).Url
}

func (_ tAgentlevel) Delete(
=======
func (_ tMember) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Login", args).Url
}

func (_ tMember) Isuse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Isuse", args).Url
}


type tProduct struct {}
var Product tProduct


func (_ tProduct) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Save", args).Url
}

<<<<<<< HEAD
func (_ tProduct) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Update", args).Url
}

func (_ tProduct) Upframe(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Upframe", args).Url
}

func (_ tProduct) Downfrmae(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Downfrmae", args).Url
}

func (_ tProduct) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Getall", args).Url
}


type tRole struct {}
var Role tRole


func (_ tRole) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Save", args).Url
}

func (_ tRole) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Update", args).Url
}

func (_ tRole) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Getall", args).Url
}


type tCart struct {}
var Cart tCart


func (_ tCart) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cart.Save", args).Url
}

func (_ tCart) Addnum(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cart.Addnum", args).Url
}

func (_ tCart) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cart.Delete", args).Url
=======
func (_ tAdvertisement) Delete(
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agentlevel.Delete", args).Url
}

func (_ tAgentlevel) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agentlevel.Update", args).Url
}

func (_ tAgentlevel) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agentlevel.Getall", args).Url
}


type tCouponClass struct {}
var CouponClass tCouponClass


func (_ tCouponClass) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CouponClass.Save", args).Url
}

func (_ tCouponClass) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CouponClass.Delete", args).Url
}

func (_ tCouponClass) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CouponClass.Update", args).Url
}

func (_ tCouponClass) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CouponClass.Getall", args).Url
}


type tMember struct {}
var Member tMember


func (_ tMember) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Save", args).Url
}

func (_ tMember) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Update", args).Url
}

func (_ tMember) Hasbyname(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Hasbyname", args).Url
}

func (_ tMember) Hascode(
		) string {
	args := make(map[string]string)
	
<<<<<<< HEAD
	return revel.MainRouter.Reverse("Member.Hascode", args).Url
}

func (_ tMember) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Login", args).Url
=======
	return revel.MainRouter.Reverse("Agentlevel.Getall", args).Url
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f
}

func (_ tMember) Isuse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Member.Isuse", args).Url
}


type tProduct struct {}
var Product tProduct


func (_ tProduct) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Save", args).Url
}

func (_ tProduct) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Update", args).Url
}

func (_ tProduct) Upframe(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Upframe", args).Url
}

func (_ tProduct) Downfrmae(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Downfrmae", args).Url
}

<<<<<<< HEAD
func (_ tProduct) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Getall", args).Url
}
=======
<<<<<<< HEAD
type tRolegrant struct {}
var Rolegrant tRolegrant


func (_ tRolegrant) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Rolegrant.Save", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Userlist(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Userlist", args).Url
}

func (_ tUser) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Save", args).Url
}

func (_ tUser) Hascellphone(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Hascellphone", args).Url
=======
type tRole struct {}
var Role tRole
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f


type tRolegrant struct {}
var Rolegrant tRolegrant


func (_ tRolegrant) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Rolegrant.Save", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Userlist(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Userlist", args).Url
}

func (_ tUser) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Save", args).Url
}

func (_ tUser) Hascellphone(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Hascellphone", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


