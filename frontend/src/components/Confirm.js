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

    const handleSubmit = () => {
        var createPayload = {
            name,
            processingSteps: [],
        };

        selected.map(function (item) {
            var parameterObj = {}
            item['Options'].map(function (options, index) {
                parameterObj[options] = item.optionsFilled[index]
            })
            var parameterJson = JSON.stringify(parameterObj)
            createPayload.processingSteps.push({
                processingStepType: item.Id,
                executionPosition: parseInt(item.sequence),
                parameterJson: parameterJson,
            })
        });

        return fetch(`https://${window._env_.API_URL}/api/auth/process`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
            body: JSON.stringify(createPayload),
        }).then((response) => {
            if (!response.ok) {
                throw new Error('Something went wrong')
            }
        })
        .then(() => {
            handleNext()
        })
        .catch((error) => {
            setError(error)
        });
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
