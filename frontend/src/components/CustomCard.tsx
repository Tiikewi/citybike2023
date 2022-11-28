export const CustomCard = ({
  name,
  ...props
}: {
  name: string;
}) => {



  return (
    <div className="card">
      <div className="top">
        <h3>{name}</h3>
      </div>
      <div className="bottom">
        <p>Address</p>
        <p>journeys from</p>
        <p>journey to</p>
      </div>
    </div>
  );
};