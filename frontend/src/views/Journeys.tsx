import { useQuery} from "@tanstack/react-query";
import { useState } from "react";
import { Button, ButtonGroup, Table } from "react-bootstrap";
import { GrNext, GrPrevious } from "react-icons/gr";
import { PageLoadingSpinner } from "../components/Spinner";
import { getJourneys,} from "../lib/apiRequests/journeyRequests";
import "../styles/journeys.css"

const formatDuration = (time: number) => {
  var seconds = time % 60;
  var minutes = (time / 60).toFixed(0);

  return `${minutes} min ${seconds} s`;
};

export enum Sort {
  DepartureName = 1,
  DepartureID = 2,
  ReturnName = 3,
  ReturnID = 4,
  Distance = 5,
  Duration = 6
}

export const Journeys = (): JSX.Element => {
  const [page, setPage] = useState(1)
  const [sort, setSort] = useState(Sort.DepartureID)


    const { isError, data, error, refetch, isFetching} = useQuery({
        queryKey: [page, sort],
        queryFn: () => getJourneys(page, sort),
    })

    if(isError) {
      if (error instanceof Error) {
        return <span>{error.message}</span>
      }
    }
    if(isFetching) {
      return   (<PageLoadingSpinner />)
    }

    const onSort = (sortValue: Sort) => {
      setSort(sortValue)
      refetch()
  }

return (
  <div className="body">
           <div className="sort">
            <label>Sort by: <span className="sortItem">{Sort[sort]}</span></label>
            <br />
            <ButtonGroup>
              <Button variant={sort === Sort.DepartureID ? 'success' : 'secondary'}
                  onClick={() => onSort(Sort.DepartureID)}>Dep. ID</Button>
                <Button variant={sort === Sort.DepartureName ? 'success' : 'secondary'}  
                    onClick={() => onSort(Sort.DepartureName)}>Dep. Name</Button>
                <Button variant={sort === Sort.ReturnName ? 'success' : 'secondary'} 
                    onClick={() => onSort(Sort.ReturnName)}>Ret. Name</Button>
                <Button variant={sort === Sort.ReturnID ? 'success' : 'secondary'}
                    onClick={() => onSort(Sort.ReturnID)}>Ret. ID</Button>
                <Button variant={sort === Sort.Distance ? 'success' : 'secondary'} 
                    onClick={() => onSort(Sort.Distance)}>Distance</Button>
                <Button variant={sort === Sort.Duration ? 'success' : 'secondary'}
                  onClick={() => onSort(Sort.Duration)}>Duration</Button>
            </ButtonGroup>
        </div>
      <Table striped>
        <thead>
          <tr>
            <th colSpan={2}>Departure</th>
            <th>|</th>
            <th colSpan={2}>Return</th>
            <th>|</th>
          </tr>
          <tr>
            <th style={{width: "10%"}}>Station ID</th>
            <th style={{width: "20%"}}>Station name</th>
            <th></th>
            <th style={{width: "10%"}}>Station ID</th>
            <th style={{width: "20%"}}>Station name</th>
            <th></th>
            <th>Distance (km)</th>
            <th>Duration</th>
          </tr>
        </thead>
        <tbody>
          {data?.data.map(journey => (
            <tr key={journey.id}>
              <td>{journey.departureStationId}</td>
              <td>{journey.departureStationName}</td>
              <td>|</td>
              <td>{journey.returnStationId}</td>
              <td>{journey.returnStationName}</td>
              <td>|</td>
              <td>{(journey.distance / 100).toFixed(2)}</td>
              <td>{formatDuration(journey.duration)}</td>
            </tr>
        ))}
        </tbody>
      </Table>
      <div className="page">
        <Button variant='white' onClick={() => 
                    {if(page > 1) setPage(page - 1)}}>
                        <GrPrevious />
        </Button>
        <label>Page: {page}</label>
        {/* TODO prevent page going over avaible pages. */}
        <Button onClick={() => setPage(page + 1)} variant='white'><GrNext /></Button>
      </div>
    </div>
  )
}


