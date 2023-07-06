/* eslint-disable @typescript-eslint/no-empty-interface */
import * as React from "react";
import ReactPannellum, { getContainer } from "react-pannellum";
import { Button, Col, Row, Space, Typography } from "antd";
import html2canvas from "html2canvas";
import { useAppDispatch, useAppSelector } from "../../../hooks";
import { checkInActions, selectBgImgURL } from "../checkInSlice";
import {
  ArrowLeftOutlined,
  HomeOutlined,
  RollbackOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { useNavigate } from "react-router-dom";

export interface IPreCheckInProps {
  setDataImgURL: React.Dispatch<React.SetStateAction<string>>;
}

export function PreCheckIn({ setDataImgURL }: IPreCheckInProps) {
  const navigate = useNavigate();

  const dispatch = useAppDispatch();

  const [chosenBgUrl, setChosenBgUrl] = React.useState("");

  const handleScreenshot = () => {
    const container = getContainer();

    const canvasDup = container.querySelector("canvas");

    html2canvas(canvasDup, {
      logging: true,
      useCORS: true,
    }).then((canvas) => {
      setDataImgURL(canvas.toDataURL("image/png"));
    });
  };

  const config = {
    autoLoad: true,
  };

  const urls = [
    "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Venice.Still001.jpeg",
    "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/NY-copy.jpeg",
    "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Lions.Still005.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Iceland.Still006.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Ice_Lagoon.Still004.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Glacier.Still012.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Elephants.Still007-copy.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Elephant_Water.Still008-copy.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Balloon.Still002.jpeg",
    // "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Vic_Falls.Still013.jpeg",
  ];

  const bgImgURL = useAppSelector(selectBgImgURL);

  const handleClickBg = (url) => {
    dispatch(checkInActions.setBgImgURL(url));
    setChosenBgUrl(url);
    window.location.reload();
  };

  return (
    <Row className="w-100 mt-5" justify="center" align="middle">
      <Col
        style={{
          borderRadius: "0.5rem",
          border: "1px solid white",
          boxShadow:
            "0px 3px 5px -1px rgba(0,0,0,0.2), 0px 6px 10px 0px rgba(0,0,0,0.14), 0px 1px 18px 0px rgba(0,0,0,0.12)",
        }}
        className="me-5 p-2"
        span={1}
      >
        <Space className="w-100" direction="vertical" align="center">
          <HomeOutlined
            style={{
              fontSize: "2rem",
              backgroundColor: "white",
              borderRadius: "0.3rem",
              padding: "0.3rem",
            }}
            onClick={() => navigate("/")}
          />

          <UserOutlined
            style={{
              fontSize: "2rem",
              backgroundColor: "white",
              borderRadius: "0.3rem",
              padding: "0.3rem",
            }}
            onClick={() => navigate("/profile/user-information")}
          />

          <ArrowLeftOutlined
            style={{
              fontSize: "2rem",
              backgroundColor: "white",
              borderRadius: "0.3rem",
              padding: "0.3rem",
            }}
            onClick={() => navigate(-1)}
          />
        </Space>
      </Col>

      <Col className="" span={16}>
        <Row
          style={{
            border: "1px solid white",
            boxShadow:
              "0px 11px 15px -7px rgba(0,0,0,0.2), 0px 24px 38px 3px rgba(0,0,0,0.14), 0px 9px 46px 8px rgba(0,0,0,0.12)",
            width: "100%",
            padding: "1rem",
            height: "30rem",
            borderRadius: "0.5rem",
          }}
          justify="space-between"
          align="middle"
        >
          <Col span={4}>
            <Space
              style={{
                border: "1px solid white",
                boxShadow:
                  "0px 2px 4px -1px rgba(0,0,0,0.2), 0px 4px 5px 0px rgba(0,0,0,0.14), 0px 1px 10px 0px rgba(0,0,0,0.12)",
                padding: "1rem",
                borderRadius: "0.5rem",
              }}
              direction="vertical"
            >
              {urls.map((item) => (
                <div
                  style={{
                    border: item === bgImgURL ? "3px solid #f68712" : "",
                    borderRadius: "0.8rem",
                  }}
                  onClick={() => handleClickBg(item)}
                >
                  <img
                    style={{
                      height: "6rem",
                      width: "auto",
                      borderRadius: "0.5rem",
                    }}
                    src={item}
                    alt=""
                  />
                </div>
              ))}
            </Space>
          </Col>

          <Col>
            <ReactPannellum
              // width="100%"
              // height="500px"
              id="1"
              sceneId="firstScene"
              imageSource={bgImgURL}
              config={config}
            />

            <Row className="mt-2" justify="center">
              <Space size={20}>
                <Button
                  style={{
                    width: "9rem",
                    height: "2.5rem",
                    backgroundColor: "#2C3892",
                    padding: "9px 30px 10px",
                  }}
                  type="primary"
                  onClick={handleScreenshot}
                >
                  Check In
                </Button>
              </Space>
            </Row>
          </Col>
        </Row>
      </Col>
    </Row>
  );
}
