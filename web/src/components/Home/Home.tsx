import React from 'react'
import { useRecoilState } from 'recoil';
import PlayerState, { PlayerResponse } from '../../types/Player';
import Scoreboard from '../Scoreboard/Scoreboard';
import './Home.css';
import Header from '../Header/Header';
import Button from 'react-bootstrap/Button';
import useWebSocket from 'react-use-websocket';
import { WebsocketUpdate } from '../../types/Websocket';


const Home: React.FC = () => {
  const [player, setPlayer] = useRecoilState(PlayerState);

  const { sendJsonMessage } = useWebSocket('ws://localhost:8000/ws/player/' + player?.id, {
    onOpen: () => console.log('WebSocket connection opened.'),
    onClose: () => console.log('WebSocket connection closed.'),
    shouldReconnect: (closeEvent) => true,
    onMessage: (event: WebSocketEventMap['message']) => onMessage(event.data),
    onError: (event: WebSocketEventMap['error']) => console.log(event),
  });

  const onMessage = (data: any) => {
    console.log('onMessage: ' + data);
    let updateData = JSON.parse(data) as WebsocketUpdate;

    switch (updateData.type) {
      case 'playerUpdate':
        onPlayerUpdate(JSON.parse(data) as PlayerResponse);
        break;
      default:
        console.log('invalid websocket data: ' + data);
    }
  }

  const onPlayerUpdate = (data: PlayerResponse) => {
    setPlayer(data.player);
    localStorage.setItem("player", JSON.stringify(data.player));
  }

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
