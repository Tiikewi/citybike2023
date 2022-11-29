import { useState } from "react";
import { FiChevronDown } from "react-icons/fi";

export const CustomCard = ({
  name,
  ...props
}: {
  name: string;
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
        <p>Address: Osoite 123</p>
        <p>Journeys from this station: 122</p>
        <p>Journeys to this station: 344</p>
      </div>
    </div>
  );
};