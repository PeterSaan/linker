import { BadgeCheck } from "lucide-react";

export default function SwipeCard() {
  return (
    <div className="rounded min-w-full min-h-full overflow-hidden shadow-lg relative">
      <img
        className="w-full rounded pointer-events-none bac"
        src="https://images.pexels.com/photos/2182970/pexels-photo-2182970.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        alt="Background Image"
      />
      <div className="absolute h-full inset-0 flex flex-col items-start justify-end p-4 backdrop-brightness-[75%] pr-12 gap-4">
        <div className="font-bold text-white text-4xl mb-2 text-start text-wrap w-[70%] flex flex-col">
          <span>Dustin</span>
          <span className="flex flex-row items-end gap-2">Ellenpark <BadgeCheck className="text-yellow-500 size-7" /></span>
          <span className="text-lg font-normal mt-2">Passionate CEO</span>
        </div>
        <div className="flex flex-col text-start">
        <span className="text-white/50 font-bold text-lg">About Me</span>
        <p className="text-white text-md">
          Lorem ipsum is simply dummy text of the printing and typesetting industry.
        </p>
        </div>
      </div>
    </div>
  );
}
