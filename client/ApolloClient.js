import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client';

const client = new ApolloClient({
  link: new HttpLink({
    uri: 'https://food-recipe-21.hasura.app/v1/graphql', // Replace with your Hasura endpoint
    headers: {
      'x-hasura-admin-secret': 'sE6xQ4q5XKJYnxqn4qrgFRbe6hnE12awTmHIo5kEnUCkJR6SJEZa8d17wTsF6ifr' // Replace with your Hasura admin secret
    }
  }),
  cache: new InMemoryCache(),
});

export default client;
