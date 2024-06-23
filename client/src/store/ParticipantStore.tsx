import { USERS } from "@/constants/User";
import React, { createContext, useContext, useState } from "react";

type ParticipantProviderState = {
  payers: string[];
  setPayers: (payers: string[]) => void;
  owers: string[];
  setOwers: (owers: string[]) => void;
};

const initialState: ParticipantProviderState = {
  payers: [USERS[0]],
  setPayers: () => {},
  owers: [],
  setOwers: () => {},
};

const ParticipantContext =
  createContext<ParticipantProviderState>(initialState);

export function ParticipantProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [payers, setPayers] = useState<string[]>([USERS[0]]);
  const [owers, setOwers] = useState<string[]>([]);
  return (
    <ParticipantContext.Provider
      value={{
        payers,
        setPayers,
        owers,
        setOwers,
      }}
    >
      {children}
    </ParticipantContext.Provider>
  );
}

export const useParticipant = () => useContext(ParticipantContext);
