import './App.css';
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