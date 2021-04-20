import React, { useCallback, useMemo } from 'react'
import useToken from './useToken'
import { useDropzone } from 'react-dropzone'
import Avatar from '@material-ui/core/Avatar'
import CssBaseline from '@material-ui/core/CssBaseline'
import BackupIcon from '@material-ui/icons/Backup'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'

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

const baseStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    padding: '20px',
    borderWidth: 2,
    borderRadius: 2,
    borderColor: '#eeeeee',
    borderStyle: 'dashed',
    backgroundColor: '#fafafa',
    color: '#bdbdbd',
    transition: 'border .3s ease-in-out',
}

const activeStyle = {
    borderColor: '#2196f3',
}

const acceptStyle = {
    borderColor: '#00e676',
}

const rejectStyle = {
    borderColor: '#ff1744',
}

export default function Upload(props) {
    const { token } = useToken()
    const classes = useStyles()
    const onDrop = useCallback(
        (acceptedFiles) => {
            var formData = new FormData()
            formData.append('file', acceptedFiles[0])
            fetch(`https://${window._env_.API_URL}/api/auth/upload`, {
                method: 'POST',
                headers: {
                    Authorization: 'Bearer ' + token,
                },
                body: formData,
            })
        },
        [token]
    )

    const { getRootProps, getInputProps, isDragActive, isDragAccept, isDragReject } = useDropzone({
        onDrop,
        accept: 'image/jpeg, image/png',
    })

    const style = useMemo(
        () => ({
            ...baseStyle,
            ...(isDragActive ? activeStyle : {}),
            ...(isDragAccept ? acceptStyle : {}),
            ...(isDragReject ? rejectStyle : {}),
        }),
        [isDragActive, isDragReject, isDragAccept]
    )

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Avatar className={classes.avatar}>
                    <BackupIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                    Upload
                </Typography>

                <div {...getRootProps({ style })}>
                    <input {...getInputProps()} />
                    <div>Drag and drop your images here.</div>
                </div>
            </div>
        </Container>
    )
}
