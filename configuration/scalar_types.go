package configuration

// scalartypes := schema.SchemaResponseScalarTypes{
// 	"Int": schema.ScalarType{
// 		AggregateFunctions: schema.ScalarTypeAggregateFunctions{
// 			"max": schema.AggregateFunctionDefinition{
// 				ResultType: schema.NewNullableNamedType("Int").Encode(),
// 			},
// 			"min": schema.AggregateFunctionDefinition{
// 				ResultType: schema.NewNullableNamedType("Int").Encode(),
// 			},
// 		},
// 		ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
// 			"eq": schema.NewComparisonOperatorCustom(schema.NewNamedType("Int")).Encode(),
// 			"in": schema.NewComparisonOperatorCustom(schema.NewNamedType("Int")).Encode(),
// 			// "eq": schema.NewComparisonOperatorEqual().Encode(),
// 			// "in": schema.NewComparisonOperatorIn().Encode(),
// 		},
// 	},
// 	"Float": schema.ScalarType{
// 		AggregateFunctions: schema.ScalarTypeAggregateFunctions{
// 			"max": schema.AggregateFunctionDefinition{
// 				ResultType: schema.NewNullableNamedType("Float").Encode(),
// 			},
// 			"min": schema.AggregateFunctionDefinition{
// 				ResultType: schema.NewNullableNamedType("Float").Encode(),
// 			},
// 		},
// 		ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
// 			// "eq": schema.NewComparisonOperatorEqual().Encode(),
// 			// "in": schema.NewComparisonOperatorIn().Encode(),
// 			"eq": schema.NewComparisonOperatorCustom(schema.NewNamedType("Float")).Encode(),
// 			"in": schema.NewComparisonOperatorCustom(schema.NewNamedType("Float")).Encode(),
// 		},
// 	},
// 	"String": schema.ScalarType{
// 		AggregateFunctions: schema.ScalarTypeAggregateFunctions{},
// 		ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
// 			// "eq":   schema.NewComparisonOperatorEqual().Encode(),
// 			// "in":   schema.NewComparisonOperatorIn().Encode(),
// 			"eq":   schema.NewComparisonOperatorCustom(schema.NewNamedType("String")).Encode(),
// 			"in":   schema.NewComparisonOperatorCustom(schema.NewNamedType("String")).Encode(),
// 			"like": schema.NewComparisonOperatorCustom(schema.NewNamedType("String")).Encode(),
// 		},
// 	},
// }

import (
	"context"

	"github.com/hasura/ndc-sdk-go/schema"
)

type AggregateFunction struct {
	Name       string
	ResultType string
}

type ComparisonOperators struct {
	Name         string
	ArgumentType string
}

type ScalarType struct {
	AggregateFunction   []AggregateFunction
	ComparisonOperators []ComparisonOperators
}

var scalar_types = map[string]ScalarType{
	"integer": {
		AggregateFunction: []AggregateFunction{
			{
				Name:       "max",
				ResultType: "integer",
			},
			{
				Name:       "min",
				ResultType: "integer",
			},
			{
				Name:       "sum",
				ResultType: "integer",
			},
			{
				Name:       "avg",
				ResultType: "integer",
			},
			{
				Name:       "cardinality",
				ResultType: "integer",
			},
			{
				Name:       "percentiles",
				ResultType: "integer",
			},
			{
				Name:       "stats",
				ResultType: "integer",
			},
		},
		ComparisonOperators: []ComparisonOperators{
			{
				Name:         "eq",
				ArgumentType: "integer",
			},
			{
				Name:         "in",
				ArgumentType: "integer",
			},
		},
	},
	"keyword": {
		AggregateFunction: []AggregateFunction{
			{
				Name:       "cardinality",
				ResultType: "integer",
			},
		},
		ComparisonOperators: []ComparisonOperators{
			{
				Name:         "eq",
				ArgumentType: "keyword",
			},
			{
				Name:         "in",
				ArgumentType: "keyword",
			},
			{
				Name:         "like",
				ArgumentType: "keyword",
			},
		},
	},
}

func getScalarTypes(scalartype map[string]bool, ctx context.Context) (schema.SchemaResponseScalarTypes, error) {
	scalartypes := make(schema.SchemaResponseScalarTypes)
	for datatype, _ := range scalartype {
		var scalar schema.ScalarType
		scalar.AggregateFunctions = make(schema.ScalarTypeAggregateFunctions)
		for _, n := range scalar_types[datatype].AggregateFunction {
			scalar.AggregateFunctions[n.Name] = schema.AggregateFunctionDefinition{
				ResultType: schema.NewNullableNamedType(datatype).Encode(),
			}
		}

		scalar.ComparisonOperators = make(map[string]schema.ComparisonOperatorDefinition)
		for _, n := range scalar_types[datatype].ComparisonOperators {
			scalar.ComparisonOperators[n.Name] = schema.NewComparisonOperatorCustom(schema.NewNamedType(datatype)).Encode()
		}

		scalartypes[datatype] = scalar
	}
	return scalartypes, nil
}
