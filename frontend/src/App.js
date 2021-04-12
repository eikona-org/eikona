import React from 'react'
import './App.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Dashboard from './components/Dashboard'
import AppBar from './components/AppBar'
import Login from './components/Login'
import Register from './components/Register'
import Upload from './components/Upload'
import Process from './components/Process'
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles'
import useToken from './components/useToken'
import CssBaseline from '@material-ui/core/CssBaseline'

const outerTheme = createMuiTheme({
    palette: {
        primary: {
            main: '#303F9F',
        },
        secondary: {
            light: '#FFC107',
            main: '#FF5252',
            contrastText: '#ffcc00',
        },
    },
})

export default function App() {
    const { token, setToken } = useToken()

    if (!token) {
        return (
            <BrowserRouter>
                <ThemeProvider theme={outerTheme}>
                    <CssBaseline />
                    <Routes>
                        <Route path="/register" element={<Register />} />
                        <Route path="*" element={<Login setToken={setToken} />} />
                    </Routes>
                </ThemeProvider>
            </BrowserRouter>
        )
    }
    return (
        <BrowserRouter>
            <ThemeProvider theme={outerTheme}>
                <CssBaseline />
                <AppBar />
                <Routes>
                    <Route path="/upload" element={<Upload />} />
                    <Route path="/process" element={<Process />} />
                    <Route path="/*" element={<Dashboard />} />
                </Routes>
            </ThemeProvider>
        </BrowserRouter>
    )
}
