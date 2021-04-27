import React, { Fragment, useState } from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import Divider from '@material-ui/core/Divider'
import Button from '@material-ui/core/Button'
import useToken from './useToken'

const Confirm = ({ handleNext, handleBack, values, selected }) => {
    const { name } = values
    const { token } = useToken()
    const [error, setError] = useState(null)
    const [processId, setProcessId] = useState(null)

    const handleSubmit = async () => {
        const process = await fetch(`https://${window._env_.API_URL}/api/auth/process`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
            body: JSON.stringify({ name }),
        }).then((res) => res.json())
        setProcessId(process['ProcessId'])
        console.log(processId)
        console.log(selected)
        const steps = await fetch(`https://${window._env_.API_URL}/api/auth/processsteps`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
            body: JSON.stringify({ processId, selected }),
        }).then((res) => res.json())
        handleNext()
    }

    return (
        <Fragment>
            <List disablePadding>
                <ListItem>
                    <ListItemText primary="Name" secondary={name} />
                </ListItem>

                {selected.map((selectedValues, index) => (
                    <div key={index}>
                        <Divider />
                        <ListItem>
                            {/* //TODO better handling of values and display options */}
                            <ListItemText
                                primary={selectedValues.Name}
                                secondary={
                                    'Position: ' + selectedValues.sequence + ' Options: ' + selectedValues.optionsFilled
                                }
                            />
                        </ListItem>
                    </div>
                ))}

                <Divider />
            </List>

            <div style={{ display: 'flex', marginTop: 50, justifyContent: 'flex-end' }}>
                <Button variant="contained" color="default" onClick={handleBack}>
                    Back
                </Button>
                <Button style={{ marginLeft: 10 }} variant="contained" color="primary" onClick={handleSubmit}>
                    Confirm
                </Button>
            </div>
        </Fragment>
    )
}

export default Confirm
