const { ApolloServer, gql } = require('apollo-server');
const { RESTDataSource } = require('apollo-datasource-rest');

class AuthAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = 'http://localhost:8000';
  }

  async signUp(user) {
    return this.post(`signup`, user);
  }
}

const typeDefs = gql`
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

const resolvers = {
    Query: {
        books: () => books
    },
    Mutation: {
        signUp: async (_source, user, { dataSources }) => {
            console.log(user)
            return dataSources.authAPI.signUp(user);
        },
    }
  };

const server = new ApolloServer({ 
    typeDefs, 
    resolvers,
    // cache: new MemcachedCache( TODO: Add cache
    //     ['memcached-server-1', 'memcached-server-2', 'memcached-server-3'],
    //     { retries: 10, retry: 10000 }, // Options
    // ),
    dataSources: () => {
        return {
            authAPI: new AuthAPI()
        }
    }
});

server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});