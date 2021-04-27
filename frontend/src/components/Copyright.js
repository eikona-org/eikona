import Typography from '@material-ui/core/Typography'

export default function Copyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'eikona - DS1, OST - '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    )
}
