package main

import (
    "fmt"
    "reflect"
)

type Model interface {
    Save() string
}

type MyModel struct {
    I int
    B bool
    F float64
}

func (m *MyModel) Save() string {
    return reflectModel(m)
}

func reflectModel(model Model) string {
    modelType := reflect.TypeOf(model)

    if modelType.Kind() == reflect.Ptr {
        modelType = modelType.Elem()
    }

    s := fmt.Sprintf("%v struct {\n", modelType.Name())

    for i := 0; i < modelType.NumField(); i++ {
        field := modelType.Field(i)
        s += fmt.Sprintf("\t%v %v\n", field.Name, field.Type.Name())
    }

    s += "}"

    return s
}
