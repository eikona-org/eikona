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
import Divider from '@material-ui/core/Divider'

import Typography from '@material-ui/core/Typography'
import Box from '@material-ui/core/Box'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'center',
        overflow: 'hidden',
        marginBottom: '20px',
    },
    gridList: {
        width: 1000,
    },
    icon: {
        color: 'rgba(255, 255, 255, 0.54)',
    },
    box: {
        padding: theme.spacing(3),
    },
    title: {
        marginTop: 30,
    },
}))

export default function Dashboard() {
    const classes = useStyles()
    const { token } = useToken()
    const [errorProcess, setErrorProcess] = useState(null)
    const [isLoadingProcess, setIsLoadingProcess] = useState(false)
    const [errorImage, setErrorImage] = useState(null)
    const [isLoadingImage, setIsLoadingImage] = useState(false)
    const [items, setItems] = useState([])
    const [process, setProcess] = useState([])
    useEffect(() => {
        fetch(`https://${window._env_.API_URL}/api/auth/images`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
        })
            .then((res) => res.json())
            .then(
                (resultImage) => {
                    setIsLoadingImage(false)
                    setItems(resultImage)
                },
                (errorImage) => {
                    setIsLoadingImage(false)
                    setErrorImage(errorImage)
                }
            )
    }, [token])
    useEffect(() => {
        fetch(`https://${window._env_.API_URL}/api/auth/processes`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
        })
            .then((res) => res.json())
            .then(
                (resultProcess) => {
                    setIsLoadingProcess(false)
                    setProcess(resultProcess)
                },
                (errorProcess) => {
                    setIsLoadingProcess(false)
                    setErrorProcess(errorProcess)
                }
            )
    }, [token])

    return (
        <div className={classes.root}>
            <Grid className={classes.root} item xs={12} sm={12}>
                <GridList cellHeight={220} className={classes.gridList}>
                    <GridListTile key="Subheader" cols={2} style={{ height: 'auto' }}>
                        <ListSubheader component="div">Images</ListSubheader>
                        <Divider />
                    </GridListTile>
                    <Hidden>
                        {isLoadingImage ||
                            (errorImage && (
                                <Box className={classes.box}>
                                    <Typography variant="h2" align="center">
                                        So far you haven't uploaded any pictures 😭
                                    </Typography>
                                </Box>
                            ))}
                    </Hidden>
                    {items.map((tile) => (
                        <GridListTile key={`Tile-${tile.ImageId}`}>
                            {/*TODO: Use real API path*/}
                            <img
                                key={`Img-${tile.ImageId}`}
                                src={
                                    'https://' +
                                    window._env_.API_URL +
                                    '/api/render/dynamic/' +
                                    tile.ImageId +
                                    '?resize-w=600'
                                }
                                alt={tile.Name}
                            />
                            <GridListTileBar
                                title={tile.ImageId}
                                key={`Bar-${tile.ImageId}`}
                                subtitle={tile.Name}
                                actionIcon={
                                    <IconButton
                                        aria-label={`Copy to clipboard ${tile.ImageId}`}
                                        className={classes.icon}
                                        key={`Button-${tile.ImageId}`}
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
                        <Divider />
                    </GridListTile>
                    <Hidden>
                        {isLoadingProcess ||
                            (errorProcess && (
                                <Box className={classes.box}>
                                    <Typography variant="h2" align="center">
                                        So far you haven't created any processes 😭
                                    </Typography>
                                </Box>
                            ))}
                    </Hidden>
                    {process.map((tile) => (
                        <GridListTile key={`Tile-${tile.ProcessId}`}>
                            {/*https://openmoji.org/library/#search=gear&emoji=2699*/}
                            <svg id="emoji" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
                                <g id="color">
                                    <path
                                        fill="#D0CFCE"
                                        stroke="none"
                                        d="M51.2579,29.8111c0,0-1.8149-0.4504-1.0786-2.0742l4.5037-7.4841l-3.6502-3.417l-0.73-0.6833 l-7.1357,4.6903c-0.9431,0.4213-2.037-0.1166-2.2794-1.1205l-1.887-7.8158h-5h-1l-1.8809,7.7858 c-0.2389,0.9894-1.3101,1.5798-2.2425,1.1716c-0.0148-0.0063-0.0296-0.0131-0.0445-0.0198l-7.1357-4.6904l-0.73,0.6833 l-3.6502,3.417l4.5037,7.4841c0.7364,1.6238-1.0785,2.0743-1.0785,2.0743l-8.4766,2.1074l0.1649,4.9972l0.033,0.9995 l8.3622,1.7291c0.952,0.3641,1.3501,1.4954,0.8358,2.3755l-0.0328,0.0562l-4.5037,7.484l3.6502,3.417l0.73,0.6833l7.1357-4.6902 c0.0149-0.0068,0.03-0.0129,0.0445-0.02c0,0,0.0115-0.0042,0.0332-0.0111c1.1606-0.3768,2.393,0.3463,2.6794,1.5325l1.7947,7.4336 h1h5l1.887-7.8158c0.2424-1.0039,1.3364-1.5417,2.2794-1.1205l7.1357,4.6903l0.73-0.6833l3.6502-3.417l-4.5037-7.4841 l-0.0328-0.0562c-0.5143-0.88-0.1162-2.0112,0.8358-2.3755l8.3622-1.729l0.033-0.9995l0.1649-4.9973L51.2579,29.8111z M36,44.9064 c-5.5228,0-10-4.4772-10-10s4.4772-10,10-10s10,4.4772,10,10S41.5228,44.9064,36,44.9064z"
                                    />
                                    <path
                                        fill="#9B9B9A"
                                        stroke="none"
                                        d="M36,24.9064c-5.5228,0-10,4.4772-10,10s4.4772,10,10,10s10-4.4772,10-10S41.5228,24.9064,36,24.9064z M36,39.9064c-2.7614,0-5-2.2386-5-5c0-2.7615,2.2386-5,5-5s5,2.2385,5,5C41,37.6678,38.7614,39.9064,36,39.9064z"
                                    />
                                    <path
                                        fill="none"
                                        stroke="#FFFFFF"
                                        stroke-miterlimit="10"
                                        stroke-width="2"
                                        d="M25.9551,16.18"
                                    />
                                    <path
                                        fill="none"
                                        stroke="#FFFFFF"
                                        stroke-miterlimit="10"
                                        stroke-width="2"
                                        d="M44.1915,16.0582"
                                    />
                                </g>
                                <g id="hair" />
                                <g id="skin" />
                                <g id="skin-shadow" />
                                <g id="line">
                                    <circle
                                        cx="36"
                                        cy="34.9064"
                                        r="5"
                                        fill="none"
                                        stroke="#000000"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-miterlimit="10"
                                        stroke-width="2"
                                    />
                                    <circle
                                        cx="36"
                                        cy="34.9064"
                                        r="10"
                                        fill="none"
                                        stroke="#000000"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-miterlimit="10"
                                        stroke-width="2"
                                    />
                                    <path
                                        fill="none"
                                        stroke="#000000"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-miterlimit="10"
                                        stroke-width="2"
                                        d="M31.1191,19.6921c-0.2389,0.9895-1.31,1.5799-2.2425,1.1717c-0.0147-0.0065-0.0296-0.0131-0.0445-0.0199l-7.1357-4.6903 l-0.73,0.6834l-3.6502,3.417l4.5037,7.484c0.7363,1.6239-1.0785,2.0744-1.0785,2.0744l-8.4766,2.1074l0.1649,4.9973l0.033,0.9995 l8.3622,1.7291h0c0.9521,0.3641,1.3501,1.4954,0.8359,2.3754l-0.0328,0.0562l-4.5037,7.484l3.6502,3.417l0.73,0.6834 l7.1357-4.6903c0.0149-0.0068,0.03-0.0129,0.0445-0.0199c0,0,0.0116-0.0042,0.0332-0.0112 c1.1606-0.3768,2.393,0.3463,2.6794,1.5324l1.7947,7.4336h1h5l1.8869-7.8157c0.2424-1.004,1.3364-1.5418,2.2794-1.1205h0 l7.1357,4.6903l0.73-0.6834l3.6502-3.417l-4.5037-7.484l-0.0328-0.0562c-0.5143-0.8801-0.1162-2.0113,0.8359-2.3755l0,0 l8.3622-1.7291l0.033-0.9995l0.1649-4.9973l-8.4766-2.1074c0,0-1.8149-0.4505-1.0785-2.0744l4.5037-7.484l-3.6502-3.417 l-0.73-0.6834l-7.1357,4.6903v0c-0.943,0.4212-2.037-0.1166-2.2794-1.1206L39,11.9064h-5h-1L31.1191,19.6921"
                                    />
                                </g>
                            </svg>
                            <GridListTileBar
                                title={tile.ProcessId}
                                key={`Bar-${tile.ProcessId}`}
                                subtitle={tile.Name}
                                actionIcon={
                                    <IconButton
                                        aria-label={`Copy to clipboard ${tile.ProcessId}`}
                                        className={classes.icon}
                                        key={`Button-${tile.ProcessId}`}
                                        onClick={() => {
                                            navigator.clipboard.writeText(tile.ProcessId)
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
