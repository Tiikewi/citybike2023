import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';

type newStation = {
    id: number,
    name: string,
    address: string,
    city: string,
    coordinates: {x: number, y: number},
    capacity?: number,
    operator?: string
}
let newSt: newStation;

export const StationModal = () => {
  const [show, setShow] = useState(false);
  const [station, setStation] = useState<newStation>(newSt);

  const handleClose = () => setShow(false);
  const handleAdd = () => {
    
      setShow(false);
  }
  const handleShow = () => setShow(true);

  return (
    <div>
      <Button variant="primary" onClick={handleShow}>
        Add new station
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Add new station</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="stationInfo">
              <Form.Label>Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Station name"
                autoFocus
              />
              <Form.Label>Address</Form.Label>
              <Form.Control
                type="text"
                placeholder="Station address"
              />
            <Form.Label>ID</Form.Label>
              <Form.Control
                type="number"
                placeholder="Station ID"
                />
            <Form.Label>City</Form.Label>
              <Form.Control
                type="text"
                placeholder="City name"
                />
            <Form.Label>Capacity</Form.Label>
              <Form.Control
                type="number"
                placeholder="Stations capacity"
                />
            <Form.Label>Operator</Form.Label>
              <Form.Control
                type="text"
                placeholder="Operator"
                />
            <Form.Label>Coordinates</Form.Label>
              <Form.Control
                type="number"
                step="0.01"
                placeholder="latitude"
                />
              <Form.Control
                type="number"
                step="0.01"
                placeholder="longitude"
                />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={handleAdd}>
            Add
          </Button>
        </Modal.Footer>
      </Modal>
      </div>
  );
}

