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
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { getInitials } from "@/lib/utils";
import { useParticipant } from "@/store/ParticipantStore";
import { ParticipantType } from "@/constants/ParticipantType";
import { USERS } from "@/constants/User";
import { User } from "lucide-react";

export function AddUser({ type }: { type: ParticipantType }) {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="ghost" className="px-3">
          <User size={16} />
        </Button>
      </SheetTrigger>
      <SheetContent side="bottom" className="h-[90vh]">
        <SheetHeader>
          <SheetTitle>Add Participant</SheetTitle>
          <SheetDescription></SheetDescription>
        </SheetHeader>
        <div className="flex flex-col gap-2 mt-4 pr-0">
          {USERS.map(name => (
            <UserItem name={name} type={type} key={name} />
          ))}
        </div>

        <SheetFooter>
          <SheetClose asChild>
            <Button className="float-right my-4">Save</Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
}

const UserItem = ({ name, type }: { name: string; type: ParticipantType }) => {
  const { owers, setOwers, payers, setPayers } = useParticipant();
  if (type == ParticipantType.OWER)
    return (
      <div
        className={`grid grid-cols-6 items-center p-2 ${
          owers.includes(name) ? "bg-transparent backdrop-brightness-150" : ""
        }`}
        onClick={() => {
          if (owers.includes(name)) {
            setOwers(owers.filter((user: string) => user !== name));
          } else {
            setOwers([...owers, name]);
          }
        }}
      >
        <Avatar>
          <AvatarImage src="" />
          <AvatarFallback>{getInitials(name)}</AvatarFallback>
        </Avatar>
        <p className="col-span-4">{name}</p>
      </div>
    );
  else
    return (
      <div
        className={`grid grid-cols-6 items-center p-2 ${
          payers.includes(name) ? "bg-transparent backdrop-brightness-150" : ""
        }`}
        onClick={() => {
          if (payers.includes(name)) {
            setPayers(payers.filter((user: string) => user !== name));
          } else {
            setPayers([...payers, name]);
          }
        }}
      >
        <Avatar>
          <AvatarImage src="" />
          <AvatarFallback>{getInitials(name)}</AvatarFallback>
        </Avatar>
        <p className="col-span-4">{name}</p>
      </div>
    );
};