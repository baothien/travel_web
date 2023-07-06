import {
  ArrowLeftOutlined,
  ArrowRightOutlined,
  ArrowsAltOutlined,
  RightOutlined,
} from "@ant-design/icons";
import "./Home.scss";
import { Button, Col, Row, Space, Typography } from "antd";
import React, { useState } from "react";
import { Link } from "react-router-dom";

import slider1 from "../../assets/img/slider-hp-1.jpg";
import slider2 from "../../assets/img/slider-hp-2.jpg";
import slider3 from "../../assets/img/slider-hp-3.jpg";

import bgLogoText from "../../assets/img/bg-logo-text.png";

import logoText from "../../assets/img/logo-text-slider.png";

import textSlider from "../../assets/img/textslider.jpg";
import arrow from "../../assets/img/arrow.jpg";

const { Text, Title } = Typography;

const Slider: React.FC = () => {
  const [mainSliderIndex, setMainSliderIndex] = useState(0);

  const [zoomInfo, setZoomInfo] = useState(true);

  const slider = [
    {
      name: "Nha Trang",
      thumbnail: slider1,
      like: 1,
      cmt: 2,
      description:
        "Nha Trang: Bãi biển tuyệt đẹp, du lịch thú vị, hấp dẫn.",
    },
    {
      name: "Thanh Hoá",
      thumbnail: slider2,
      like: 3,
      cmt: 4,
      description:
        "Vùng đất đa dạng, từ núi rừng đến biển cả, với di sản lịch sử, thiên nhiên và văn hóa độc đáo.",
    },
    {
      name: "Bình Định",
      thumbnail: slider3,
      like: 5,
      cmt: 6,
      description:
        "Bờ biển dài, thiên nhiên hoang sơ, di sản văn hóa độc đáo.",
    },
  ];

  const handleRightSlider = () => {
    if (mainSliderIndex === slider.length - 1) {
      setMainSliderIndex(0);
    } else {
      setMainSliderIndex((prev) => prev + 1);
    }
  };

  return (
    <Row style={{ marginTop: "6.5rem" }} justify="space-between" align="middle">
      <Col>
        <Space direction="vertical" size={10} align="end">
          <img style={{ width: "24rem" }} src={textSlider} alt="" />

          <Space>
            <img
              style={{ width: "8rem", height: "2.2rem" }}
              src={arrow}
              alt=""
            />

            <Link to="./map">
              <Button
                style={{
                  backgroundColor: "#2C3892",
                  borderRadius: "1.2rem",
                  color: "white",
                }}
                className="explore-btn rounded-0 mt-1"
              >
                TRẢI NGHIỆM NGAY
              </Button>
            </Link>
          </Space>
        </Space>
      </Col>

      <Col>
        <Row className="position-relative " justify="center" align="middle">
          <Col className="position-relative">
            <img
              style={{
                boxShadow:
                  "0 20px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23)",
              }}
              width={800}
              height={500}
              src={slider[mainSliderIndex].thumbnail}
              alt=""
            />

            {zoomInfo ? (
              <div
                style={{
                  width: "65%",
                  height: "50%",
                  borderRadius: "0.2rem",
                  backgroundColor: "white",
                }}
                className="position-absolute bottom-0 left-0"
              >
                <Row justify="space-between">
                  <Col style={{ padding: "0.8rem 0 0 2rem" }} span={22}>
                    <Title level={2}>{slider[mainSliderIndex].name}</Title>

                    <Row style={{}}>
                      <Text style={{ fontSize: "0.9rem" }}>
                        <span style={{ fontWeight: "900", fontSize: "1rem" }}>
                          |
                        </span>{" "}
                        {slider[mainSliderIndex].cmt} BÌNH LUẬN{" "}
                        <span style={{ fontWeight: "900", fontSize: "1rem" }}>
                          |
                        </span>{" "}
                        {slider[mainSliderIndex].like} LƯỢT THÍCH
                      </Text>
                    </Row>

                    <Row style={{ padding: "1rem 0 0 0" }}>
                      <Text style={{ fontWeight: "600", fontSize: "0.8rem" }}>
                        {slider[mainSliderIndex].description}
                      </Text>
                    </Row>

                    <Row style={{ padding: "1rem 0 0 0" }}>
                      <Link to="./map">
                        <button
                          style={{
                            border: "1px solid #c6c6c6",
                            fontWeight: "400",
                            padding: "9px 30px 10px",
                            color: "black",
                          }}
                        >
                          KHÁM PHÁ NGAY
                        </button>
                      </Link>
                    </Row>
                  </Col>

                  <Col>
                    <ArrowsAltOutlined
                      style={{
                        width: "100%",
                        fontSize: "2.5rem",
                      }}
                      onClick={() => setZoomInfo((prev) => !prev)}
                    />
                  </Col>
                </Row>
              </div>
            ) : (
              <div
                style={{
                  width: "36%",
                  height: "25%",
                  borderRadius: "0.2rem",
                  backgroundColor: "white",
                }}
                className="position-absolute bottom-0 left-0"
              >
                <Row justify="space-between">
                  <Col style={{ padding: "1rem 0 0 3rem" }}>
                    <Title level={2}>{slider[mainSliderIndex].name}</Title>
                  </Col>

                  <Col className="p-1">
                    <ArrowsAltOutlined
                      style={{ width: "100%", fontSize: "2rem" }}
                      onClick={() => setZoomInfo((prev) => !prev)}
                    />
                  </Col>
                </Row>

                <Row style={{ paddingLeft: "3rem" }}>
                  <Text>
                    <span style={{ fontWeight: "900", fontSize: "1rem" }}>
                      |
                    </span>{" "}
                    {slider[mainSliderIndex].cmt} BÌNH LUẬN{" "}
                    <span style={{ fontWeight: "900", fontSize: "1rem" }}>
                      |
                    </span>{" "}
                    {slider[mainSliderIndex].like} LƯỢT THÍCH
                  </Text>
                </Row>
              </div>
            )}
          </Col>

          <Col style={{ position: "absolute", right: "1rem" }}>
            <RightOutlined
              style={{ fontSize: "3rem" }}
              className="text-white"
              onClick={handleRightSlider}
            />
          </Col>
        </Row>
      </Col>
    </Row>
  );
};

export default Slider;
