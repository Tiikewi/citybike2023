import { QueryClient, QueryClientProvider, useIsFetching, useQueryClient } from '@tanstack/react-query';
import { Routes, Route } from 'react-router-dom';
import { Home } from './views/Home';
import { Navbar } from './views/Navbar';
import { Journeys } from './views/Journeys';
import { Stations } from './views/Stations';

// Create a client
const queryClient = new QueryClient()

const Content = (): JSX.Element => {
  // Access the client


  return(
    <main>
      <Navbar />
      <div className="body">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/journeys" element={<Journeys />} />
          <Route path="/stations" element={<Stations />} />
        </Routes>
      </div>
    </main>
  )
}

const App = (): JSX.Element => {
  return (
    <QueryClientProvider client={queryClient}>
      <Content />
    </QueryClientProvider>
)}




export default App;