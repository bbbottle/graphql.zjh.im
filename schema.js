var graphql = require('graphql');

var queryType = new graphql.GraphQLObjectType({
    name: 'Query',
    fields: {
        hello: {
            type: graphql.GraphQLString,
            resolve: function () {
                return 'Hello World!';
            }
        }
    }
});

var schema = new graphql.GraphQLSchema({
    query: queryType
});
