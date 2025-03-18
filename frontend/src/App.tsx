import { useState } from 'react'
import { Button } from "@/components/ui/button"
import TinderCard from 'react-tinder-card'
import './App.css'

import Login from './pages/login';

function App() {
  const [health, setHealth] = useState<string>('')

  const onSwipe = (direction : string) => {
    console.log('You swiped: ' + direction)
  }

  const onCardLeftScreen = (myIdentifier : string) => {
    console.log(myIdentifier + ' left the screen')
  }

  return (
    <div className="App">
        <Login/>    
    </div>
  )
}

export default App
