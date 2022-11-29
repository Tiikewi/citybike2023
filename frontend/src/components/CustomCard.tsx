import { useState } from "react";
import { FiChevronDown } from "react-icons/fi";

export const CustomCard = ({
  name,
  address,
  coordinates,
  ...props
}: {
  name: string;
  address: string;
  coordinates: {x: number, y: number}
}) => {

  const [isCollapsed, setIsCollapsed] = useState(true)

const cardClickHandler = () => {
  setIsCollapsed(!isCollapsed)
}

  return (
    <div className="card" onClick={cardClickHandler}>
      <div className="top">
        <h3>{name}</h3>
        <span className="icon">
          <FiChevronDown />
        </span>
      </div>
      <div className="bottom" id={isCollapsed ? 'collapsed' : ''}>
        <p>Address: {address}</p>
        <p>Journeys from this station: </p>
        <p>Journeys to this station: </p>
        <p>Coordinates: ({coordinates.x}, {coordinates.y})</p>
      </div>
    </div>
  );
};