import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";

export default function Login(){
  return (
    <Card className="w-full max-w-sm p-6 shadow-lg rounded-2xl">
      <h2 className="text-2xl font-bold text-center">Create an account</h2>
      <p className="text-sm text-gray-600 text-center mb-4">
        Enter your email below to create your account
      </p>
      <div className="flex mb-5">
        <Button variant="outline" className="flex-1 flex items-center">
         LinkedIn
        </Button>
      </div>
      <div className="relative">
        <Separator />
        <span className="absolute left-1/2 top-[-10px] bg-white px-2 text-xs text-gray-500 transform -translate-x-1/2">
          OR CONTINUE WITH
        </span>
      </div>
      <form className="space-y-4 mt-4">
        <Input type="email" placeholder="m@example.com" className="w-full" />
        <Input type="password" placeholder="Password" className="w-full" />
        <Button className="w-full">Create account</Button>
      </form>
    </Card>
  );
}