import React, { useEffect, useRef } from 'react'
import axios from 'axios';
import { ScoreboardResponse } from '../types/Scoreboard';

const Scoreboard: React.FC = () => {
  const [scoreboardResponse, setScoreboardResponse] = React.useState<ScoreboardResponse>();
  const getScoreboard = useRef(false);

  useEffect(() => {
    if (getScoreboard.current) {
      return;
    }
    getScoreboard.current = true;

    axios
      .get<ScoreboardResponse>("http://localhost:8000/api/latest/scoreboard")
      .then(resp => {
        if (resp.data) {
          setScoreboardResponse(resp.data)
        }
      })
      .catch(ex => {
        alert(ex.message);
      });
  }, [])

  return <div>
    <h1>Scoreboard</h1>

    <h2>Houses</h2>
    <table>
      <thead>
        <tr key={"house-header"}>
          <th>House</th>
          <th>Score</th>
        </tr>
      </thead>
      <tbody>
        {scoreboardResponse?.houses.map((house) =>
          <tr key={house.name}>
            <td>{house.name}</td>
            <td>{house.score}</td>
          </tr>)}
      </tbody>
    </table>

    <h2>Players</h2>
    <table>
      <thead>
        <tr key={"player-header"}>
          <th>House</th>
          <th>Player</th>
          <th>Score</th>
        </tr>
      </thead>
      <tbody>
        {scoreboardResponse?.players.map((player, idx) =>
          <tr key={player.name + '-' + idx}>
            <td>{player.house}</td>
            <td>{player.name}</td>
            <td>{player.score}</td>
          </tr>)}
      </tbody>
    </table>

  </div>
}

export default Scoreboard
