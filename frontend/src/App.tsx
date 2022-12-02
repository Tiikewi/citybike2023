import { QueryClient, QueryClientProvider, useIsFetching } from '@tanstack/react-query';
import { Routes, Route } from 'react-router-dom';
import { PageLoadingSpinner } from './components/Spinner';
import { Home } from './views/Home';
import { Navbar } from './views/Navbar';
import { Journeys } from './views/Journeys';
import { Stations } from './views/Stations';

// Create a client
const queryClient = new QueryClient()

const Content = (): JSX.Element => {
  // Access the client
  // const queryClient = useQueryClient()

  const isFetching = useIsFetching();

  return(
    <main>
      {isFetching > 0 ? <PageLoadingSpinner /> : null}
      <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/journeys" element={<Journeys />} />
          <Route path="/stations" element={<Stations />} />
        </Routes>
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