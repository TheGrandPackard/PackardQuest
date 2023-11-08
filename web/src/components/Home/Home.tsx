import React, { useState } from "react";
import { useRecoilState } from "recoil";
import PlayerState, { PlayerResponse } from "../../types/Player";
import Scoreboard from "../Scoreboard/Scoreboard";
import "./Home.css";
import axios from 'axios';
import Header from "../Header/Header";
import useWebSocket from "react-use-websocket";
import { WebsocketUpdate } from "../../types/Websocket";
import Button from "react-bootstrap/Button";
import PlayerInfoModal from "../PlayerInfoModal/PlayerInfoModal";

const Home: React.FC = () => {
  const [player, setPlayer] = useRecoilState(PlayerState);
  const [show, setShow] = useState(false);

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);
  const handleUpdateWandId = () => {
    axios
            .put<PlayerResponse>("http://localhost:8000/api/latest/player/" + player?.id, {wandID:player?.wandID}, {})
  } 
  const handleResetPlayer = () => {
    setPlayer(undefined)
    localStorage.clear();
  }
    
  useWebSocket("ws://localhost:8000/ws/player/" + player?.id, {
    onOpen: () => console.log("WebSocket connection opened."),
    onClose: () => console.log("WebSocket connection closed."),
    shouldReconnect: (closeEvent) => true,
    onMessage: (event: WebSocketEventMap["message"]) => onMessage(event.data),
    onError: (event: WebSocketEventMap["error"]) => console.log(event),
  });

  const onMessage = (data: any) => {
    console.log("onMessage: " + data);
    let updateData = JSON.parse(data) as WebsocketUpdate;

    switch (updateData.type) {
      case "playerUpdate":
        onPlayerUpdate(JSON.parse(data) as PlayerResponse);
        break;
      default:
        console.log("invalid websocket data: " + data);
    }
  };

  const onPlayerUpdate = (data: PlayerResponse) => {
    setPlayer(data.player);
    localStorage.setItem("player", JSON.stringify(data.player));
  };

  const renderBody = () => {
    // once registered, direct the player to the sorting hat
    if (player && player.progress.sortingHat === false) {
      return <div>Proceed to the Great Hall to be sorted into your house.</div>;
    }

    // once sorted, direct the player to the pensieve
    if (player && player.progress.sortingHat === true && player.progress.pensieve === false) {
      return <div className={player.house + "-bg"}>The headmaster wishes to see you.</div>;
    }

    // once the player has interacted with the pensieve, show the scoreboard
    return (
      <div className={player?.house + "-bg"}>
        <>Good luck in your studies, and may the best house win the house cup!</>
        <Scoreboard />
      </div>
    );
  };

  return (
    <div className="container">
      <Header />
      <Button variant="primary" onClick={handleShow}>
        Open Player Profile Modal
      </Button>

      <PlayerInfoModal
        player={player!}
        show={show}
        handleClose={handleClose}
        handleUpdateWandId={handleUpdateWandId}
        handleResetPlayer={handleResetPlayer}
      />
      {renderBody()}
    </div>
  );
};

export default Home;
