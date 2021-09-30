package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var availableValidationTags = map[reflect.Kind]map[string]string{
	reflect.String: {
		"len":    "StringLen",    // `len:32` - длина строки должна быть ровно 32 символа;
		"regexp": "StringRegexp", // `regexp:\\d+` - согласно регулярному выражению строка должна состоять из цифр
		"in":     "StringIn",     // `min:10` - число не может быть меньше 10;
	},
	reflect.Int: {
		"min": "IntMin", // `min:10` - число не может быть меньше 10;
		"max": "IntMax", // `max:20` - число не может быть больше 20;
		"in":  "IntIn",  // `in:256,1024` - число должно входить в множество чисел {256, 1024};
	},
	reflect.Int8: {
		"notImplement": "notImplement", // - не реализованый метод для тестов
	},
}

var validateTagName = "validate"

type validationTag struct {
	tagName  string
	tagValue string
}

func (v *validationTag) getMethodNameForTag(valueType reflect.Kind) (methodName string, ok bool) {
	if tagsMap, ok := availableValidationTags[valueType]; ok {
		if methodName, ok := tagsMap[v.tagName]; ok {
			st := reflect.TypeOf(v)
			_, ok := st.MethodByName(methodName)
			return methodName, ok
		}
	}
	return "", false
}

func (v *validationTag) validate(value interface{}) []error {
	valueType := reflect.TypeOf(value).Kind()
	var valuesSlice []interface{}
	switch valueType {
	case reflect.Slice:
		s := reflect.ValueOf(value)
		if s.Len() == 0 {
			return []error{fmt.Errorf("длинна слайса '%s' = 0", value)}
		}
		valuesSlice = make([]interface{}, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			valuesSlice = append(valuesSlice, s.Index(i).Interface())
		}
		valueType = reflect.TypeOf(value).Elem().Kind()
	default:
		valuesSlice = append(valuesSlice, value)
	}

	methodName, ok := v.getMethodNameForTag(valueType)
	if !ok {
		errStr := fmt.Sprintf("отсутствует метод валидации для тэга: %s и типа значения %s", v.tagName, valueType)
		if methodName != "" {
			errStr += fmt.Sprintf("\nне реализован метод: %s(value interface{}) (err error)", methodName)
		}
		return []error{fmt.Errorf(errStr)}
	}

	// вызываем метод валидации для конкретного типа значения
	var errors []error
	for _, val := range valuesSlice {
		validateMethodRes := reflect.ValueOf(v).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(val)})
		err := validateMethodRes[0].Interface()
		if err != nil {
			errors = append(errors, err.(error))
		}
	}
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *validationTag) StringLen(value interface{}) (err error) {
	equalLen, convErr := strconv.Atoi(v.tagValue)
	if convErr != nil {
		return fmt.Errorf("'%s' не получилось преобразовать значение = '%s' в число", v.tagName, v.tagValue)
	}
	valueStr := fmt.Sprintf("%v", value)
	strLen := utf8.RuneCountInString(valueStr)
	if strLen != equalLen {
		return fmt.Errorf("длина строки '%s' должна быть '%d', а не '%d'", valueStr, equalLen, strLen)
	}

	return nil
}

func (v *validationTag) StringIn(value interface{}) (err error) {
	inValuesStr := strings.Split(v.tagValue, ",")

	valueStr := fmt.Sprintf("%v", value)

	for _, inV := range inValuesStr {
		if valueStr == inV {
			return nil
		}
	}
	return fmt.Errorf("'%s' должен находиться в массиве '%v'", valueStr, inValuesStr)
}

func (v *validationTag) StringRegexp(value interface{}) (err error) {
	regexpValue := v.tagValue
	valueStr := fmt.Sprintf("%v", value)
	matched, err := regexp.Match(regexpValue, []byte(valueStr))
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("значение '%s' не соответвует регулярному выражению '%s'", valueStr, regexpValue)
	}

	return nil
}

func (v *validationTag) IntMin(value interface{}) (err error) {
	minValue, convErr := strconv.Atoi(v.tagValue)
	if convErr != nil {
		return fmt.Errorf("'%s' не получилось преобразовать значение = '%s' в число", v.tagName, v.tagValue)
	}
	intValue, ok := value.(int)
	if !ok {
		return fmt.Errorf("ошибка преобразования в целое '%d'", value)
	}

	if intValue < minValue {
		return fmt.Errorf("число '%d' должно быть не меньше чем '%d'", intValue, minValue)
	}

	return nil
}

func (v *validationTag) IntMax(value interface{}) (err error) {
	minValue, convErr := strconv.Atoi(v.tagValue)
	if convErr != nil {
		return fmt.Errorf("%w; '%s' не получилось преобразовать значение = '%s' в число", convErr, v.tagName, v.tagValue)
	}
	intValue, ok := value.(int)
	if !ok {
		return fmt.Errorf("ошибка преобразования в целое '%d'", value)
	}
	if intValue > minValue {
		return fmt.Errorf("число '%d' должно быть не больше чем '%d'", intValue, minValue)
	}

	return nil
}

func (v *validationTag) IntIn(value interface{}) (err error) {
	inValuesStr := strings.Split(v.tagValue, ",")
	inMap := map[int]interface{}{}

	valueInt, ok := value.(int)
	if !ok {
		return fmt.Errorf("ошибка преобразования в целое '%d'", value)
	}

	for _, inV := range inValuesStr {
		inVInt, convErr := strconv.Atoi(inV)
		if convErr != nil {
			errM := fmt.Errorf("%w; не получилось преобразовать в число '%s'", convErr, inV)
			if err != nil {
				err = fmt.Errorf("%w; %s", err, errM)
			} else {
				err = errM
			}
		} else {
			inMap[inVInt] = nil
		}
	}
	if err != nil {
		return err
	}

	if _, ok := inMap[valueInt]; !ok {
		return fmt.Errorf("'%s' должен находиться в массиве '%v'", value, inValuesStr)
	}

	return nil
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	resMessage := make([]string, 0, len(v))
	for _, valErr := range v {
		if valErr.Err != nil {
			resMessage = append(
				resMessage,
				fmt.Sprintf("ошибка в поле: %s \n %s", valErr.Field, valErr.Err.Error()),
			)
		}
	}
	return strings.Join(resMessage, "\n")
}

func Validate(v interface{}) error {
	value := reflect.ValueOf(v)
	fmt.Println(value.Kind())
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, but received %T", v)
	}

	t := value.Type()
	var resValidationErrors ValidationErrors
	for i := 0; i < t.NumField(); i++ {
		fieldInfo := t.Field(i)
		fieldValue := value.Field(i)
		errV, ok := validateField(fieldInfo, fieldValue)
		if !ok {
			resValidationErrors = append(resValidationErrors, errV)
		}
	}

	if len(resValidationErrors) > 0 {
		return resValidationErrors
	}

	return nil
}

func validateField(fieldInfo reflect.StructField, fieldValue reflect.Value) (err ValidationError, ok bool) {
	if !fieldValue.CanInterface() { // игнорируем приватные поля структуры
		return ValidationError{}, true
	}

	tagValue := fieldInfo.Tag.Get(validateTagName)
	if tagValue == "" {
		return ValidationError{}, true
	}

	validateTags := getValidateTags(tagValue)
	if len(validateTags) == 0 {
		return ValidationError{
				Field: fieldInfo.Name,
				Err:   fmt.Errorf("не удалось создать validateTags для значения: %s ", tagValue),
			},
			false
	}

	valueInterface := fieldValue.Interface()

	var validateErrors []error
	for _, validateTag := range validateTags {
		errs := validateTag.validate(valueInterface)
		if len(errs) > 0 {
			validateErrors = append(validateErrors, errs...)
		}
	}

	if len(validateErrors) > 0 {
		var resErr error
		for _, err := range validateErrors {
			if resErr != nil {
				resErr = fmt.Errorf("%w; %s", resErr, err.Error())
			} else {
				resErr = err
			}
		}
		return ValidationError{
			Field: fieldInfo.Name,
			Err:   resErr,
		}, false
	}
	return ValidationError{}, true
}

/*
Создаем []validationTag из описания поля
`validate:"min:18|max:1"` создаст []validationTag{[min]}.
*/
func getValidateTags(tagValue string) []validationTag {
	if len(tagValue) == 0 {
		return nil
	}

	tags := strings.Split(tagValue, "|")
	res := make([]validationTag, 0, len(tags))
	for _, tag := range tags {
		name, value := func() (string, string) {
			x := strings.Split(tag, ":")
			x1 := x[0]
			x2 := ""
			if len(x) == 2 {
				x2 = x[1]
			} else if len(x) > 2 {
				x2 = strings.Join(x[1:], ":")
			}
			return x1, x2
		}()
		res = append(res, validationTag{
			tagName:  name,
			tagValue: value,
		})
	}
	return res
}
