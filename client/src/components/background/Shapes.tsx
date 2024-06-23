import React from "react";
import { JSX } from "react/jsx-runtime";

export const PinkShape = (
  props: JSX.IntrinsicAttributes & React.SVGProps<SVGSVGElement>
) => (
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200" {...props}>
    <path
      fill="#FF669A"
      d="M143.9 41.9c12.1 6.5 20.3 20.5 27.9 36.5 7.5 16 14.4 33.9 10 48.4-4.5 14.4-20.2 25.3-35.8 30.9-15.5 5.6-30.7 5.9-44.7 4.1-14-1.8-26.6-5.7-42.3-11.2-15.6-5.6-34.3-12.8-41.7-26-7.3-13.1-3.4-32.2 5.3-47 8.6-14.7 22-25.1 35.6-31.1 13.7-5.9 27.8-7.3 42.7-8.6 15-1.3 30.9-2.4 43 4Z"
    />
  </svg>
);

export const PurpleShape = (
  props: JSX.IntrinsicAttributes & React.SVGProps<SVGSVGElement>
) => (
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200" {...props}>
    <path
      fill="#8884FF"
      d="M139 35.7c13.8 4.6 28.8 10.6 35 21.5 6.1 10.9 3.4 26.9 3.9 43.1.5 16.2 4 32.6-1.9 44-5.9 11.3-21.3 17.5-35.6 23.3-14.2 5.7-27.3 11.1-40.9 11.9-13.5.8-27.5-3-38.3-10.8-10.7-7.8-18.2-19.6-26.5-31.8-8.4-12.1-17.6-24.5-20.7-38.7-3.1-14.2-.2-30.3 7.4-43.6 7.7-13.3 20-24 33.9-28.6 13.9-4.5 29.3-3 43.3-.6 14 2.4 26.6 5.7 40.4 10.3Z"
    />
  </svg>
);

export const WhiteShape = (
  props: JSX.IntrinsicAttributes & React.SVGProps<SVGSVGElement>
) => (
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200" {...props}>
    <path
      fill="#FBF5F3"
      d="M131.8 53.9c12-2.1 26 1.3 34.2 9.9 8.2 8.6 10.7 22.4 11.2 36.5.5 14.1-.8 28.6-7.5 39.8-6.7 11.3-18.7 19.5-31.1 25.6-12.5 6.2-25.6 10.4-38.6 10.5-13.1 0-26.2-4-28.2-16.3-1.9-12.2 7.4-32.7 12.4-44.1 5.1-11.3 6.1-13.6 4.6-16.7-1.5-3-5.4-7-6.6-11.7-1.3-4.8 0-10.4 3.4-13.1 3.4-2.8 8.9-2.8 16.6-6.6 7.7-3.9 17.7-11.6 29.6-13.8Z"
    />
  </svg>
);

export const Grain = (
  props: JSX.IntrinsicAttributes & React.SVGProps<SVGSVGElement>
) => (
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 250 250" {...props}>
    <filter id="a">
      <feTurbulence
        baseFrequency={0.65}
        numOctaves={3}
        stitchTiles="stitch"
        type="fractalNoise"
      />
    </filter>
    <rect width="100%" height="100%" filter="url(#a)" />
  </svg>
);
