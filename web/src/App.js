import React from 'react';
import { gql, useQuery } from '@apollo/client';
import './App.css';

const GET_BOOKS = gql`
  query GetBooks {
    books {
      title
    }
  }
`;

function App() {
  const { loading, error, data } = useQuery(GET_BOOKS);
  console.log({loading}, {data}, data.books)
  if (loading) return <p>Loading ...</p>;
  return (
    <div className="App">
      {data.books.map(b => b.title)}
    </div>
  );
 }
 
 export default App;