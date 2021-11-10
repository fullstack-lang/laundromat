package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/laundromat"
	laundromat_controllers "github.com/fullstack-lang/laundromat/go/controllers"
	laundromat_models "github.com/fullstack-lang/laundromat/go/models"
	laundromat_orm "github.com/fullstack-lang/laundromat/go/orm"

	gongsim_controllers "github.com/fullstack-lang/gongsim/go/controllers"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gongsim/go/orm"
	_ "github.com/fullstack-lang/gongsim/ng"

	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"
	_ "github.com/fullstack-lang/gongdoc/ng"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"
	_ "github.com/fullstack-lang/gong/ng"
)

var (
	logDBFlag         = flag.Bool("logDB", false, "log mode for db")
	logGINFlag        = flag.Bool("logGIN", false, "log mode for gin")
	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

//
// generic code
//
// specific code is in laundromat_engine
func main() {

	log.SetPrefix("laundromat: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// Gongsim
	if *clientControlFlag {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.CLIENT_CONTROL
	} else {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.AUTONOMOUS
	}

	// setup GORM
	db := laundromat_orm.SetupModels(*logDBFlag, ":memory:")
	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	// add gongdocatabase
	gongdoc_orm.AutoMigrate(db)

	// add gongsim database
	gongsim_orm.AutoMigrate(db)

	// add gong database
	gong_orm.AutoMigrate(db)

	//
	// stage gongdoc stack
	//
	var pkgelt gongdoc_models.Pkgelt
	pkgelt.Unmarshall("../../diagrams")
	pkgelt.SerializeToStage()

	//
	// stage simulation stack (to be done after the gongdoc load)
	//
	simulation := laundromat_models.NewSimulation()
	simulation.Stage()

	//
	// stage gong stack
	//
	modelPkg := &gong_models.ModelPkg{}
	gong_models.Walk("../../models", modelPkg)
	modelPkg.SerializeToStage()

	//
	//  setup controlers
	//
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	laundromat_controllers.RegisterControllers(r)
	gongsim_controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)

	r.Use(static.Serve("/", EmbedFolder(laundromat.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	// put all to database
	gongsim_models.Stage.Commit()
	gongdoc_models.Stage.Commit()
	laundromat_models.Stage.Commit()
	gong_models.Stage.Commit()

	log.Printf("simulation ready to run")
	r.Run()
	os.Exit(0)
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
