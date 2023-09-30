import React from 'react'
import axios from 'axios';

interface FormProps {
    onSubmit?: (data: UserData) => void;
}

interface PlayerResponse {
    player: Player;
}

interface Player {
    id: number;
    name: string;
    wandID: number;
    house: string;
    progress: PlayerProgress;
}

interface PlayerProgress {
    sortingHat: boolean;
    pensieve: boolean;
}

export interface UserData {
    name: string;
}

const RegistrationFormHogwarts: React.FC<FormProps> = ({onSubmit}: FormProps) => {
    const [userData, setUserData] = React.useState<UserData>({name: ''});

function handleInputChange(event: React.ChangeEvent<HTMLInputElement>) {
    const {name, value} = event.target;
    setUserData({...userData, [name]: value});
}


function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    axios
    .post<PlayerResponse>("http://localhost:8000/api/latest/player", userData, {})
    .then(resp => {
      if (resp.data.player) {
          //setToken(resp.player);
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


export default RegistrationFormHogwarts;