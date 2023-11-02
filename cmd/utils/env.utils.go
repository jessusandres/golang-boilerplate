package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func ParseEnvSchema[T any](cfg *T) []string {
	schemaType := reflect.TypeOf(*cfg)
	schemaValue := reflect.ValueOf(cfg).Elem()

	if schemaType.Kind() != reflect.Struct {
		// Fatal logs terminate the app by os.Exit(1)
		log.Fatal("❌  Provided schema is not a struct")

		return []string{}
	}

	schemaKeyLength := schemaType.NumField()
	errorsSlice := make([]string, 0, schemaKeyLength)

	for i := 0; i < schemaKeyLength; i++ {
		field := schemaType.Field(i)

		envKey := field.Tag.Get("env")
		required := field.Tag.Get("required")
		defaultValue := field.Tag.Get("default")

		if envKey == "" {
			log.Printf("⚠️  Field '%s' does not have an 'env' tag, skipping", field.Name)
			continue
		}

		envValue := os.Getenv(envKey)

		if envValue == "" && (required == "true") {
			errorString := fmt.Sprintf("❌ Required environment variable '%s' is not set", envKey)

			errorsSlice = append(errorsSlice, errorString)
		}

		fieldValue := schemaValue.Field(i)
		if fieldValue.CanSet() {
			v := envValue

			switch fieldValue.Kind() {
			case reflect.String:
				if v == "" {
					v = defaultValue
				}

				fieldValue.SetString(v)
			case reflect.Bool:
				if v == "" {
					v = defaultValue
				}

				fieldValue.SetBool(v == "true")
			case reflect.Int:
				if v == "" {
					v = defaultValue
				}

				value, err := strconv.Atoi(v)

				if err != nil {
					errorsSlice = append(errorsSlice, fmt.Sprintf("❌ Key '%s' is not a valid integer", envKey))
				}

				fieldValue.SetInt(int64(value))
			default:
				log.Fatalf("❌ Unsupported field type for '%s' variable", envKey)
			}

		}

	}

	return errorsSlice
}
