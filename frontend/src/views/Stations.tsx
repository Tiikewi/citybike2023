import { useQuery } from '@tanstack/react-query'
import { useState } from 'react'
import { Button, Form } from 'react-bootstrap'
import { CustomCard } from '../components/StationCard'
import { STATIONS_QUERY_KEY } from '../lib/apiRequests/queryKeys'
import { getStations } from '../lib/apiRequests/stationRequest'
import '../styles/stations.css'
import { GrNext, GrPrevious } from "react-icons/gr";


export const Stations = (): JSX.Element => {
    const [searhField, setSearchField] = useState('')
    const [page, setPage] = useState(1)


    
    const { isError, data, error, refetch} = useQuery({
        queryKey: [STATIONS_QUERY_KEY, page],
        queryFn: () => getStations(page, searhField),
    })

    if(isError) {
      if (error instanceof Error) {
        return <span>{error.message}</span>
      }
    }


    const onInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchField(e.target.value)
    }

    const onSearch = () => {
        setPage(1)
        refetch()
        setSearchField('')
    }

    const handleKeyDown = (e: React.KeyboardEvent<HTMLFormElement>) => {
        if (e.key === 'Enter') {
            onSearch()
        }
      }

    return(
        <div className='body'>
            <Form onSubmit={(e) => e.preventDefault()}  onKeyDown={handleKeyDown}>
                <Form.Group className="mb-3">
                    <Form.Control onChange={onInputChange} type="text" placeholder="Station name" value={searhField} />
                    <Form.Text className="text-muted">
                        Search station by name. 
                    </Form.Text>
                </Form.Group>
            <Button variant="primary" onClick={onSearch}>Search</Button>
            </Form>

            <div className="page">
                <Button variant='white' onClick={() => 
                    {if(page > 1) setPage(page - 1)}}>
                        <GrPrevious /></Button>
                <label>Page: {page}</label>
                {/* TODO prevent page going over avaible pages. */}
                <Button onClick={() => setPage(page + 1)} variant='white'><GrNext /></Button>
            </div>

            <div className="stations">
               {data?.data === null ? (<p>No stations</p>) : (
                data?.data.map(st =>  (
                <CustomCard station={st}></CustomCard>
                )))}
            </div>
        </div>
    )
}