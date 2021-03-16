import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import { DropzoneAreaBase } from 'material-ui-dropzone'

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(3),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}))

export default function LogIn({ setToken }) {
    const classes = useStyles()

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <DropzoneAreaBase
                    acceptedFiles={['image/*']}
                    dropzoneText={'Drag and drop an image here or click'}
                    onChange={(files) => console.log('Files:', files)}
                    onAlert={(message, variant) => console.log(`${variant}: ${message}`)}
                />
            </div>
        </Container>
    )
}

LogIn.propTypes = {
    setToken: PropTypes.func.isRequired,
}
