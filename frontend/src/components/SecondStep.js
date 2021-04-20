import React, { useState } from 'react'
import Grid from '@material-ui/core/Grid'
import Button from '@material-ui/core/Button'
import CustomChip from './CustomChip'
import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        justifyContent: 'center',
        flexWrap: 'wrap',
        '& > *': {
            margin: theme.spacing(0.5),
        },
    },
}))

const SecondStep = ({ handleNext, handleBack, handleChange, values: { Blabla } }) => {
    const classes = useStyles()
    return (
        <>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <div className={classes.root}>
                        <CustomChip></CustomChip>
                    </div>
                </Grid>
            </Grid>
            <div style={{ display: 'flex', marginTop: 50, justifyContent: 'flex-end' }}>
                <Button variant="contained" color="default" onClick={handleBack} style={{ marginRight: 10 }}>
                    Back
                </Button>
                <Button variant="contained" color="primary" onClick={handleNext}>
                    Next
                </Button>
            </div>
        </>
    )
}

export default SecondStep
