package container

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goioc/di"
	"reflect"
	"simple-web/helper"
	"simple-web/internal/domain/repository"
	"simple-web/internal/domain/service"
	"simple-web/internal/infrastructure/database"
	"simple-web/internal/interface/api/controller"
)

type bean struct {
	beanType reflect.Type
	beanID   string
}

func registerBean(beans ...bean) {
	for _, b := range beans {
		_, err := di.RegisterBean(b.beanID, b.beanType)
		helper.PanicIfError(err)
	}
}

func InitGoioc() {
	registerBean(
		bean{
			beanID:   "userRepository",
			beanType: reflect.TypeOf((*repository.UserRepositoryImpl)(nil)),
		},
		bean{
			beanID:   "userService",
			beanType: reflect.TypeOf((*service.UserServiceImpl)(nil)),
		},
		bean{
			beanID:   "userController",
			beanType: reflect.TypeOf((*controller.UserControllerImpl)(nil)),
		},
	)

	_, err := di.RegisterBeanInstance("database", database.NewDB())
	helper.PanicIfError(err)

	_, err = di.RegisterBeanInstance("validate", validator.New())
	helper.PanicIfError(err)

	err = di.InitializeContainer()
	helper.PanicIfError(err)
}
