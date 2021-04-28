import React, { Fragment, useState } from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import Divider from '@material-ui/core/Divider'
import Button from '@material-ui/core/Button'
import useToken from './useToken'
import Alert from '@material-ui/lab/Alert'
import Hidden from '@material-ui/core/Hidden'

const Confirm = ({ handleNext, handleBack, values, selected }) => {
    const { name } = values
    const { token } = useToken()
    const [error, setError] = useState(null)

    const createProcess = async (process) => {
        const response = await fetch(`https://${window._env_.API_URL}/api/auth/process`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
            body: JSON.stringify(process),
        })
        if (!response.ok) {
            const message = `An error has occured: ${response.status}`;
            throw new Error(message);
        }
    }

    const buildStepsArray = () => {
        var process = {
            name: name,
            processingSteps: [],
        }

        selected.map(function (item) {
            var parameterObj = {}
            item['Options'].map(function (options, index) {
                parameterObj[options] = item.optionsFilled[index]
            })
            process.processingSteps.push({
                processingStepType: item.Id,
                executionPosition: parseInt(item.sequence),
                parameterJson: JSON.stringify(parameterObj),
            })
        })
        return process
    }
    const handleSubmit = () => {
        const process = buildStepsArray()
        createProcess(process).then(function(){
            handleNext()
        }).catch(error => {
            setError(error)
        });
    }

    return (
        <Fragment>
            <Hidden>{error && <Alert severity="error">Upps, something went wrong!</Alert>}</Hidden>
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
