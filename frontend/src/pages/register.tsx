import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useState } from "react";
import api from "@/api/axios";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormMessage,
} from "@/components/ui/form";

export default function Register() {
    const [error, setError] = useState("");

    const formSchema = z
        .object({
            email: z.string().email(),
            password: z.string(),
            password_confirm: z.string(),
        })
        .refine((data) => data.password === data.password_confirm, {
            message: "Passwords must match",
            path: ["password_confirm"],
        });

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            email: "",
            password: "",
            password_confirm: "",
        },
    });

    const submit = async (values: z.infer<typeof formSchema>) => {
        await api
            .post("/auth/register", {
                email: values.email,
                password: values.password,
                password_confirm: values.password_confirm,
            })
            .catch(function (error) {
                setError(error.response.data.error);
            });
    };

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
            <Form {...form}>
                <form className="space-y-4 mt-4" onSubmit={form.handleSubmit(submit)}>
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input
                                        type="email"
                                        placeholder="mail@example.com"
                                        className="w-full"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input
                                        type="password"
                                        placeholder="password"
                                        className="w-full"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password_confirm"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input
                                        type="password"
                                        placeholder="password confirmation"
                                        className="w-full"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    {error && (
                        <p className="text-[0.8rem] font-medium text-red-600">{error}</p>
                    )}
                    <Button type="submit" className="w-full">
                        Create account
                    </Button>
                </form>
            </Form>
        </Card>
    );
}
