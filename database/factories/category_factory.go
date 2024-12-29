package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/category"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.Username(),
			Description: faker.Sentence(),
		}
		objs = append(objs, categoryModel)
	}

	return objs
}
