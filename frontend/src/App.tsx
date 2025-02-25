import { useState, useEffect, useMemo, useRef } from "react";
import api from "./api/axios";
import { Button } from "@/components/ui/button";
import SwipeCard from "./components/ui/swipeCard";
import TinderCard from "react-tinder-card";
import "./App.css";
import React from "react";

const db = [
  {
    name: "Dustin Ellenpark",
    url: "https://images.pexels.com/photos/2182970/pexels-photo-2182970.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    name: "Erlich Bachman",
    url: "./img/erlich.jpg",
  },
  {
    name: "Monica Hall",
    url: "./img/monica.jpg",
  },
  {
    name: "Jared Dunn",
    url: "./img/jared.jpg",
  },
  {
    name: "Dinesh Chugtai",
    url: "./img/dinesh.jpg",
  },
];

export default function App() {
  const [health, setHealth] = useState<string>("");
  const [currentIndex, setCurrentIndex] = useState(db.length - 1);
  const [lastDirection, setLastDirection] = useState();

  const currentIndexRef = useRef(currentIndex);

  useEffect(() => {
    // Test API connection
    api
      .get("/health")
      .then((response) => {
        setHealth(response.data.status);
      })
      .catch((error) => {
        console.error("Error:", error);
        setHealth("Error");
      });
  }, []);
  const childRefs = useMemo(
    () =>
      Array(db.length)
        .fill(0)
        .map(() => React.createRef<any>()),
    []
  );

  useEffect(() => {
    // Test API connection
    api
      .get("/health")
      .then((response) => {
        setHealth(response.data.status);
      })
      .catch((error) => {
        console.error("Error:", error);
        setHealth("Error");
      });
  }, []);

  const updateCurrentIndex = (val: number) => {
    setCurrentIndex(val);
    currentIndexRef.current = val;
  };

  const canGoBack = currentIndex < db.length - 1;
  const canSwipe = currentIndex >= 0;

  // Triggered when a card is swiped
  const swiped = (direction: string, nameToDelete: string, index: number) => {
    setLastDirection(direction);
    updateCurrentIndex(index - 1);
  };

  // Called when a card leaves the screen
  const outOfFrame = (name: string, idx: number) => {
    console.log(`${name} (${idx}) left the screen!`, currentIndexRef.current);
    // If goBack was triggered before card fully left, restore it.
    if (currentIndexRef.current >= idx) {
      childRefs[idx].current.restoreCard();
    }
  };

  // Programmatically trigger a swipe
  const swipe = async (dir: "left" | "right") => {
    if (canSwipe && currentIndex < db.length) {
      await childRefs[currentIndex].current.swipe(dir);
    }
  };

  // Restore last swiped card
  const goBack = async () => {
    if (!canGoBack) return;
    const newIndex = currentIndex + 1;
    updateCurrentIndex(newIndex);
    await childRefs[newIndex].current.restoreCard();
  };

  return (
    <div className="App flex flex-col items-center min-h-screen">
      <h1 className="text-3xl font-bold underline">Vite + React + Go</h1>
      <Button>Click me</Button>
      <p>API Status: {health}</p>

      {db.map((character, index) => (
        <TinderCard
          ref={childRefs[index]}
          className="swipe"
          key={character.name}
          onSwipe={(dir) => swiped(dir, character.name, index)}
          onCardLeftScreen={() => outOfFrame(character.name, index)}
        >
          <SwipeCard />
        </TinderCard>
      ))}

      <div className="buttons flex gap-4 mt-4">
        <button
          style={{ backgroundColor: !canSwipe ? "#c3c4d3" : "#fff" }}
          onClick={() => swipe("left")}
          className="p-2 rounded"
        >
          Swipe left!
        </button>
        <button
          style={{ backgroundColor: !canGoBack ? "#c3c4d3" : "#fff" }}
          onClick={() => goBack()}
          className="p-2 rounded"
        >
          Undo swipe!
        </button>
        <button
          style={{ backgroundColor: !canSwipe ? "#c3c4d3" : "#fff" }}
          onClick={() => swipe("right")}
          className="p-2 rounded"
        >
          Swipe right!
        </button>
      </div>

      {lastDirection ? (
        <h2 key={lastDirection} className="infoText mt-4">
          You swiped {lastDirection}
        </h2>
      ) : (
        <h2 className="infoText mt-4">
          Swipe a card or press a button to restore!
        </h2>
      )}
    </div>
  );
}
