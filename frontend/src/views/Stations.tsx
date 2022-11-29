import { useState } from 'react'
import { Button, Form } from 'react-bootstrap'
import { CustomCard } from '../components/CustomCard'
import '../styles/stations.css'

export const Stations = (): JSX.Element => {
    const [searhField, setSearchField] = useState("")

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
                <CustomCard name="Station name 1"></CustomCard>
                <CustomCard name="Station name 2"></CustomCard>
                <CustomCard name="Station name 3"></CustomCard>
                <CustomCard name="Station name 4"></CustomCard>
                <CustomCard name="Station name 5"></CustomCard>
                <CustomCard name="Station name 6"></CustomCard>
                <CustomCard name="Station name 7"></CustomCard>
                <CustomCard name="Station name 8"></CustomCard>
                <CustomCard name="Station name 9"></CustomCard>
                <CustomCard name="Station name 10"></CustomCard>
       
            </div>
        </div>
    )
}