package project

import (
    "gohub/pkg/database"
)

func Get(idstr string) (project Project) {
    database.DB.Where("id", idstr).First(&project)
    return
}

func GetBy(field, value string) (project Project) {
    database.DB.Where("? = ?", field, value).First(&project)
    return
}

func All() (projects []Project) {
    database.DB.Find(&projects)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Project{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}