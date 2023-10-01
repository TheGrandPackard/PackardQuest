
import * as React from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import { FC, useEffect } from 'react';
import { useRecoilState } from 'recoil';
import PlayerState from '../types/Player';

// https://blog.devgenius.io/implementing-react-router-v6-with-code-splitting-in-a-react-typescript-project-14d98e2cab79

const Loading: FC = () => <p>Loading ...</p>;

const Registration = React.lazy(() => import('./RegistrationFormHogwarts'));
const Sorting = React.lazy(() => import('./SortingHat'));

const Router: FC = () => {
    const navigate = useNavigate();
    const [player, setPlayer] = useRecoilState(PlayerState);
  
    // Load token from local storage and set into recoil state
    const localStorageToken = localStorage.getItem("player");
    if(localStorageToken && player === undefined) {
      console.log(localStorageToken);
      setPlayer(JSON.parse(localStorageToken));
    }
    
    // If the player has not registered, navigate to registration
    useEffect(() => {
      if(player === undefined) {
        navigate('/registration', { replace: true });
      }
    }, [navigate, player])

    return (
        <React.Suspense fallback={<Loading/>}>
            <Routes>
                <Route path='/' element={<Sorting/>}/>
                <Route path='/registration' element={<Registration/>}/>
            </Routes>
        </React.Suspense>
    )
}

export default Router;
