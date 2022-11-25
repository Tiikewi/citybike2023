import { useQuery } from "@tanstack/react-query";
import { getPing } from "../lib/apiRequests/pingRequests";
import { message } from "../lib/apiRequests/queryKeys";


export const Ping = (): JSX.Element => {
    const { isError, data, error } = useQuery({
        queryKey: [message],
        queryFn: getPing,
                
    })
    
    if(isError) {
        if (error instanceof Error) {
            return <span>{error.message}</span>
        }
    }

    return(
      <div>
        <p>Data from /ping:</p>
        <p>{data?.data.message}</p>
      </div>
    )
}