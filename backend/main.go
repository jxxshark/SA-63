package main

import (
	"context"
	"log"
	"fmt"
    "github.com/jxxshark/app/ent/user"
	"github.com/jxxshark/app/controllers"
	_ "github.com/jxxshark/app/docs"
	"github.com/jxxshark/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	NAME  string
	EMAIL string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	NAME  string
	AGE   int
	
}


type Specializeddiags struct {
	Specializeddiag []Specializeddiag
}

type Specializeddiag struct {
	Specializeddiagtype string
	Specialistid int
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:test.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewSpecializeddiagController(v1, client)
	controllers.NewSpecializedappointController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"นพ.ทดสอบ 1", "todsorb1@example.com"},
			User{"นพ.ทดสอบ 2", "todsorb2@example.com"},
			User{"นพ.ทดสอบ 3", "todsorb3@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUseremail(u.EMAIL).
			SetUsername(u.NAME).
			Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"คนไข้ 1",20 },
			Patient{"คนไข้ 2",40 },
			Patient{"คนไข้ 3",30 },
		},
	}
	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetName(p.NAME).
		    SetAge(p.AGE).
			Save(context.Background())
	}

	
  // Set Specializeddiags Data
	specializeddiags := Specializeddiags{
		Specializeddiag: []Specializeddiag{
			Specializeddiag{"ตรวจสุขภาพ",1},
			Specializeddiag{"ตรวจภายใน",3},
			Specializeddiag{"ตรวจจักษุ",2},
		},
	}

	for _, sd := range specializeddiags.Specializeddiag {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(sd.Specialistid))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Specializeddiag.
			Create().
			SetSpecializeddiacnostictype(sd.Specializeddiagtype).
			SetUser(u).
			Save(context.Background())
	}
  
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
	
}

