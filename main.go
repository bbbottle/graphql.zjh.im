package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "Hello World!", nil
            },
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
