import { useState, useEffect, useMemo, useRef } from "react";
import api from "./api/axios";
import { Button } from "@/components/ui/button";
import SwipeCard from "./components/ui/swipeCard";
import TinderCard from "react-tinder-card";
import "./App.css";
import React from "react";

const db = [
  {
    firstname: "Dustin",
    lastName: "Ellenpark",
    jobTitle: "Senior Software Engineer",
    location: "New York, USA",
    languages: ["English", "Estonian"],
    about: "Passionate about building scalable applications that make an impact. 10+ years in software engineering, currently leading teams in AI and cloud computing.",
    imageUrl: "https://images.pexels.com/photos/2182970/pexels-photo-2182970.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    firstname: "Mari",
    lastName: "Maasikas",
    jobTitle: "Marketing Expert",
    location: "Tallinn, Estonia",
    languages: ["English", "Estonian"],
    about:
      "I'm a marketing expert with a passion for building brands and creating engaging content.",
    imageUrl:
      "https://images.pexels.com/photos/227294/pexels-photo-227294.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    firstname: "Liam",
    lastName: "Callahan",
    jobTitle: "Data Scientist",
    location: "San Francisco, USA",
    languages: ["English", "Spanish"],
    about:
      "Data enthusiast with a love for AI and machine learning. Always exploring new ways to leverage data for business insights.",
    imageUrl:
      "https://images.pexels.com/photos/1183266/pexels-photo-1183266.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    firstname: "Sofia",
    lastName: "Lindström",
    jobTitle: "UX/UI Designer",
    location: "Stockholm, Sweden",
    languages: ["Swedish", "English"],
    about:
      "Creative UX/UI designer focused on crafting intuitive and user-friendly digital experiences.",
    imageUrl:
      "https://images.pexels.com/photos/1130626/pexels-photo-1130626.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    firstname: "Kenji",
    lastName: "Takahashi",
    jobTitle: "Cybersecurity Specialist",
    location: "Tokyo, Japan",
    languages: ["Japanese", "English"],
    about:
      "Cybersecurity professional with expertise in ethical hacking and digital forensics. Passionate about securing the digital world.",
    imageUrl:
      "https://images.pexels.com/photos/8765940/pexels-photo-8765940.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
  {
    firstname: "Elena",
    lastName: "Rodriguez",
    jobTitle: "Biomedical Engineer",
    location: "Madrid, Spain",
    languages: ["Spanish", "English"],
    about:
      "Bridging the gap between technology and healthcare to create innovative medical solutions.",
    imageUrl:
      "https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  },
];

export default function App() {
  const [health, setHealth] = useState<string>("");
  const [currentIndex, setCurrentIndex] = useState(db.length - 1);
  const [lastDirection, setLastDirection] = useState<string | undefined>();
  const [cardHeight, setCardHeight] = useState(0);
  const currentIndexRef = useRef(currentIndex);
  const cardContainerRef = useRef(null);
  const cardRefs = useRef<(HTMLDivElement | null)[]>([]);

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

  // Update card height when cards are rendered
  useEffect(() => {
    const updateCardHeight = () => {
      // Find the maximum height among all card elements
      let maxHeight = 0;
      cardRefs.current.forEach((ref) => {
        if (ref) {
          const height = (ref as HTMLElement).offsetHeight;
          console.log("Card height:", height);
          if (height > maxHeight) {
            maxHeight = height;
          }
        }
      });

      // Set the container height to match the tallest card
      if (maxHeight > 0) {
        setCardHeight(maxHeight);
      } else {
        // Fallback height if cards aren't measured yet
        setCardHeight(600); // Default height
      }
    };

    // Initial height calculation
    updateCardHeight();

    // Add window resize listener for responsive behavior
    window.addEventListener("resize", updateCardHeight);

    // Cleanup listener
    return () => window.removeEventListener("resize", updateCardHeight);
  }, [currentIndex]);

  const childRefs = useMemo(
    () =>
      Array(db.length)
        .fill(0)
        .map(() => React.createRef<any>()),
    []
  );

  const updateCurrentIndex = (val: number) => {
    setCurrentIndex(val);
    currentIndexRef.current = val;
  };

  const canGoBack = currentIndex < db.length - 1;
  const canSwipe = currentIndex >= 0;

  // Triggered when a card is swiped
  const swiped = (direction: string, index: number) => {
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
    <div className="flex flex-col items-center min-h-screen p-4 overflow-x-hidden">
      <header className="w-full max-w-md flex justify-between items-center p-4 mb-6">
        <div className="text-2xl font-bold bg-gradient-to-r from-blue-500 to-purple-500 text-transparent bg-clip-text">
          Linker
        </div>
        <div className="flex items-center space-x-2">
          <div
            className={`h-3 w-3 rounded-full ${
              health === "Connected" ? "bg-green-500" : "bg-red-500"
            }`}
          ></div>
          <span className="text-xs text-gray-500">{health}</span>
        </div>
      </header>

      <div
        ref={cardContainerRef}
        className="relative w-full max-w-md flex justify-center mb-8"
        style={{ height: `${cardHeight}px` }}
      >
        {db.map((character, index) => (
          <TinderCard
            ref={childRefs[index]}
            className="absolute"
            key={character.firstname}
            onSwipe={(dir) => swiped(dir, index)}
            onCardLeftScreen={() => outOfFrame(character.firstname, index)}
            preventSwipe={["up", "down"]}
            onSwipeRequirementUnfulfilled={() => {
              // Prevent swipe if currentIndex is not equal to the index of the card
              if (currentIndex !== index) {
                return true;  // Prevent swipe action
              }
              return false;  // Allow swipe action
            }}
          >
            <div
              ref={(el) => (cardRefs.current[index] = el)}
              className="relative rounded-xl shadow-xl h-full w-full overflow-hidden"
            >
              <div className="absolute inset-0 bg-center bg-cover" />
              <SwipeCard user={character} />
            </div>
          </TinderCard>
        ))}

        {db.length === 0 && (
          <div className="text-center p-4 rounded-lg shadow">
            <h3 className="text-xl font-semibold text-gray-700">
              No more profiles!
            </h3>
            <p className="text-gray-500 mt-2">
              Check back later for more matches
            </p>
          </div>
        )}
      </div>

      <div className="flex justify-center items-center space-x-4 mt-2">
        <button
          onClick={() => swipe("left")}
          disabled={!canSwipe}
          className="flex items-center justify-center w-14 h-14 rounded-full bg-white shadow-lg disabled:opacity-50 disabled:cursor-not-allowed transition-transform hover:scale-110"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-8 w-8 text-red-500"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clipRule="evenodd"
            />
          </svg>
        </button>

        <button
          onClick={() => goBack()}
          disabled={!canGoBack}
          className="flex items-center justify-center w-10 h-10 rounded-full bg-white shadow-lg disabled:opacity-50 disabled:cursor-not-allowed transition-transform hover:scale-110"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-6 w-6 text-yellow-500"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L4.414 9H17a1 1 0 110 2H4.414l5.293 5.293a1 1 0 010 1.414z"
              clipRule="evenodd"
            />
          </svg>
        </button>

        <button
          onClick={() => swipe("right")}
          disabled={!canSwipe}
          className="flex items-center justify-center w-14 h-14 rounded-full bg-white shadow-lg disabled:opacity-50 disabled:cursor-not-allowed transition-transform hover:scale-110"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-8 w-8 text-green-500"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
              clipRule="evenodd"
            />
          </svg>
        </button>
      </div>

      {lastDirection && (
        <div
          className={`mt-8 px-6 py-2 rounded-full bg-white shadow-md ${
            lastDirection === "right" ? "text-green-500" : "text-red-500"
          }`}
        >
          {lastDirection === "right"
            ? "You liked this profile!"
            : "You passed on this profile"}
        </div>
      )}
    </div>
  );
}
