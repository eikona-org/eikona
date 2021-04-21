import React, { useState } from 'react'
import Box from '@material-ui/core/Box'
import Typography from '@material-ui/core/Typography'
import Stepper from '@material-ui/core/Stepper'
import Step from '@material-ui/core/Step'
import StepLabel from '@material-ui/core/StepLabel'
import FirstStep from './FirstStep'
import SecondStep from './SecondStep'
import Confirm from './Confirm'
import Success from './Success'
import Paper from '@material-ui/core/Paper'
import { makeStyles } from '@material-ui/core/styles'

const labels = ['Set Name', 'Add Actions', 'Confirmation']

const initialValues = {
    name: '',
}
const useStyles = makeStyles((theme) => ({
    layout: {
        width: 'auto',
        marginLeft: theme.spacing(2),
        marginRight: theme.spacing(2),
        [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
            width: 600,
            marginLeft: 'auto',
            marginRight: 'auto',
        },
    },
    paper: {
        marginTop: theme.spacing(3),
        marginBottom: theme.spacing(3),
        padding: theme.spacing(2),
        [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
            marginTop: theme.spacing(8),
            marginBottom: theme.spacing(8),
            padding: theme.spacing(3),
        },
    },
}))

const StepForm = () => {
    const classes = useStyles()
    const [activeStep, setActiveStep] = useState(0)
    const [formValues, setFormValues] = useState(initialValues)
    const [selected, setSelected] = useState([])

    const handleNext = () => setActiveStep((prev) => prev + 1)
    const handleBack = () => setActiveStep((prev) => prev - 1)

    const handleChange = (e) => {
        const { name, value } = e.target

        setFormValues((prev) => ({
            ...prev,
            [name]: value,
        }))
    }

    const handleSteps = (step) => {
        switch (step) {
            case 0:
                return <FirstStep handleNext={handleNext} handleChange={handleChange} values={formValues} />
            case 1:
                return (
                    <SecondStep
                        handleNext={handleNext}
                        handleBack={handleBack}
                        handleChange={handleChange}
                        updateSelected={(s) => setSelected(s)}
                    />
                )
            case 2:
                return (
                    <Confirm handleNext={handleNext} handleBack={handleBack} values={formValues} selected={selected} />
                )
            default:
                break
        }
    }

    return (
        <>
            <main className={classes.layout}>
                <Paper className={classes.paper}>
                    {activeStep === labels.length ? (
                        // Last Component
                        <Success values={formValues} />
                    ) : (
                        <>
                            <Box style={{ margin: '30px 0 50px' }}>
                                <Typography variant="h4" align="center">
                                    Create your process
                                </Typography>
                            </Box>
                            <Stepper activeStep={activeStep} style={{ margin: '30px 0 15px' }} alternativeLabel>
                                {labels.map((label) => (
                                    <Step key={label}>
                                        <StepLabel>{label}</StepLabel>
                                    </Step>
                                ))}
                            </Stepper>
                            {handleSteps(activeStep)}
                        </>
                    )}
                </Paper>
            </main>
        </>
    )
}

export default StepForm
