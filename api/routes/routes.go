// api/routes/RoleRoutes.go

package routes

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/defrindr/gowisata/handlers"
)

func SetupRoleRoutes(app *fiber.App) {
	panic("aaaa")
    app.Post("/roles", handlers.CreateRole)
    app.Get("/roles/:id", handlers.GetRoleByID)
    app.Get("/roles", handlers.GetAllRoles)
    app.Put("/roles/:id", handlers.UpdateRole)
    app.Delete("/roles/:id", handlers.DeleteRole)
}
