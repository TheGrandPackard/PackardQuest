import React from 'react'
import { useRecoilState } from 'recoil';
import PlayerState from '../../types/Player';
import Scoreboard from '../Scoreboard/Scoreboard';
import './Home.css';
import Header from '../Header/Header';
import Button from 'react-bootstrap/Button';

const Home: React.FC = () => {
  const [player] = useRecoilState(PlayerState);

  return <div className="container">
    <Header />

    <Button variant="primary" className="mr-1">
      Primary
    </Button>

    {(player && player.progress.sortingHat === false && player.progress.pensieve === false) && (
      <div>
        Proceed to the Great Hall to be sorted into your house.
      </div>
    )}

    {(player && player.progress.pensieve === true) && (
      <div className={player.house + '-bg'}>
        <>Good luck in your studies, and may the best house win the house cup!</>
        <Scoreboard />
      </div>
    )}

    {(player && player.progress.sortingHat === true) && (
      <div className={player.house + '-bg'}>
        The headmaster wishes to see you.
      </div>
    )}

  </div>
}

export default Home
