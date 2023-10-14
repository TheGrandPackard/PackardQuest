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
            <Form.Control type="number" placeholder="1000" />
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
