import React, { useEffect, useState } from 'react'
import FAB from './SpeedDial'
import useToken from './useToken'
import FileCopyIcon from '@material-ui/icons/FileCopy'
import IconButton from '@material-ui/core/IconButton'
import { makeStyles } from '@material-ui/core/styles'
import GridList from '@material-ui/core/GridList'
import GridListTile from '@material-ui/core/GridListTile'
import GridListTileBar from '@material-ui/core/GridListTileBar'
import ListSubheader from '@material-ui/core/ListSubheader'
import Hidden from '@material-ui/core/Hidden'
import Grid from '@material-ui/core/Grid'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'center',
        overflow: 'hidden',
    },
    gridList: {
        width: 1000,
    },
    icon: {
        color: 'rgba(255, 255, 255, 0.54)',
    },
}))

export default function Dashboard() {
    const classes = useStyles()
    const { token } = useToken()
    const [error, setError] = useState(null)
    const [isLoading, setIsLoading] = useState(false)
    const [items, setItems] = useState([])
    const [process, setProcess] = useState([])
    useEffect(() => {
        fetch(`https://${window._env_.API_URL}/api/auth/getAllImages`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
        })
            .then((res) => res.json())
            .then(
                (result) => {
                    setIsLoading(false)
                    setItems(result)
                },
                (error) => {
                    setIsLoading(false)
                    setError(error)
                }
            )
    }, [token])
    useEffect(() => {
        fetch(`https://${window._env_.API_URL}/api/auth/getAllProcess`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
        })
            .then((res) => res.json())
            .then(
                (result) => {
                    setIsLoading(false)
                    setProcess(result)
                },
                (error) => {
                    setIsLoading(false)
                    setError(error)
                }
            )
    }, [token])

    return (
        <div className={classes.root}>
            <Grid className={classes.root} item xs={12} sm={12}>
                <GridList cellHeight={220} className={classes.gridList}>
                    <GridListTile key="Subheader" cols={2} style={{ height: 'auto' }}>
                        <ListSubheader component="div">Images</ListSubheader>
                    </GridListTile>
                    <Hidden>
                        {isLoading ||
                            (error && (
                                <GridListTile key="1">
                                    <img src="https://pascalchristen.ch/images/thumbs/7.jpg" alt="Placeholder" />
                                    <GridListTileBar
                                        title="No Images for you 😢 ...but here's one"
                                        subtitle="Start uploading a new image"
                                        actionIcon={
                                            <IconButton
                                                aria-label={`Copy to clipboard`}
                                                className={classes.icon}
                                                onClick={() => {
                                                    navigator.clipboard.writeText('Upload your own image first')
                                                }}
                                            >
                                                <FileCopyIcon />
                                            </IconButton>
                                        }
                                    />
                                </GridListTile>
                            ))}
                    </Hidden>
                    {items.map((tile) => (
                        <GridListTile key={tile.ImageId}>
                            {/*TODO: Use real API path*/}
                            <img src={tile.ImageId} alt={tile.Name} />
                            <GridListTileBar
                                title={tile.ImageId}
                                subtitle={tile.Name}
                                actionIcon={
                                    <IconButton
                                        aria-label={`Copy to clipboard ${tile.ImageId}`}
                                        className={classes.icon}
                                        onClick={() => {
                                            navigator.clipboard.writeText(tile.ImageId)
                                        }}
                                    >
                                        <FileCopyIcon />
                                    </IconButton>
                                }
                            />
                        </GridListTile>
                    ))}
                </GridList>
            </Grid>
            <Grid className={classes.root} item xs={12} sm={12}>
                <GridList cellHeight={220} className={classes.gridList}>
                    <GridListTile key="Subheader" cols={2} style={{ height: 'auto' }}>
                        <ListSubheader component="div">Processes</ListSubheader>
                    </GridListTile>
                    <Hidden>
                        {isLoading ||
                            (error && (
                                <GridListTile key="1">
                                    <img src="https://pascalchristen.ch/images/thumbs/6.jpg" alt="Placeholder" />
                                    <GridListTileBar
                                        title="No Process for you 😢 ...but here's one"
                                        subtitle="Start creating your process"
                                        actionIcon={
                                            <IconButton
                                                aria-label={`Copy to clipboard`}
                                                className={classes.icon}
                                                onClick={() => {
                                                    navigator.clipboard.writeText('Create your process first')
                                                }}
                                            >
                                                <FileCopyIcon />
                                            </IconButton>
                                        }
                                    />
                                </GridListTile>
                            ))}
                    </Hidden>
                    {process.map((tile) => (
                        <GridListTile key={tile.ImageId}>
                            {/*TODO: Use real API path*/}
                            <img src={tile.ImageId} alt={tile.Name} />
                            <GridListTileBar
                                title={tile.ImageId}
                                subtitle={tile.Name}
                                actionIcon={
                                    <IconButton
                                        aria-label={`Copy to clipboard ${tile.ImageId}`}
                                        className={classes.icon}
                                        onClick={() => {
                                            navigator.clipboard.writeText(tile.ImageId)
                                        }}
                                    >
                                        <FileCopyIcon />
                                    </IconButton>
                                }
                            />
                        </GridListTile>
                    ))}
                </GridList>
            </Grid>
            <FAB></FAB>
        </div>
    )
}
