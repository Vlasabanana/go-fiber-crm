package lead

import (
	"github.com/Vlasabanana/go-fiber-crm/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {

}

func GetLead() {

}

func NewLead() {

}

func DeleteLead() {

}

