import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import SpeedDial from '@material-ui/lab/SpeedDial'
import SpeedDialIcon from '@material-ui/lab/SpeedDialIcon'
import SpeedDialAction from '@material-ui/lab/SpeedDialAction'
import AddAPhotoIcon from '@material-ui/icons/AddAPhoto'
import PostAddIcon from '@material-ui/icons/PostAdd'
import { Link } from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
    speedDial: {
        position: 'fixed',
        bottom: theme.spacing(8),
        right: theme.spacing(8),
    },
}))

const withLink = (to, children) => (
    <Link style={{ textDecoration: 'none' }} to={to}>
        {children}
    </Link>
)

const actions = [
    { icon: withLink('/upload', <AddAPhotoIcon />), name: 'Upload Image' },
    { icon: withLink('/process', <PostAddIcon />), name: 'Create Process' },
]

export default function FAB() {
    const classes = useStyles()
    const [open, setOpen] = React.useState(false)

    const handleClose = () => {
        setOpen(false)
    }

    const handleOpen = () => {
        setOpen(true)
    }
    return (
        <div className={classes.root}>
            <SpeedDial
                ariaLabel="Eikona Actions"
                className={classes.speedDial}
                icon={<SpeedDialIcon />}
                onClose={handleClose}
                onOpen={handleOpen}
                open={open}
            >
                {actions.map((action) => (
                    <SpeedDialAction
                        key={action.name}
                        icon={action.icon}
                        tooltipTitle={action.name}
                        onClick={handleClose}
                    />
                ))}
            </SpeedDial>
        </div>
    )
}
