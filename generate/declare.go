package generate

import (
	"github.com/ekokurniadi/micagen"
	"gorm.io/gorm"
)

func Generate(db *gorm.DB, model interface{}) {
	micagen.GenerateTable(db, model)
	micagen.CreateStructInput(model)
	micagen.CreateService(db, model)
	micagen.CreateRepository(db, model)
	micagen.CreateHandler(db, model)
	micagen.CreateFormatter(db, model)
}
