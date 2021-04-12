import React from 'react'
import { useNavigate } from 'react-router-dom'
import { makeStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import Hidden from '@material-ui/core/Hidden'
import { Favorite, GitHub } from '@material-ui/icons'
import ImageSearchIcon from '@material-ui/icons/ImageSearch';

const useStyles = makeStyles((theme) => ({
    root: {
        flexWrap: 'wrap',
    },
    toolbar: {
        flexWrap: 'wrap',
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
        textTransform: 'none',
    },
}))

export default function HeaderBar() {
    const classes = useStyles()
    let navigate = useNavigate()
    return (
        <React.Fragment>
            <AppBar position="relative" className={classes.root}>
                <Toolbar className={classes.toolbar}>
                    <Grid container direction="row" justify="space-between" alignItems="center">
                        <Grid item xs={12} md={4}>
                            <Button onClick={() => navigate('/')} aria-label="Image Processing'" color="inherit">
                                <Typography variant="h6" className={classes.title}>
                                    Image Processing
                                </Typography>
                            </Button>
                        </Grid>
                        <Grid item xs={12} md={8}>
                            <Hidden only={['md', 'sm', 'xs']}>
                                <Grid container direction="row" justify="flex-end" alignItems="center">
                                    <Grid item>
                                        <IconButton
                                            onClick={() => navigate('/dashboard')}
                                            aria-label="Test"
                                            color="inherit"
                                        >
                                            <ImageSearchIcon fontSize="small" />
                                        </IconButton>
                                        <IconButton
                                            onClick={() => window.open('https://github.com/imgProcessing', '_blank')}
                                            aria-label="Source Code"
                                            color="inherit"
                                        >
                                            <GitHub fontSize="small" />
                                        </IconButton>
                                        <IconButton
                                            onClick={() => window.open('https://ost.ch/', '_blank')}
                                            aria-label="Support"
                                            color="secondary"
                                        >
                                            <Favorite fontSize="small" />
                                        </IconButton>
                                    </Grid>
                                </Grid>
                            </Hidden>
                            <Hidden only={['xl', 'lg']}>
                                <Grid container direction="row" justify="flex-start" alignItems="center">
                                    <Grid item>
                                        <IconButton
                                            onClick={() => navigate('/dashboard')}
                                            aria-label="Test"
                                            color="inherit"
                                        >
                                            <ImageSearchIcon fontSize="small" />
                                        </IconButton>
                                        <IconButton
                                            onClick={() => window.open('https://github.com/imgProcessing', '_blank')}
                                            aria-label="Source Code"
                                            color="inherit"
                                        >
                                            <GitHub fontSize="small" />
                                        </IconButton>
                                        <IconButton
                                            onClick={() => window.open('https://ost.ch/', '_blank')}
                                            aria-label="Support"
                                            color="secondary"
                                        >
                                            <Favorite fontSize="small" />
                                        </IconButton>
                                    </Grid>
                                </Grid>
                            </Hidden>
                        </Grid>
                    </Grid>
                </Toolbar>
            </AppBar>
        </React.Fragment>
    )
}
