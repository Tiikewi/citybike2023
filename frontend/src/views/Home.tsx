import '../styles/home.css'

export const Home = (): JSX.Element => {
    return(
      <div className='body'>
        <h1>This is homepage</h1>
          <p>Some useful information here.</p>
          <ul>
            <li>List item 1</li>
            <li>List item 2</li>
            <li>List item 3</li>
          </ul>
        </div>
    )
  }