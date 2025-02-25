import { useState, useEffect } from 'react'
import api from './api/axios'
import { Button } from "@/components/ui/button"
import TinderCard from 'react-tinder-card'
import './App.css'

import Login from './pages/login';

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

  const onSwipe = (direction : string) => {
    console.log('You swiped: ' + direction)
  }

  const onCardLeftScreen = (myIdentifier : string) => {
    console.log(myIdentifier + ' left the screen')
  }

  return (
    <div className="App">
      
    </div>
  )
}

export default App