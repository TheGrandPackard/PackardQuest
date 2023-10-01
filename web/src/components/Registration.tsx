import React, { useEffect } from 'react'
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useRecoilState } from 'recoil';
import { Player, PlayerResponse, PlayerState } from '../types/Player';

export interface UserData {
    name: string;
    wandId: number;
}

const Registration: React.FC = () => {
    const [userData, setUserData] = React.useState<UserData>({ name: '', wandId: 0 });
    const [player, setPlayer] = useRecoilState(PlayerState);
    const navigate = useNavigate();

    // Navigate the user back to home after registering
    useEffect(() => {
        if (player) {
            navigate('/', { replace: true });
        }
    }, [navigate, player])

    function handleInputChange(event: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = event.target;
        setUserData({ ...userData, [name]: value });
    }

    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        axios
            .post<PlayerResponse>("http://localhost:8000/api/latest/player", userData, {})
            .then(resp => {
                if (resp.data.player) {
                    setPlayer(resp.data.player);
                    localStorage.setItem("player", JSON.stringify(resp.data.player));
                }
            })
            .catch(ex => {
                alert(ex.message);
            });
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Name:
                <input type="text" name="name" value={userData.name} onChange={handleInputChange}></input>
            </label>
            <br />
            <button type="submit">Submit</button>
        </form>
    );

}


export default Registration;