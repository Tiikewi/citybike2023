import { useQuery } from '@tanstack/react-query'
import { useState } from 'react'
import { Button, Form } from 'react-bootstrap'
import { CustomCard } from '../components/CustomCard'
import { STATIONS_QUERY_KEY } from '../lib/apiRequests/queryKeys'
import { getStations } from '../lib/apiRequests/stationRequest'
import '../styles/stations.css'

export const Stations = (): JSX.Element => {
    const [searhField, setSearchField] = useState("")
    const [page, setPage] = useState(1)

    const { isError, data, error} = useQuery({
        queryKey: [STATIONS_QUERY_KEY, page],
        queryFn: () => getStations(page),
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
        console.log("search value:", searhField)
        setSearchField("")
    }
    return(
        <div className='stBody'>
            <Form>
                <Form.Group className="mb-3">
                    <Form.Control onChange={onInputChange} type="text" placeholder="Station name" value={searhField} />
                    <Form.Text className="text-muted">
                        Search station by name.
                    </Form.Text>
                </Form.Group>
            <Button variant="primary" onClick={onSearch}>Search</Button>
            </Form>

            <div className="stations">
                {data?.data.map(st =>  (
                <CustomCard name={st.name} address={st.address} coordinates={st.coordinates}></CustomCard>
                ))}
            </div>
        </div>
    )
}