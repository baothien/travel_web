import * as React from "react";
import { SyntheticEvent, useEffect, useRef, useState } from "react";

import html2canvas from "html2canvas";
import { SourceConfig, SourcePlayback } from "../helpers/sourceHelper";
import { BackgroundConfig } from "../helpers/backgroundHelper";
import useTFLite from "../hooks/useTFLite";

import OutputViewer from "./OutputViewer";
import uploadApi from "../../../apis/uploadApi";
import { Button, Row, Space, Spin } from "antd";
import { useAppDispatch } from "../../../hooks";
import { checkInActions } from "../checkInSlice";
import { useNavigate } from "react-router-dom";

export interface IHandleCheckInProps {
  bgURL: string;
  setDataImgURL: React.Dispatch<React.SetStateAction<string>>;
}

export function HandleCheckIn({ bgURL, setDataImgURL }: IHandleCheckInProps) {
  const dispatch = useAppDispatch();

  const navigate = useNavigate();

  // const sourceConfig: SourceConfig = {
  //   type: "camera",
  // };

  const backgroundConfig: BackgroundConfig = { type: "image", url: bgURL };

  const { tflite } = useTFLite();

  const [sourcePlayback, setSourcePlayback] = useState<SourcePlayback>();

  const [isLoading, setLoading] = useState(false);

  const videoRef = useRef<HTMLVideoElement>(null);

  const [imgURL, setImgURL] = useState("");

  useEffect(() => {
    setLoading(true);

    async function getCameraStream() {
      try {
        const constraint = { video: true };
        const stream = await navigator.mediaDevices.getUserMedia(constraint);
        if (videoRef.current) {
          videoRef.current.srcObject = stream;
          return;
        }
      } catch (error) {
        console.error("Error opening video camera.", error);
      }
      setLoading(false);
    }

    getCameraStream();
  }, []);

  function handleVideoLoad(event: SyntheticEvent) {
    const video = event.target as HTMLVideoElement;
    setSourcePlayback({
      htmlElement: video,
      width: video.videoWidth,
      height: video.videoHeight,
    });
    setLoading(false);
  }

  const handleCapture = () => {
    const canvasDup = document.getElementById("canvas-output");

    if (canvasDup) {
      html2canvas(canvasDup, {
        logging: true,
        useCORS: true,
      }).then((canvas) => {
        setImgURL(canvas.toDataURL("image/png"));

        fetch(canvas.toDataURL("image/png"))
          .then((response) => response.blob())
          .then((blob) => {
            const file = new File([blob], "checkIn.png", { type: blob.type });

            uploadApi
              .upload({ file: file, type: "check_in" })
              .then((res) => {
                dispatch(checkInActions.setCapturedImgUrl(res.data.full_path));
                navigate("/post-check-in");
              })
              .catch((err) => console.log(err));
          });
      });
    }
  };

  return (
    <Row
      style={{
        position: "relative",
        width: "50.8%",
        padding: "1rem",
        height: "36rem",
        border: "1px solid white",
        borderRadius: "0.5rem",
        boxShadow:
          "rgba(0, 0, 0, 0.2) 0px 11px 15px -7px, rgba(0, 0, 0, 0.14) 0px 24px 38px 3px, rgba(0, 0, 0, 0.12) 0px 9px 46px 8px",
      }}
      justify="center"
    >
      <Space
        style={{
          position: "absolute",
        }}
        direction="vertical"
        align="center"
      >
        {sourcePlayback && tflite ? (
          <OutputViewer
            sourcePlayback={sourcePlayback}
            backgroundConfig={backgroundConfig}
            tflite={tflite}
          />
        ) : (
          <Spin size="large" />
        )}

        <div
          style={{
            position: "absolute",
            top: 0,
            left: 0,
            zIndex: "-1",
          }}
        >
          <video
            ref={videoRef}
            src={""}
            hidden={isLoading}
            autoPlay
            playsInline
            controls={false}
            muted
            loop
            onLoadedData={handleVideoLoad}
          />
        </div>

        <Space>
          <Button
            style={{
              backgroundColor: "#999999",
              width: "7rem",
              height: "2.5rem",
            }}
            onClick={() => setDataImgURL("")}
            type="primary"
          >
            Quay lại
          </Button>

          <Button
            style={{
              width: "9rem",
              height: "2.5rem",
              backgroundColor: "#51ade5",
              padding: "9px 30px 10px",
            }}
            type="primary"
            onClick={handleCapture}
          >
            Check In
          </Button>
        </Space>
      </Space>
    </Row>
  );
}
