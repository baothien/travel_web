/* eslint-disable @typescript-eslint/no-empty-interface */
import * as React from "react";
import "./CheckIn.scss";

import { HandleCheckIn } from "./components/HandleCheckIn";
import { PreCheckIn } from "./components/PreCheckIn";

import { Row } from "antd";

export interface ICheckInProps {}

export default function CheckIn(props: ICheckInProps) {
  const [dataImgURL, setDataImgURL] = React.useState<string>("");

  return (
    <div style={{ position: "relative" }}>
      <div className="vw-100 bg-image"></div>

      <Row
        style={{
          width: "100%",
          position: "absolute",
          top: "5%",
        }}
        className="check-in"
        justify="center"
        align="middle"
      >
        {dataImgURL ? (
          <HandleCheckIn bgURL={dataImgURL} setDataImgURL={setDataImgURL} />
        ) : (
          <PreCheckIn setDataImgURL={setDataImgURL} />
        )}
      </Row>
    </div>
  );
}
