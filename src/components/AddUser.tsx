import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { SetStateAction, useState } from "react";
import { Label } from "./ui/label";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { getInitials } from "@/lib/utils";

export function AddUser() {
  const [inputValue, setInputValue] = useState("");
  const [selectedUsers, setSelectedUsers] = useState<string[]>([]);

  const MEMBERS = ["Somil Gupta", "Divyanshu Sharma"];

  const handleInputChange = (event: {
    target: { value: SetStateAction<string> };
  }) => {
    setInputValue(event.target.value);
  };
  return (
    <Sheet>
      <SheetTrigger asChild>
        <div className="grid grid-cols-6 items-center mb-8">
          <Avatar>
            <AvatarImage src="" />
            {/* TODO: set user name SheetHeader */}
            <AvatarFallback>+</AvatarFallback>
          </Avatar>
          <p className="col-span-4">Add a Friend</p>
        </div>
      </SheetTrigger>
      <SheetContent side="bottom" className="h-[90vh]">
        <SheetHeader>
          <SheetTitle>Add Participant</SheetTitle>
          <SheetDescription></SheetDescription>
        </SheetHeader>
        <div className="flex flex-col gap-4 mt-4">
          <Label>User Name</Label>
          <Input value={inputValue} onChange={handleInputChange} />
          {MEMBERS.map(name => (
            <div
              className={`grid grid-cols-6 items-center p-2 ${
                selectedUsers.includes(name) ? "bg-accent bg-opacity-10" : ""
              }`}
              onClick={() => {
                if (selectedUsers.includes(name)) {
                  setSelectedUsers(selectedUsers.filter(user => user !== name));
                } else {
                  setSelectedUsers([...selectedUsers, name]);
                }
              }}
            >
              <Avatar>
                <AvatarImage src="" />
                <AvatarFallback>{getInitials(name)}</AvatarFallback>
              </Avatar>
              <p className="col-span-4">{name}</p>
            </div>
          ))}
        </div>

        <SheetFooter>
          <SheetClose asChild>
            <Button
              className="float-right my-4"
              onClick={() => {
                console.log(selectedUsers);
              }}
            >
              Add
            </Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
}
