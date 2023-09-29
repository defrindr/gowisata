// repositories/RoleRepository.go

package repositories

import (
    "github.com/defrindr/gowisata/models"
)

func CreateRole(role *models.Role) error {
    if err := DB.Create(role).Error; err != nil {
        return err
    }
    return nil
}

func GetRoleByID(id uint) (*models.Role, error) {
    var role models.Role
    if err := DB.First(&role, id).Error; err != nil {
        return nil, err
    }
    return &role, nil
}

func GetAllRoles() ([]models.Role, error) {
    var roles []models.Role
    if err := DB.Find(&roles).Error; err != nil {
        return nil, err
    }
    return roles, nil
}

func UpdateRole(role *models.Role) error {
    if err := DB.Save(role).Error; err != nil {
        return err
    }
    return nil
}

func DeleteRole(id uint) error {
    if err := DB.Delete(&models.Role{}, id).Error; err != nil {
        return err
    }
    return nil
}
