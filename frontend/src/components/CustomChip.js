import React, { Component } from 'react'
import Chip from '@material-ui/core/Chip'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogTitle from '@material-ui/core/DialogTitle'
import { makeStyles } from '@material-ui/core/styles'

class CustomChip extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: [],
            data: [
                {
                    id: 10,
                    name: 'Resize',
                    options: ['width', 'height'],
                    optionsFilled: ['', ''],
                    sequence: '',
                    isSelected: false,
                    isOpen: false,
                    color: 'default',
                },
                { id: 20, name: 'Grayscale', options: [], isSelected: false, isOpen: false, color: 'default' },
                { id: 30, name: 'Blub', options: [], isSelected: false, isOpen: false, color: 'default' },
            ],
        }
        this.handleChipDelete = this.handleChipDelete.bind(this)
        this.handleChipClick = this.handleChipClick.bind(this)
    }

    callBackMethod() {
        this.props.sendData(this.state.selected)
    }

    handleChange = (e, dataindex, index) => {
        const { data } = this.state
        data[dataindex].optionsFilled[index] = e.target.value
        this.setState({ data })
    }

    handleChangeSequence = (e, dataindex) => {
        const { data } = this.state
        data[dataindex].sequence = e.target.value
        this.setState({ data })
    }

    handleDialogSubmit = (e) => {
        e.preventDefault()
        const { id } = e.currentTarget
        let tmp = this.state.data
        tmp[id].isOpen = !tmp[id].isOpen
        this.setState({ arr: tmp })
        this.state.open = true
    }

    handleDialogCancel = (e) => {
        e.preventDefault()
        const { id } = e.currentTarget
        let tmp = this.state.data
        tmp[id].isSelected = !tmp[id].isSelected
        tmp[id].isOpen = !tmp[id].isOpen
        tmp[id].color = 'default'
        this.setState({ arr: tmp })
    }

    handleSubmit = (e, id, index) => {
        const { data, selected } = this.state
        var d = data.find((d) => d.id === id)
        selected.push(d)
        this.setState({ data, selected })
        console.log(this.state.selected)
        e.preventDefault()

        data[index].isOpen = !data[index].isOpen
        this.setState({ arr: data })
        this.state.open = true
    }

    handleChipClick = (e, options) => {
        e.preventDefault()
        const { id } = e.currentTarget
        let tmp = this.state.data
        tmp[id].isSelected = !tmp[id].isSelected
        tmp[id].isOpen = !tmp[id].isOpen
        tmp[id].color = 'primary'
        this.setState({ arr: tmp })
    }

    handleChipDelete = (index) => (event) => {
        console.log(index)
    }

    render() {
        return (
            <>
                {this.state.data.map((step, index) => (
                    <>
                        <Chip
                            label={step.name}
                            id={index}
                            key={index}
                            clickable
                            onClick={(e) => this.handleChipClick(e, step.options)}
                            color={step.color}
                        />
                        <Dialog open={step.isOpen} onClose={this.handleClose} aria-labelledby="form-dialog-title">
                            <DialogTitle id="form-dialog-title">{step.name}</DialogTitle>
                            <DialogContent>
                                <TextField
                                    name="Sequence"
                                    onChange={(e) => this.handleChangeSequence(e, index)}
                                    autoFocus
                                    margin="dense"
                                    label="Sequence"
                                    type="number"
                                    value={step.sequence}
                                    fullWidth
                                    required
                                    key={step.id}
                                />
                                {step.options.map((option, i) => (
                                    <TextField
                                        margin="dense"
                                        name={option}
                                        onChange={(e) => this.handleChange(e, index, i)}
                                        key={option}
                                        label={option}
                                        value={step.optionsFilled[i]}
                                        type="text"
                                        fullWidth
                                    />
                                ))}
                            </DialogContent>
                            <DialogActions>
                                <Button id={index} onClick={(e) => this.handleDialogCancel(e)}>
                                    Cancel
                                </Button>
                                <Button
                                    id={index}
                                    onClick={(e) => this.handleSubmit(e, step.id, index)}
                                    color="primary"
                                >
                                    Submit
                                </Button>
                            </DialogActions>
                        </Dialog>
                    </>
                ))}
            </>
        )
    }
}

export default CustomChip
