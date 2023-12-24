import { Button } from "../components/ui/button";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useState } from "react";
import { AddUser } from "@/components/AddUser";

const formSchema = z.object({
  description: z.string().min(2).max(50),
  amount: z.coerce.number().min(0.01).max(999999999.99),
});

const AddTransaction = () => {
  const [step, setStep] = useState(0);
  // const [participants, setParticipants] = useState<string[]>([]);
  // -------------------------first-------------------
  // Form
  // -------------------------------------------------

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
    if (step === 0) {
      setStep(1);
    }
  }

  return (
    <div className="text-left mt-8">
      <h1 className="text-2xl font-semibold mb-8">Add Expense</h1>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="flex flex-col gap-4"
        >
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <Input placeholder="Enter a description" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          {step === 0 && (
            <Button
              onClick={() => {
                setStep(1);
              }}
              className="mt-4"
            >
              Next
            </Button>
          )}

          {step === 1 && (
            <>
              <div className="flex flex-col gap-4">
                <p className="text-sm uppercase text-slate-300">Paid By</p>
                <div className="grid grid-cols-6 items-center">
                  <Avatar>
                    <AvatarImage src="" />
                    {/* TODO: set user name SheetHeader */}
                    <AvatarFallback>MS</AvatarFallback>
                  </Avatar>
                  <p className="col-span-4">Manavendra Sen</p>
                  <FormField
                    control={form.control}
                    name="amount"
                    render={({ field }) => (
                      <FormItem>
                        {/* <FormLabel>Amount</FormLabel> */}
                        <FormControl>
                          <Input placeholder="0.00" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <div className="grid grid-cols-6 items-center">
                  <Avatar>
                    <AvatarImage src="" />
                    <AvatarFallback>MS</AvatarFallback>
                  </Avatar>
                  <p className="col-span-4">Divyanshu Sharma</p>
                  <FormField
                    control={form.control}
                    name="amount"
                    render={({ field }) => (
                      <FormItem>
                        <FormControl>
                          <Input placeholder="0.00" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <AddUser />
                <div className="flex justify-between items-center">
                  <p className="text-sm uppercase text-slate-300">
                    Split Between
                  </p>
                  {/* <Button variant="outline">Edit</Button> */}
                </div>
                {/* {MEMBERS.map(name => ( */}
                {/* <div className="grid grid-cols-6 items-center">
                      <Avatar>
                        <AvatarImage src="" />
                        <AvatarFallback>
                          {name
                            .split(" ")
                            .map(n => n[0].toUpperCase())
                            .join("")}
                        </AvatarFallback>
                      </Avatar>
                      <p className="col-span-4">{name}</p>
                      <p className="text-right">100</p>
                    </div> */}
                {/* ))} */}
                {/* <hr /> */}
                {/* <AddUser */}
                {/* trigger={ */}

                {/* } */}
                {/* /> */}
                <div className="grid grid-cols-6 items-center">
                  <Avatar>
                    <AvatarImage src="" />
                    {/* TODO: set user name SheetHeader */}
                    <AvatarFallback>MS</AvatarFallback>
                  </Avatar>
                  <p className="col-span-4">Somil Gupta</p>
                  <FormField
                    control={form.control}
                    name="amount"
                    render={({ field }) => (
                      <FormItem>
                        {/* <FormLabel>Amount</FormLabel> */}
                        <FormControl>
                          <Input placeholder="0.00" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <AddUser />
              </div>
            </>
          )}

          {step === 1 && (
            <Button type="submit" className="mt-4">
              Split
            </Button>
          )}
        </form>
      </Form>
    </div>
  );
};

export default AddTransaction;
