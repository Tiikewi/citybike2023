import { useQuery} from "@tanstack/react-query";
import { useState } from "react";
import { Button, Table } from "react-bootstrap";
import { GrNext, GrPrevious } from "react-icons/gr";
import { getJourneys,} from "../lib/apiRequests/journeyRequests";
import { JOURNEYS_QUERY_KEY } from "../lib/apiRequests/queryKeys";

const formatDuration = (time: number) => {
  var seconds = time % 60;
  var minutes = (time / 60).toFixed(0);

  return `${minutes} min ${seconds} s`;
};


export const Journeys = (): JSX.Element => {
  const [page, setPage] = useState(1)

    const { isError, data, error} = useQuery({
        queryKey: [JOURNEYS_QUERY_KEY, page],
        queryFn: () => getJourneys(page),
    })

    if(isError) {
      if (error instanceof Error) {
        return <span>{error.message}</span>
      }
    }

return (
    <div className="body">
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


