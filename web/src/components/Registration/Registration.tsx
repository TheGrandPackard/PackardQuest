import React, { useEffect } from 'react'
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useRecoilState } from 'recoil';
import { PlayerResponse, PlayerState } from '../../types/Player';
import Form from 'react-bootstrap/esm/Form';
import "./Registration.css";

export interface UserData {
    name: string;
    wandId: number;
}

const Registration: React.FC = () => {
    const [userData, setUserData] = React.useState<UserData>({ name: '', wandId: 0 });
    const [formError, setFormError] = React.useState('');
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
        console.log(value);
    }

    function handleSelectChange(event: React.ChangeEvent<HTMLSelectElement>) {
        const {name, value} = event.target;
        setUserData({...userData, [name]: Number(value)})
        console.log(value);
    }



    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (userData.name === '' || userData.wandId === 0) {
            setFormError("Invalid name or wand id")
            return;
        }

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
        <>
            <form className="form-control pt-4" onSubmit={handleSubmit}>
                <div className="form-group mb-3">
                <label>
                    Name:
                    <input type="text" name="name" value={userData.name} onChange={handleInputChange}></input>
                </label>
                </div>
                <div className="form-group mb-3">
                <label>Wand:
                <Form.Select name="wandId" onChange={handleSelectChange} aria-label="Select your wand">
                    <option>Select your wand</option>
                    <option value="403796">Pearl</option>
                    <option value="506728">Ruby</option>
                    <option value="1633612">Wolf</option>
                    <option value="1705713">Leaf</option>
                </Form.Select>
                </label>
                {formError !== '' && (<label>{formError}</label>)}
                </div>
                
                <button className="btn btn-primary submit" type="submit">Submit</button>
            </form>

            <div className="symbol">
            <div className="deathly"></div>
            <div className="hallows"></div>
            </div>
        </>
    );

}


export default Registration;