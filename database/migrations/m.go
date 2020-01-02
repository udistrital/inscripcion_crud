package main

import(
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/migration"

	_ "github.com/lib/pq"
)

func init(){
	orm.RegisterDataBase("default", "postgres","postgres://desarrollooas:test@inscripcion_db/test?sslmode=disable&search_path=public")
}

func main(){
	task := "upgrade"
	switch task {
	case "upgrade":
		if err := migration.Upgrade(1577779743); err != nil {
			os.Exit(2)
		}
	case "rollback":
		if err := migration.Rollback("AgregarCamposGrupoInvestigacion_20191231_080903"); err != nil {
			os.Exit(2)
		}
	case "reset":
		if err := migration.Reset(); err != nil {
			os.Exit(2)
		}
	case "refresh":
		if err := migration.Refresh(); err != nil {
			os.Exit(2)
		}
	}
}

