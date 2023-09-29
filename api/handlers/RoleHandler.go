// api/handlers/RoleHandler.go

package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/defrindr/gowisata/models"
    "github.com/defrindr/gowisata/repositories"
)

func CreateRole(c *fiber.Ctx) error {
    role := new(models.Role)
    if err := c.BodyParser(role); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request payload",
        })
    }

    if err := repositories.CreateRole(role); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error creating role",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Role created successfully",
        "role":    role,
    })
}

func GetAllRoles(c *fiber.Ctx) error {
    roles, err := repositories.GetAllRoles()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error retrieving roles",
        })
    }
    return c.JSON(fiber.Map{
        "roles": roles,
    })
}

// Implement similar functions for GetRoleByID, GetAllRoles, UpdateRole, and DeleteRole
