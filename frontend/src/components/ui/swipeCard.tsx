import { BadgeCheck } from 'lucide-react';

export default function SwipeCard() {
  return (
    <div className="rounded min-w-full min-h-full overflow-hidden shadow-lg relative">
      <img
        className="w-full rounded pointer-events-none"
        src="https://images.pexels.com/photos/2182970/pexels-photo-2182970.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        alt="Background Image"
      />
      <div className="absolute h-full inset-0 flex flex-col items-start backdrop-brightness-[75%]">
        <div className="font-bold text-4xl mb-2 text-start text-wrap w-[70%]">Dustin Ellenpark
        <BadgeCheck className='text-blue-500' />
        </div>
        <p className="text-gray-700 text-base text-start">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus
          quia, nulla! Maiores et perferendis eaque, exercitationem praesentium
          nihil.
        </p>
      </div>
    </div>
  );
}
