import { useState, useEffect } from 'react'
import api from './api/axios'
import './App.css'

function App() {
  const [health, setHealth] = useState<string>('')

  useEffect(() => {
    // Test API connection
    api.get('/health')
      .then(response => {
        setHealth(response.data.status)
      })
      .catch(error => {
        console.error('Error:', error)
      })
  }, [])

  return (
    <div className="App">
      <h1>Vite + React + Go</h1>
      <p>API Status: {health}</p>
    </div>
  )
}

export default App