package rest

import "github.com/getkin/kin-openapi/openapi3"

func newOpenApi() *openapi3.T {
	swagger := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Event calendar API",
			Description: "REST API для календаря ивентов",
			Version:     "0.0.1",
		},
		Servers: []*openapi3.Server{
			{
				URL: "http://0.0.0.0:8000/api/v1",
			},
		},
		Tags: []*openapi3.Tag{
			{
				Name:        "tags",
				Description: "Тэги для мероприятий",
			},
			{
				Name:        "cities",
				Description: "Города проведения мероприятий",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Tag": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(map[string]*openapi3.Schema{
				"id":   openapi3.NewIntegerSchema(),
				"name": openapi3.NewStringSchema(),
			}),
		),
		"City": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(map[string]*openapi3.Schema{
				"id":       openapi3.NewIntegerSchema(),
				"name":     openapi3.NewIntegerSchema(),
				"timezone": openapi3.NewIntegerSchema(),
			}),
		),
		"Event": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(map[string]*openapi3.Schema{
				"id":        openapi3.NewIntegerSchema(),
				"name":      openapi3.NewStringSchema(),
				"startDate": openapi3.NewDateTimeSchema(),
				"endDate":   openapi3.NewDateTimeSchema(),
				"url":       openapi3.NewStringSchema().WithNullable(),
			}),
		),
		"EventPart": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(map[string]*openapi3.Schema{
				"id":          openapi3.NewIntegerSchema(),
				"name":        openapi3.NewStringSchema(),
				"description": openapi3.NewStringSchema().WithNullable(),
				"address":     openapi3.NewStringSchema().WithNullable(),
				"place":       openapi3.NewStringSchema(),
				"age":         openapi3.NewIntegerSchema().WithEnum(0, 6, 12, 16, 18),
				"startTime":   openapi3.NewDateTimeSchema(),
				"endTime":     openapi3.NewDateTimeSchema(),
			}),
		),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateTagsRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Тело запроса на создание пачки тэгов").
				WithRequired(true).WithJSONSchema(
				openapi3.NewObjectSchema().
					WithProperty("tags", openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema())),
			),
		},
		"UpdateCreateCityRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Тело запроса на создание города").
				WithRequired(true).
				WithJSONSchema(openapi3.NewObjectSchema().WithProperties(map[string]*openapi3.Schema{
					"name":     openapi3.NewStringSchema(),
					"timezone": openapi3.NewIntegerSchema(),
				})),
		},
		"UpdateTagRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Тело запроса на обновление названия тэга").
				WithRequired(true).WithJSONSchema(
				openapi3.NewObjectSchema().WithPropertyRef("", nil),
			),
		},
	}

	return nil
}
