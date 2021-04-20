import React, { Fragment } from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import Divider from '@material-ui/core/Divider'
import Button from '@material-ui/core/Button'

const Confirm = ({ handleNext, handleBack, values }) => {
    const { name, process } = values

    const handleSubmit = () => {
        handleNext()
    }

    return (
        <Fragment>
            <List disablePadding>
                <ListItem>
                    <ListItemText primary="Name" secondary={name} />
                </ListItem>

                <Divider />

                <ListItem>
                    <ListItemText primary="Process Steps" secondary={process} />
                </ListItem>

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
