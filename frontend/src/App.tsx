import { useState, useEffect } from 'react'
import api from './api/axios'
import { Button } from "@/components/ui/button"
import SwipeCard from './components/ui/swipeCard'
import TinderCard from 'react-tinder-card'
import './App.css'

export default function App() {
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
    <div className="App flex flex-col items-center min-h-screen">
      <h1 className='text-3xl font-bold underline'>Vite + React + Go</h1>
      <Button>Click me</Button>
      <p>API Status: {health}</p>

      <TinderCard className="max-w-sm" onSwipe={onSwipe} onCardLeftScreen={() => onCardLeftScreen('fooBar')} preventSwipe={['up', 'down']}>
        <SwipeCard />
      </TinderCard>
    </div>
  )
}