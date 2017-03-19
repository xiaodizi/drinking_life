// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "drinking_life/app"
	controllers "drinking_life/app/controllers"
	_ "drinking_life/app/models"
	tests "drinking_life/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
<<<<<<< HEAD
	revel.RegisterController((*controllers.Advertisement)(nil),
=======
	revel.RegisterController((*controllers.AdClass)(nil),
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Getall",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
<<<<<<< HEAD
	revel.RegisterController((*controllers.Member)(nil),
=======
	revel.RegisterController((*controllers.Agent)(nil),
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Update",
=======
				Name: "Delete",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Hasbyname",
=======
				Name: "Update",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Hascode",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Isuse",
=======
				Name: "Getall",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
<<<<<<< HEAD
	revel.RegisterController((*controllers.Product)(nil),
=======
	revel.RegisterController((*controllers.Member)(nil),
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Upframe",
=======
				Name: "Hasbyname",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Downfrmae",
=======
				Name: "Hascode",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Getall",
=======
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Isuse",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
<<<<<<< HEAD
	revel.RegisterController((*controllers.Role)(nil),
=======
	revel.RegisterController((*controllers.Product)(nil),
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Getall",
=======
				Name: "Upframe",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Cart)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Save",
=======
				Name: "Downfrmae",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Addnum",
=======
				Name: "Getall",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Rolegrant)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Delete",
=======
				Name: "Save",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Userlist",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Hascellphone",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
<<<<<<< HEAD
	revel.RegisterController((*controllers.Rolegrant)(nil),
=======
	revel.RegisterController((*controllers.Advertisement)(nil),
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Userlist",
=======
				Name: "Delete",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Getall",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
<<<<<<< HEAD
				Name: "Save",
=======
				Name: "Sort",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Agentlevel)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
>>>>>>> bf010d9cd0c3733b8cf6a9f2dab02f291c05810e
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Hascellphone",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Modular)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Getall",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Role)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Getall",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
