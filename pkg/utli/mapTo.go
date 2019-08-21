package utli

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Json2Map(jsonBuf string) (map[string]interface{}, error) {
	maps := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonBuf), &maps)

	if err != nil {
		return maps, err
	}
	return maps, nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		k := typeOfType.Field(i).Name
		k = CamelString(k)

		data[k] = field.Interface()
	}
	return data
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		//return fmt.Errorf("No such field: %s in obj", name)
		fmt.Println("No such field: %s in obj", name)
		return nil
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		//i, err := strconv.Atoi(value)
		//return reflect.ValueOf(i), err
		i := StrTo(value).MustInt()
		return reflect.ValueOf(i), nil
	} else if ntype == "int8" {
		//i, err := strconv.ParseInt(value, 10, 64)
		//return reflect.ValueOf(int8(i)), err
		i := StrTo(value).MustInt()
		return reflect.ValueOf(i), nil

	} else if ntype == "int32" {
		//i, err := strconv.ParseInt(value, 10, 64)
		//return reflect.ValueOf(int64(i)), err
		i := StrTo(value).MustInt()
		return reflect.ValueOf(i), nil
	} else if ntype == "int64" {
		//i, err := strconv.ParseInt(value, 10, 64)
		//return reflect.ValueOf(i), err

		i := StrTo(value).MustInt()
		return reflect.ValueOf(i), nil
	} else if ntype == "float32" {
		//i, err := strconv.ParseFloat(value, 64)
		//return reflect.ValueOf(float32(i)), err

		i := StrTo(value).MustFloat64()
		return reflect.ValueOf(i), nil
	} else if ntype == "float64" {
		//i, err := strconv.ParseFloat(value, 64)
		//return reflect.ValueOf(i), err
		i := StrTo(value).MustFloat64()
		return reflect.ValueOf(i), nil
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
