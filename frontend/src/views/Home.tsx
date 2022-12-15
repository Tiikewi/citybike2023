import '../styles/home.css'


export const Home = (): JSX.Element => {

  

    return(
      <div className='body'>
        <h1>Citybike 2023</h1>
        <p>This is the pre-assignment for Solita Dev Academy 2023.</p>
          <p>Citybike is app for displaying data of journeys made with city bikes in Helsinki Capital area.</p>
          <p>Going to implement some improvements during application period, depending amount of time I have available.</p>
          <br></br>

          <h3>TODOS</h3>
          <ul>
            <li>Really improve and implement tests</li>
            <li>Make UI responsive for various screen sizes</li>
            <li>UI element for creating journeys/stations</li>
            <li>Improve CSV reading and validation and perhaps add ability to add csv files to db from client.</li>
          </ul>

        </div>
    )
  }

  