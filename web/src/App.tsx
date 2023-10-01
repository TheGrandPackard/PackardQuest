import React from 'react';
import logo from './logo.svg';
import './App.css';
import RegistrationFormHogwarts, {UserData} from './components/RegistrationFormHogwarts';
import { RecoilRoot } from 'recoil';
import Router from './components/Router';

function App() {




  return (
      <>
      <RecoilRoot>
            <Router/>
      </RecoilRoot>
    </>
  );
}

export default App;