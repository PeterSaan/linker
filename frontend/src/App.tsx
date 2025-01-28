import { useState, useEffect } from 'react'
import api from './api/axios'
import { Button } from "@/components/ui/button"
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
        setHealth('Error')
      })
  }, [])

  return (
    <div className="App">
      <h1 className='text-3xl font-bold underline'>Vite + React + Go</h1>
      <Button>Click me</Button>
      <p>API Status: {health}</p>
    </div>
  )
}

export default App