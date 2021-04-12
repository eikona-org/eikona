import React, { useEffect, useState } from 'react'
import FAB from './SpeedDial'
import useToken from './useToken'
import { makeStyles } from '@material-ui/core/styles'
import GridList from '@material-ui/core/GridList'
import GridListTile from '@material-ui/core/GridListTile'
import GridListTileBar from '@material-ui/core/GridListTileBar'
import ListSubheader from '@material-ui/core/ListSubheader'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'center',
        overflow: 'hidden',
    },
    gridList: {
        width: 800,
    },
    icon: {
        color: 'rgba(255, 255, 255, 0.54)',
    },
}))

export default function Dashboard() {
    const classes = useStyles()
    const { token } = useToken()
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [items, setItems] = useState([])
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
                    setIsLoaded(true)
                    setItems(result)
                },
                (error) => {
                    setIsLoaded(true)
                    setError(error)
                }
            )
    }, [token])

    if (error) {
        return <div>Error: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Loading...</div>
    } else {
        return (
            <div className={classes.root}>
                <GridList cellHeight={220} className={classes.gridList}>
                    <GridListTile key="Subheader" cols={2} style={{ height: 'auto' }}>
                        <ListSubheader component="div">Uploaded Images</ListSubheader>
                    </GridListTile>
                    {items.map((tile) => (
                        <GridListTile key={tile.img}>
                            <img src={tile.img} alt={tile.id} />
                            <GridListTileBar title={tile.id} subtitle={tile.name} />
                        </GridListTile>
                    ))}
                </GridList>
                <FAB></FAB>
            </div>
        )
    }
}
