import React from 'react';
import ReactDOM from 'react-dom/client';
import SignUp from './components/auth/signup.tsx';
import SignIn from './components/auth/signin.tsx';


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <div>
    <SignUp/>
    hello world
    <SignIn></SignIn>
  </div>
  

);
