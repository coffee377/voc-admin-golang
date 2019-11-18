package main

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist/file-adapter"
	"github.com/coffee377/voc-admin/model"
	_ "github.com/coffee377/voc-admin/router"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("hello"))
//}

func init() {
	// set default database
	_ = orm.RegisterDataBase("default", "mysql", "sysadmin:sysadmin@tcp(47.99.116.58:3306)/test?charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModelWithPrefix("", new(model.User), new(model.Role))

	//	// update
	//	user.Name = "astaxie"
	//	num, err := o.Update(&user)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//	// read one
	//	u := User{Id: user.Id}
	//	err = o.Read(&u)
	//	fmt.Printf("ERR: %v\n", err)
	//
	//	// delete
	//	num, err = o.Delete(&u)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// create table
	_ = orm.RunSyncdb("default", false, true)

	//o := orm.NewOrm()
	//
	//user := model.User{}
	//user.Id = "2"
	//user.Username = "demo2"
	//user.Status = 1
	//user.CreateTime = time.Now()
	////	// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

func main() {
	model := casbin.NewModel("conf/model.conf")

	//model.AddDef("r", "r", "sub, obj, act")
	//model.AddDef("p", "p", "sub, obj, act")
	//model.AddDef("e", "e", "some(where (p.eft == allow))")
	//model.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")

	a := fileadapter.NewAdapter("conf/policy.csv")
	//a, _ := xormadapter.NewAdapter("mysql", "sysadmin:sysadmin@tcp(47.99.116.58:3306)/test", true)

	e := casbin.NewEnforcer(model, a)

	// Load the policy from DB.
	_ = e.LoadPolicy()

	// Check the permission.
	sub := "coffee377"     // the user that wants to access a resource.
	obj := "/portal/login" // the resource that is going to be accessed.
	act := "read"          // the operation that the user performs on the resource.
	s := e.Enforce(sub, obj, act)

	if s {
		// permit alice to read data1
		log.Printf("YES %s -> %s -> %s", sub, obj, act)
	} else {
		// deny the request, show an error
		log.Printf("NO %s -> %s -> %s", sub, obj, act)
	}

	// Modify the policy.
	//e.AddRoleForUser("eve","ddd")
	e.AddRoleForUser("coffee377", "ROLE_ADMIN")
	e.DeletePermissionForUser("coffee377", "/*", "a", "b", "c", "d")
	//e
	//e.AddPolicy("eve", "data3", "read")
	e.AddGroupingPolicy("group1", "data2_admin")
	e.AddGroupingPolicy("ROLE_ADMIN", "/*")
	e.DeleteRole("ROLE_ADMIN")
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()

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

	//beego.Run()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", helloHandler)
	//_ = http.ListenAndServe(":12345", mux)
}

//func noRoute(c *gin.Context) {
//	c.JSON(404, result.Of(nil, error2.NotFound))
//}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
