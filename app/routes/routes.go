// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


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
	
	return revel.MainRouter.Reverse("Member.Hascode", args).Url
}

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


type tModular struct {}
var Modular tModular


func (_ tModular) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Modular.Save", args).Url
}

func (_ tModular) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Modular.Update", args).Url
}

func (_ tModular) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Modular.Getall", args).Url
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

func (_ tProduct) Getall(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Getall", args).Url
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
}

