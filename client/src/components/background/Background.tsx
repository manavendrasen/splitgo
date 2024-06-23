import { PinkShape, PurpleShape, WhiteShape } from "./Shapes";

export const Background = () => (
  <div className="absolute left-0 top-0 -z-40 overflow-hidden opacity-100 bg-gradient-to-t from-background to-transparent">
    <div className="relative h-screen w-screen scale-150">
      <PurpleShape className="absolute -top-20  -right-20 blur-3xl opacity-80 animate-spin repeat-infinite spin-in-360 duration-20s direction-alternate delay-30" />
      <PurpleShape className="absolute -top-20   -left-20 blur-3xl opacity-50  animate-spin repeat-infinite spin-out-45 duration-20s delay-150" />
      <PinkShape className="absolute -top-48  -right-20  blur-3xl animate-spin repeat-infinite spin-out-360 duration-20s delay-300" />
      <PinkShape className="absolute -top-60  -left-32  blur-3xl scale-129 opacity-80 animate-spin  repeat-infinite spin-in-360 duration-20s delay-700" />
      <WhiteShape className="absolute -top-48  -left-20  blur-3xl animate-spin  repeat-infinite spin-out-360 duration-20s opacity-50" />
    </div>
  </div>
);

export const TopBackground = () => (
  <div className="absolute left-0 top-0 -z-40 overflow-hidden h-[calc(100vh-74px)] opacity-80">
    <div className="relative h-screen w-screen ">
      <PurpleShape className="absolute -top-96 -left-20 blur-3xl opacity-80 animate-spin repeat-infinite spin-in-360 duration-20s direction-alternate delay-30" />
      <PurpleShape className="absolute -top-96  -left-20 blur-3xl opacity-50  animate-spin repeat-infinite spin-out-45 duration-20s delay-150" />
      <PinkShape className="absolute -top-96 -left-20  blur-3xl animate-spin repeat-infinite spin-out-360 duration-20s delay-300" />
      <PinkShape className="absolute -top-96  -left-20  blur-3xl scale-129 opacity-80 brightness-125 animate-spin  repeat-infinite spin-in-360 duration-20s delay-700" />
      <WhiteShape className="absolute -top-96 -left-20  blur-3xl animate-spin  repeat-infinite spin-out-360 duration-20s opacity-55" />
    </div>
  </div>
);
