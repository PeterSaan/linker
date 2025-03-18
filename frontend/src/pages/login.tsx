import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useState } from "react";
import api from "@/api/axios";

export default function Login(){
  const [password, setPassword] = useState("")
  const [email, setEmail] = useState("")
  const [error, setError] = useState("")

  const submit = async (e: any) => {
    e.preventDefault()

    await api.post("/auth/login", {
      password: password,
      email: email,
    }).catch(function (error) {
        setError(error.response.data.error)
    })
  }
  
  return (
    <Card className="w-full max-w-sm p-6 shadow-lg rounded-2xl">
      <h2 className="text-2xl font-bold text-center">Login</h2>
      <p className="text-sm text-gray-600 text-center mb-4">
        Enter your email below to login
      </p>
      <div className="flex mb-5">
        <Button asChild variant="outline" className="flex-1 flex items-center">
            <a href="/api/auth/linkedin/login">LinkedIn</a>
        </Button>
      </div>
      <div className="relative">
        <Separator />
        <span className="absolute left-1/2 top-[-10px] bg-white px-2 text-xs text-gray-500 transform -translate-x-1/2">
          OR CONTINUE WITH
        </span>
      </div>
      <form method="POST" onSubmit={submit} className="space-y-4 mt-4">
        <Input
            type="email" 
            placeholder="m@example.com"
            className="w-full"
            required
            onChange={(e) => setEmail(e.target.value)}
        />
        <Input
            type="password"
            placeholder="Password"
            className="w-full"
            required
            onChange={(e) => setPassword(e.target.value)}
        />
        {error && (
            <p className="text-red-600">{error}</p>
        )}
        <Button className="w-full" type="submit">Login</Button>
      </form>
    </Card>
  );
}
