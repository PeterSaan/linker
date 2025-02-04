import { useState, useEffect } from 'react'
import api from './api/axios'
import { Button } from "@/components/ui/button"
import TinderCard from 'react-tinder-card'
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

  const onSwipe = (direction : string) => {
    console.log('You swiped: ' + direction)
  }

  const onCardLeftScreen = (myIdentifier : string) => {
    console.log(myIdentifier + ' left the screen')
  }

  return (
    <div className="App">
      <h1 className='text-3xl font-bold underline'>Vite + React + Go</h1>
      <Button>Click me</Button>
      <p>API Status: {health}</p>

      <TinderCard onSwipe={onSwipe} onCardLeftScreen={() => onCardLeftScreen('fooBar')} preventSwipe={['up', 'down']}>Hello, World!</TinderCard>
    </div>
  )
}

export default App