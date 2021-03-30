import Link from '@material-ui/core/Link'
import Typography from '@material-ui/core/Typography'

export default function Copyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <a color="inherit" href="https:/ost.ch/">
                imgprocessing - DS1 - FH OST
            </a>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    )
}
