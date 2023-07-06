import * as React from "react";
import {
  ReactCompareSlider,
  ReactCompareSliderImage,
} from "react-compare-slider";

export interface ICompareSliderProps {
  original: string;
  restored: string;
}

export function CompareSlider({ original, restored }: ICompareSliderProps) {
  return (
    <ReactCompareSlider
      style={{ height: "auto" }}
      itemOne={<ReactCompareSliderImage src={original} alt="original photo" />}
      itemTwo={<ReactCompareSliderImage src={restored} alt="restored photo" />}
      portrait
    />
  );
}
