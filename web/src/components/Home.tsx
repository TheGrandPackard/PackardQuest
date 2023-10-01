import React from 'react'
import { useRecoilState } from 'recoil';
import PlayerState from '../types/Player';
import Scoreboard from './Scoreboard';

const Home:React.FC = () => {
  const [player] = useRecoilState(PlayerState);

  if(player && player.progress.pensieve === true) {
    return <>
      <>Good luck in your studies, and may the best house win the house cup!</>
      <Scoreboard/>
    </>
  }

  if(player && player.progress.sortingHat === true) {
    return <>The headmaster wishes to see you.</>
  }
  
  return <>Proceed to the Great Hall to be sorted into your house.</>
}

export default Home