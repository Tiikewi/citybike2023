import { useState } from "react";
import { Table } from "react-bootstrap";
import { FiChevronDown } from "react-icons/fi";
import { Station } from "../lib/apiRequests/stationRequest";
import { Wrapper, Status } from "@googlemaps/react-wrapper";
import {Map} from "./Map"

const render = (status: Status) => {
  return <h1>{status}</h1>;
};

// DO NOT STORE SECRETS ON REACT APP!
// api key I use on tiikewi.fi is restricted to spesific IP.
const GOOGLE_MAPS_API_KEY = process.env.REACT_APP_MAPS_API_KEY || ""

export const StationCard = ({
  station,
}: {
  station: Station;
}) => {
  
  const [isCollapsed, setIsCollapsed] = useState(true)

  const cardClickHandler = () => {
    setIsCollapsed(!isCollapsed)
  }

  return (
    <div className="card">
      <div className="top" onClick={cardClickHandler}>
        <h3>{station.name}</h3>
        <span className="icon">
          <FiChevronDown />
        </span>
      </div>
      <div className="bottom" id={isCollapsed ? 'collapsed' : ''}>      
        <Table striped bordered>
          <thead>
            <tr>
              <th>Address</th>
              <th>Departures</th>
              <th>Returns</th>
              <th>Coordinates</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{station.address}</td>
              <td>{station.departures}</td>
              <td>{station.returns}</td>
              <td>{station.coordinates !== undefined ? <p>({station.coordinates.x}, {station.coordinates.y})</p> : null}</td>
            </tr>
          </tbody>
        </Table>
        <Wrapper apiKey={GOOGLE_MAPS_API_KEY} render={render}>
          <Map lat={station.coordinates.y} lng={station.coordinates.x}/>
        </Wrapper>
      </div>
    </div>
  );
};
