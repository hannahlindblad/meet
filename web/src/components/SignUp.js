import React from 'react';
import { gql, useMutation } from '@apollo/client';
import { useFormik } from "formik";
import { Link } from "react-router-dom";

const SIGN_UP = gql`
  mutation SignUp($first_name: String!, $last_name: String!, $email: String!, $password: String!) {
    signUp(first_name: $first_name, last_name: $last_name, email: $email, password: $password) {
      first_name,
      last_name
    }
  }
`;

export function SignUp() {
  const [signUp, error] = useMutation(SIGN_UP);
  const formik = useFormik({
    initialValues: {
      firstName: '',
      lastName: '',
      email: '',
      passwpord: ''
    },
    onSubmit: values => {
      signUp({variables: { 
        first_name: values.firstName, 
        last_name: values.lastName, 
        email: values.email, 
        password: values.password 
      }});
    },
  });
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Create an account
          </h2>
        </div>
        <form className="mt-8 space-y-6" action="#" method="POST">
          <input type="hidden" name="remember" value="true" />
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <label for="email-address" className="sr-only">Email address</label>
              <input 
                id="email-address"
                name="email"
                type="email"      
                onChange={formik.handleChange}
                value={formik.values.email}
                autocomplete="email"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
            </div>
            <div>
              <label for="password" className="sr-only">Password</label>
              <input 
                id="password"
                name="password"
                type="password"
                onChange={formik.handleChange}
                value={formik.values.password}
                autocomplete="current-password"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Password" />
            </div>
            <div>
              <label for="first_name" className="sr-only">First name</label>
              <input 
                id="first_name"
                name="first_name"
                type="password"
                onChange={formik.handleChange}
                value={formik.values.password}
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="First name" />
            </div>
            <div>
              <label for="last_name" className="sr-only">Last name</label>
              <input 
                id="last_name"
                name="last_name"
                type="last_name"
                onChange={formik.handleChange}
                value={formik.values.password}
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Last name" />
            </div>
          </div>

          <div>
            <button type="submit" className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Create an account
            </button>
          </div>

          <div className="flex items-center justify-end">
            <div className="text-sm">
            <Link className="font-medium text-indigo-600 hover:text-indigo-500" to="/">← Back to sign in</Link>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
 }