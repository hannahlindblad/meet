const { ApolloServer, gql } = require('apollo-server');
const { RESTDataSource } = require('apollo-datasource-rest');
const { MemcachedCache } = require('apollo-server-cache-memcached');

class AuthAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = 'http://localhost:8000';
  }

  async signUp(user) {
    return this.post(`signup`, user);
  }
}

// A schema is a collection of type definitions (hence "typeDefs")
// that together define the "shape" of queries that are executed against
// your data.
const typeDefs = gql`
  # This "Book" type defines the queryable fields for every book in our data source.
  type Book {
    title: String
    author: String
  }

  type User {
    first_name: String
    last_name: String
  }

  type Mutation {
    signUp(first_name: String, last_name: String, email: String, password: String): User
  }

  # The "Query" type is special: it lists all of the available queries that
  # clients can execute, along with the return type for each. In this
  # case, the "books" query returns an array of zero or more Books (defined above).
  type Query {
    books: [Book]
  }
`;

const books = [
    {
      title: 'The Awakening',
      author: 'Kate Chopin',
    },
    {
      title: 'City of Glass',
      author: 'Paul Auster',
    },
  ];

  // Resolvers define the technique for fetching the types defined in the
// schema. This resolver retrieves books from the "books" array above.
const resolvers = {
    Query: {
        books: () => books
    },
    Mutation: {
        signUp: async (_source, user, { dataSources }) => {
            console.log(user)
            // return {"first_name": "ha", "last_name": "ha"}
            return dataSources.authAPI.signUp(user);
        },
    }
  };

  // The ApolloServer constructor requires two parameters: your schema
// definition and your set of resolvers.
const server = new ApolloServer({ 
    typeDefs, 
    resolvers,
    // cache: new MemcachedCache(
    //     ['memcached-server-1', 'memcached-server-2', 'memcached-server-3'],
    //     { retries: 10, retry: 10000 }, // Options
    // ),
    dataSources: () => {
        return {
            authAPI: new AuthAPI()
        }
    }
});

// The `listen` method launches a web server.
server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});