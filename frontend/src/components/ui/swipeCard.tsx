import { BadgeCheck, Languages, MapPin, Briefcase, Info } from "lucide-react";
import { Badge } from "@/components/ui/badge";


export default function SwipeCard() {
  return (
    <div className="rounded min-w-full min-h-full w-[150%] overflow-hidden shadow-lg relative ">
      <img
        className="w-full rounded pointer-events-none"
        src="https://images.pexels.com/photos/2182970/pexels-photo-2182970.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        alt="Background Image"
      />
      <div className="absolute pointer-events-none h-full inset-0 flex flex-col items-start justify-end p-4 bg-gradient-to-t from-black/75 to-black/25 pr-12 gap-4">
        <div className="font-bold text-white text-4xl mb-2 text-start text-wrap w-[100%] flex flex-col">
          <span>Dustin</span>
          <span className="flex flex-row items-end gap-2">
            Ellenpark <BadgeCheck className="text-yellow-500 size-7" />
          </span>

          <div className="flex flex-col text-start text-base mt-4 gap-1">
            <span className="text-white font-semibold text-lg flex flex-row items-center gap-2 text-md">
              About Me
            </span>
            <p className="text-white font-normal text-sm">
            Passionate about building scalable applications that make an impact.
            10+ years in software engineering, currently leading teams in AI
            and cloud computing.
            </p>
          </div>

        </div>
        <div className="flex flex-col gap-2">
        <span className="text-lg font-normal flex flex-row gap-2 items-center">
            <Briefcase className="text-white size-5 mr-1" />
            <Badge variant="outline" className="px-2 text-sm text-nowrap font-normal">
              Senior Software Engineer
            </Badge>
          </span>
          <span className="flex flex-row gap-2 items-center">
            <MapPin className="text-white size-5 mr-1" />
            <Badge variant="outline" className="font-normal">New York, USA</Badge>
          </span>
          <span className="flex flex-row gap-2 items-center">
            <Languages className="text-white size-5 mr-1" />
            <Badge variant="outline" className="font-normal">English</Badge>
            <Badge variant="outline" className="font-normal">Estonian</Badge>
          </span>
        </div>
      </div>
    </div>
  );
}
