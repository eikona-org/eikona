import React, { useEffect, useState } from 'react'
import Chip from '@material-ui/core/Chip'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogTitle from '@material-ui/core/DialogTitle'
import useToken from './useToken'

const CustomChip = ({ updateSelected }) => {
    const { token } = useToken()
    const [data, setData] = useState([])
    const [selected, setSelected] = useState([])
    const [errorSteps, setErrorSteps] = useState(null)

    useEffect(() => {
        fetch(`https://${window._env_.API_URL}/api/auth/processingsteptypes`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + token,
            },
        })
            .then((res) => res.json())
            .then(
                (resultSteps) => {
                    Object.keys(resultSteps).map((e, i) => {
                        const length = resultSteps[i]['Options'].length
                        resultSteps[i].sequence = ''
                        resultSteps[i].color = 'default'
                        resultSteps[i].selected = false
                        resultSteps[i].isOpen = false
                        resultSteps[i].optionsFilled = Array(length).fill('')
                    })
                    setData(resultSteps)
                },
                (errorSteps) => {
                    setErrorSteps(errorSteps)
                }
            )
    }, [token])

    const handleChange = (e, dataindex, index) => {
        updateItemArray(dataindex, 'optionsFilled', index, e.target.value)
    }

    const handleChangeSequence = (e, dataindex) => {
        updateItem(dataindex, 'sequence', e.target.value)
    }

    const handleDialogCancel = (e, dataindex) => {
        updateItem(dataindex, 'isOpen', false)
    }

    const updateItemArray = (index, whichvalue, optionindex, newvalue) => {
        //https://stackoverflow.com/questions/37662708/react-updating-state-when-state-is-an-array-of-objects
        let temporaryarray = data.slice()
        temporaryarray[index][whichvalue][optionindex] = newvalue
        setData(temporaryarray)
    }

    const updateItem = (index, whichvalue, newvalue) => {
        //https://stackoverflow.com/questions/37662708/react-updating-state-when-state-is-an-array-of-objects
        let temporaryarray = data.slice()
        temporaryarray[index][whichvalue] = newvalue
        setData(temporaryarray)
    }

    const handleSubmit = (e, id, index) => {
        e.preventDefault()
        var d = data.find((d) => d.Id === id)
        const selectedHistory = [...selected, d]
        setSelected(selectedHistory)
        updateItem(index, 'color', 'primary')
        updateItem(index, 'selected', true)
        updateItem(index, 'isOpen', false)
        updateSelected(selectedHistory)
    }

    const handleChipClick = (e, index) => {
        updateItem(index, 'isOpen', true)
    }

    const handleChipDelete = (index, id) => {
        updateItem(index, 'color', 'default')
        updateItem(index, 'selected', false)
        const selectedDelete = selected.filter((item) => item.Id !== id)
        setSelected(selectedDelete)
        updateSelected(selectedDelete)
    }

    return (
        <>
            {!errorSteps && (
                <>
                    {data.map((step, index) => (
                        <div key={index}>
                            {step.selected ? (
                                <Chip
                                    label={step.Name}
                                    id={index}
                                    key={index + 'test'}
                                    clickable
                                    onClick={(e) => handleChipClick(e, index)}
                                    color={step.color}
                                    onDelete={(e) => handleChipDelete(index, step.Id)}
                                />
                            ) : (
                                <Chip
                                    label={step.Name}
                                    id={index}
                                    key={index + 'test'}
                                    clickable
                                    onClick={(e) => handleChipClick(e, index)}
                                    color={step.color}
                                />
                            )}
                            <Dialog open={step.isOpen} onClose={handleSubmit} aria-labelledby="form-dialog-title">
                                <DialogTitle id="form-dialog-title">{step.Name}</DialogTitle>
                                <DialogContent>
                                    <TextField
                                        name="Execution position"
                                        onChange={(e) => handleChangeSequence(e, index)}
                                        autoFocus
                                        margin="dense"
                                        label="Execution position"
                                        type="number"
                                        value={step.sequence}
                                        fullWidth
                                        required
                                        key={step.Id}
                                    />
                                    {step.Options.map((option, i) => (
                                        <TextField
                                            margin="dense"
                                            name={option}
                                            onChange={(e) => handleChange(e, index, i)}
                                            key={option}
                                            label={option}
                                            value={step.optionsFilled[i]}
                                            type="text"
                                            fullWidth
                                        />
                                    ))}
                                </DialogContent>
                                <DialogActions>
                                    <Button id={index} onClick={(e) => handleDialogCancel(e, index)}>
                                        Cancel
                                    </Button>
                                    <Button id={index} onClick={(e) => handleSubmit(e, step.Id, index)} color="primary">
                                        OK
                                    </Button>
                                </DialogActions>
                            </Dialog>
                        </div>
                    ))}
                </>
            )}
        </>
    )
}

export default CustomChip
