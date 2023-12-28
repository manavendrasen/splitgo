import React, { useEffect, useMemo } from "react";
import { Button } from "../components/ui/button";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
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
import { useParticipant } from "@/store/ParticipantStore";
import { ParticipantType } from "@/constants/ParticipantType";
import { getInitials } from "@/lib/utils";

const AddTransaction = () => {
  const [totalPayerAmount, setTotalPayerAmount] = useState(0);
  const { owers, payers } = useParticipant();

  const formSchema = useMemo(() => {
    const scheme = {
      description: z.string().min(2).max(50),
    } as Record<string, z.ZodString | z.ZodNumber>;

    owers.forEach((ower: string) => {
      scheme[`ower-${ower}`] = z.coerce
        .number()
        .min(0, { message: "Amount should be greater than 0" })
        .max(1000000, { message: "Amount should be less than 1000000" });
    });

    payers.forEach((payer: string) => {
      scheme[`payer-${payer}`] = z.coerce
        .number()
        .min(0, { message: "Amount should be greater than 0" })
        .max(1000000, { message: "Amount should be less than 1000000" });
    });

    return z.object(scheme);
  }, [owers, payers]);

  // -------------------------first-------------------
  // Form
  // -------------------------------------------------

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
    const owerTotal = Object.keys(values).reduce<number>((acc, key) => {
      if (key.includes("ower")) {
        return acc + (!values[key] ? 0 : Number(values[key]));
      }
      return acc;
    }, 0);

    console.log(owerTotal, totalPayerAmount);

    if (
      owerTotal !== totalPayerAmount ||
      totalPayerAmount >= 0 ||
      owerTotal >= 0
    ) {
      toast.error(
        `Split amount should be equal to Paid amount. Adjust Rs.${Math.abs(
          owerTotal - totalPayerAmount
        )}.`
      );
    }
  }

  useEffect(() => {
    const subscription = form.watch((value, { name, type }) => {
      if (type === "change" && name?.includes("payer")) {
        const total = Object.keys(value).reduce<number>((acc, key) => {
          if (key.includes("payer")) {
            return acc + (!value[key] ? 0 : Number(value[key]));
          }
          return acc;
        }, 0);

        setTotalPayerAmount(total);
      }
    });
    return () => subscription.unsubscribe();
  }, [form, owers, payers]);

  useEffect(() => {
    if (totalPayerAmount) {
      owers.forEach((ower: string) => {
        form.setValue(
          `ower-${ower}`,
          (totalPayerAmount / owers.length).toFixed(2)
        );
      });
    }
  }, [form, owers, totalPayerAmount]);

  useEffect(() => {
    if (form.formState.errors) {
      toast.error("Please fill all the fields with valid values.");
    }
  }, [form]);

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

          <div className="flex flex-col gap-4">
            <div className="flex justify-between items-center">
              <p className="text-sm uppercase text-slate-300">Paid By</p>
              <AddUser type={ParticipantType.PAYER} />
            </div>
            {payers.map((payer: string) => (
              <UserItem
                key={payer}
                name={payer}
                form={
                  owers.length > 0 && (
                    <FormField
                      control={form.control}
                      name={`payer-${payer}`}
                      render={({ field }) => (
                        <div className="col-span-2 text-right">
                          <FormItem>
                            <FormControl>
                              <Input
                                placeholder="0.00"
                                inputMode="numeric"
                                {...field}
                                className="text-right"
                              />
                            </FormControl>
                          </FormItem>
                        </div>
                      )}
                    />
                  )
                }
              />
            ))}
            {payers.length == 0 && (
              <p className="text-xs text-slate-400 text-center mb-12 ">
                Add Friend to split expense.
              </p>
            )}
            {payers.length > 0 && (
              <>
                <div className="flex justify-between items-center">
                  <p className="text-sm uppercase text-slate-300">
                    Split Between
                  </p>
                  <AddUser type={ParticipantType.OWER} />
                </div>
                {owers.length == 0 && (
                  <p className="text-xs text-slate-400 text-center">
                    Add Friend
                  </p>
                )}
              </>
            )}
            {owers.map((ower: string) => (
              <UserItem
                key={ower}
                name={ower}
                form={
                  <FormField
                    control={form.control}
                    name={`ower-${ower}`}
                    render={({ field }) => (
                      <div className="col-span-2 text-right">
                        <FormItem>
                          <FormControl>
                            <Input
                              placeholder="0.00"
                              inputMode="numeric"
                              {...field}
                              className="text-right"
                            />
                          </FormControl>
                        </FormItem>
                      </div>
                    )}
                  />
                }
              />
            ))}
          </div>

          {/* <div className="grid grid-cols-6 items-center">
                  <Avatar>
                    <AvatarImage src="" />
                    <AvatarFallback>MS</AvatarFallback>
                  </Avatar>
                  <p className="col-span-4">Manavendra Sen</p>
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
                </div> */}

          {owers.length > 0 && payers.length > 0 && (
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

const UserItem = ({ name, form }: { name: string; form: React.ReactNode }) => {
  return (
    <div key={name} className="grid grid-cols-6 items-center">
      <Avatar>
        <AvatarImage src="" />
        <AvatarFallback>{getInitials(name)}</AvatarFallback>
      </Avatar>
      <p className="col-span-3">{name}</p>
      {form}
    </div>
  );
};
