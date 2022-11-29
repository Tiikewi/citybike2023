import { CustomCard } from '../components/CustomCard'
import '../styles/stations.css'

export const Stations = (): JSX.Element => {


    return(
        <div>
            <div className="stations">
                <CustomCard name="Station name 1"></CustomCard>
                <CustomCard name="Station name 2"></CustomCard>
                <CustomCard name="Station name 3"></CustomCard>
            </div>
        </div>
    )
}