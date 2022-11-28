import { useQuery} from "@tanstack/react-query";
import { useState } from "react";
import { getJourneys,} from "../lib/apiRequests/journeyRequests";
import { JOURNEYS_QUERY_KEY } from "../lib/apiRequests/queryKeys";



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
    <div className="container">
      <table>
        <thead>
          <tr>
            <th>Dep id</th>
            <th>Dep name</th>
            <th>Ret id</th>
            <th>Ret name</th>
            <th>Distance</th>
            <th>Duration</th>
          </tr>
        </thead>
        <tbody>
          {data?.data.map(journey => (
            <tr key={journey.id}>
              <td>{journey.departureStationId}</td>
              <td>{journey.departureStationName}</td>
              <td>{journey.returnStationId}</td>
              <td>{journey.returnStationName}</td>
              <td>{journey.distance}</td>
              <td>{journey.duration}</td>
            </tr>
        ))}
        </tbody>
      </table>

      <button onClick={() => {if(page > 1) setPage(page - 1)}}>Previous</button>
      <button onClick={() => setPage(page + 1)}>Next</button>
    </div>
  )
}


