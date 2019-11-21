package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coffee377/voc-admin/internal/app"
	"github.com/coffee377/voc-admin/internal/app/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("hello"))
//}

var (
	configFile string
	//modelFile  string
	//wwwDir     string
	//swaggerDir string
	//menuFile   string
)

func init() {
	flag.StringVar(&configFile, "c", "", "配置文件(.json,.yaml,.toml)")
	//flag.StringVar(&modelFile, "m", "", "Casbin的访问控制模型(.conf)")
	//flag.StringVar(&wwwDir, "www", "", "静态站点目录")
	//flag.StringVar(&swaggerDir, "swagger", "", "swagger目录")
	//flag.StringVar(&menuFile, "menu", "", "菜单数据文件(.json)")
}

//func init() {

//}

func main() {
	flag.Parse()

	// init config
	if err := config.Init(configFile); err != nil {
		panic(err)
	}

	app.InitLogger()

	app.Init(context.Background())
	//model := casbin.NewModel("conf/model.conf")

	//model.AddDef("r", "r", "sub, obj, act")
	//model.AddDef("p", "p", "sub, obj, act")
	//model.AddDef("e", "e", "some(where (p.eft == allow))")
	//model.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")

	//a := fileadapter.NewAdapter("conf/policy.csv")
	//a, _ := xormadapter.NewAdapter("mysql", "sysadmin:sysadmin@tcp(47.99.116.58:3306)/test", true)

	//e := casbin.NewEnforcer(model, a)

	// Load the policy from DB.
	//_ = e.LoadPolicy()

	// Check the permission.
	//sub := "coffee377"     // the user that wants to access a resource.
	//obj := "/portal/login" // the resource that is going to be accessed.
	//act := "read"          // the operation that the user performs on the resource.
	//s := e.Enforce(sub, obj, act)

	//if s {
	//	// permit alice to read data1
	//	log.Printf("YES %s -> %s -> %s", sub, obj, act)
	//} else {
	//	// deny the request, show an error
	//	log.Printf("NO %s -> %s -> %s", sub, obj, act)
	//}

	// Modify the policy.
	//e.AddRoleForUser("eve","ddd")
	//e.AddRoleForUser("coffee377", "ROLE_ADMIN")
	//e.DeletePermissionForUser("coffee377", "/*", "a", "b", "c", "d")
	//e
	//e.AddPolicy("eve", "data3", "read")
	//e.AddGroupingPolicy("group1", "data2_admin")
	//e.AddGroupingPolicy("ROLE_ADMIN", "/*")
	//e.DeleteRole("ROLE_ADMIN")
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	//e.SavePolicy()

	//roles, err := e.GetRolesForUser("alice")
	//
	//log.Printf("%s", roles)
	//log.Printf("%s", err)

	//app.Run(noRoute)
	//api.Run()

	// Ping the server to make sure the router is working.
	//go func() {
	//	if err := pingServer(); err != nil {
	//		log.Fatal("The router has no response, or it might took too long to start up.", err)
	//	}
	//	log.Print("The router has been deployed successfully.")
	//}()
	time.Sleep(time.Microsecond * 100)
	//beego.Run()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", helloHandler)
	//_ = http.ListenAndServe(":12345", mux)
}
