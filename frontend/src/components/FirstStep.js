import React, { Fragment } from 'react'
import Grid from '@material-ui/core/Grid'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'

const FirstStep = ({ handleNext, handleChange, values: { name } }) => {
    return (
        <Fragment>
            <Grid container spacing={2} noValidate>
                <Grid item xs={12} sm={12}>
                    <TextField
                        fullWidth
                        label="Name"
                        name="name"
                        placeholder="Name"
                        margin="normal"
                        value={name || ''}
                        onChange={handleChange}
                        required
                    />
                </Grid>
            </Grid>
            <div style={{ display: 'flex', marginTop: 50, justifyContent: 'flex-end' }}>
                <Button variant="contained" color="primary" onClick={handleNext}>
                    Next
                </Button>
            </div>
        </Fragment>
    )
}

export default FirstStep
