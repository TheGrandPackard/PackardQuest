
import * as React from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import { FC, useEffect } from 'react';
import { useRecoilState } from 'recoil';
import PlayerState from '../types/Player';

const Loading: FC = () => <p>Loading ...</p>;
const Home = React.lazy(() => import('./Home'));
const Registration = React.lazy(() => import('./Registration'));

const Router: FC = () => {
    const navigate = useNavigate();
    const [player, setPlayer] = useRecoilState(PlayerState);
  
    // Load token from local storage and set into recoil state
    const localStorageToken = localStorage.getItem("player");
    if(localStorageToken && player === undefined) {
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
                <Route path='/' element={<Home/>}/>
                <Route path='/registration' element={<Registration/>}/>
            </Routes>
        </React.Suspense>
    )
}

export default Router;
