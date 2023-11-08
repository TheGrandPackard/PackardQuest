import React from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import { Player } from "../../types/Player";

interface PlayerInfoModalProps {
  player: Player;
  show: boolean;
  handleClose: () => void;
  handleUpdateWandId: () => void;
  handleResetPlayer: () => void;
}

const PlayerInfoModal: React.FC<PlayerInfoModalProps> = (props) => {
  return (
    <Modal show={props.show} onHide={props.handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>Player Profile</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>Player Name</Form.Label>
            <Form.Control type="text" disabled={true} value={props.player?.name} />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formWandId">
            <Form.Label>Wand ID</Form.Label>
            <Form.Select name="wandId" aria-label="Select your wand">
                    <option>Select your wand</option>
                    <option value="403796">Pearl</option>
                    <option value="506728">Ruby</option>
                    <option value="1633612">Wolf</option>
                    <option value="1705713">Leaf</option>
                </Form.Select>
          </Form.Group>
          <Button variant="primary" onClick={props.handleUpdateWandId}>
            Submit
          </Button>
          <Button variant="danger" onClick={props.handleResetPlayer}>
            Reset
          </Button>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={props.handleClose}>
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default PlayerInfoModal;
