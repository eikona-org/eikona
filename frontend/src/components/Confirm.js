import React, { Fragment } from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import Divider from '@material-ui/core/Divider'
import Button from '@material-ui/core/Button'

const Confirm = ({ handleNext, handleBack, values, selected }) => {
    const { name } = values

    const handleSubmit = () => {
        //TODO Push to API after Endpoints are done...
        console.log(name)
        console.log(selected)
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
