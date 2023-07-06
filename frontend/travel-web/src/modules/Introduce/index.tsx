import * as React from "react";
import { Footer, Header } from "../../components";
import { Button, Col, Row, Space, Typography } from "antd";
import {
  FacebookFilled,
  InstagramFilled,
  TwitterCircleFilled,
} from "@ant-design/icons";

import aboutIcon from "../../assets/about-icon-1.png";
import bridgeBg from "../../assets/h4-slider-img-2.png";

import book1 from "../../assets/book1.png";
import book2 from "../../assets/book2.png";
import { useNavigate } from "react-router-dom";

const { Title, Text } = Typography;

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface IIntroduceProps {}

export default function Introduce(props: IIntroduceProps) {
  const navigate = useNavigate();
  return (
    <div>
      <Header />

      <Row style={{ margin: "8rem 0rem" }} className="" justify="center">
        <Row
          style={{ width: "80%" }}
          className="mb-5"
          justify="space-between"
          align="middle"
        >
          <Col span={11}>
            <Space direction="vertical" size={10}>
              <Text style={{ fontSize: "1.7rem" }} className="m-0">
                LET’S GO TRAVEL :)
              </Text>

              <Text style={{ fontSize: "0.9rem" }}>
                Chào mừng bạn đến với trang web du lịch hàng đầu, nơi bạn có thể
                khám phá những điểm đến tuyệt vời trên khắp thế giới. Tận hưởng
                trải nghiệm du lịch độc đáo thông qua thông tin chi tiết, đánh
                giá chất lượng và đặt chỗ nhanh chóng.
              </Text>

              <Text style={{ fontSize: "0.9rem" }}>
                Với giao diện thân thiện và nội dung phong phú, chúng tôi sẽ
                đồng hành cùng bạn trong mọi chuyến hành trình đáng nhớ. 😎
              </Text>

              <Space>
                <div>
                  <FacebookFilled
                    style={{ fontSize: "1.5rem", color: "#3B5998" }}
                  />
                </div>
                <div>
                  <InstagramFilled
                    style={{ fontSize: "1.5rem", color: "#CB2027" }}
                  />
                </div>
                <div>
                  <TwitterCircleFilled
                    style={{ fontSize: "1.5rem", color: "#55ACEE" }}
                  />
                </div>
              </Space>
            </Space>
          </Col>

          <Col>
            <Row style={{ position: "relative" }} className="mb-4">
              <div className="border p-3">
                <img
                  style={{ width: "30rem" }}
                  src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/09/about-img-video-1.jpg"
                  alt=""
                />
              </div>
            </Row>

            {/* <Row justify="space-between">
              <Col>
                <img src={aboutIcon} alt="" />
              </Col>
              <Col span={20} className="">
                <Text style={{ fontSize: "1rem", fontWeight: "800" }}>
                  Morbi accumsan ipsum velit. nam nec tellus a odio tincidunt.
                </Text>
              </Col>
            </Row> */}
          </Col>
        </Row>

        <Row
          style={{ width: "100%", height: "30rem", backgroundColor: "#F7F2EE" }}
          justify="start"
          align="middle"
        >
          <Col style={{ position: "relative", marginRight: "12rem" }} span={6}>
            <img style={{ height: "25rem" }} src={bridgeBg} alt="" />

            <img
              style={{
                width: "15rem",
                position: "absolute",
                top: "20%",
                left: "30%",
                zIndex: "1",
              }}
              src={book1}
              alt=""
            />
            <img
              style={{
                width: "15rem",
                position: "absolute",
                top: "20%",
                left: "72%",
                zIndex: "0",
              }}
              src={book2}
              alt=""
            />
          </Col>

          <Col span={10}>
            <Row>
              {" "}
              <Text style={{ fontSize: "2.2rem" }}>ĐĂNG KÝ NGAY!</Text>
            </Row>
            {/* <Row style={{ width: "40rem" }}>
              <Text style={{ fontSize: "0.9rem" }}>
                See how i helped real readers plan, save and go on the trip of a
                lifetime. And get actionable steps you can use to travel
                anywhere - no matter your income or where you're from.
              </Text>
            </Row> */}
            <Row className="mt-3">
              <button
                style={{
                  width: "8rem",
                  height: "2.5rem",
                  padding: "0.5rem",
                  backgroundColor: "black",
                  color: "white",
                  border: "none",
                }}
                onClick={() => navigate("/register")}
              >
                ĐĂNG KÝ
              </button>
            </Row>
          </Col>
        </Row>

        <Row style={{ marginTop: "6rem" }} justify="center">
          <Space className="w-100" direction="vertical" size={10}>
            <Row justify="center">
              <div>
                <Text style={{ fontSize: "2.2rem" }}>CÁC ĐÁNH GIÁ NỔI BẬT</Text>
              </div>
            </Row>

            {/* <Row className="w-100" justify="center">
              <Text style={{ fontSize: "1rem" }} type="secondary">
                See how I’ve helped others save money and plan trips of their
                lifetime
              </Text>
            </Row> */}
          </Space>
        </Row>

        <Row
          className="position-relative w-100"
          style={{
            padding: "0 15rem",
            marginBottom: "3rem",
            marginTop: "2rem",
          }}
          justify="end"
        >
          <Space direction="vertical">
            <Space>
              <Title
                style={{
                  position: "absolute",
                  fontSize: "2rem",
                  top: "6rem",
                  right: "8rem",
                }}
              >
                Sara & Tom
              </Title>
              <img
                width={400}
                height={400}
                style={{
                  borderRadius: "100%",
                }}
                alt=""
                src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/08/t-3.jpg"
              />
            </Space>
            <div
              style={{
                position: "absolute",
                top: "9rem",
                left: "20rem",
                width: "22rem",
                textAlign: "center",
                marginTop: "3rem",
              }}
            >
              <Title level={5}>
                Trang web du lịch này thực sự xuất sắc! Nó cung cấp thông tin
                đáng tin cậy về các điểm đến, đặc điểm nổi bật và hoạt động thú
                vị.
              </Title>
            </div>
          </Space>
        </Row>

        <Row
          className="position-relative"
          style={{
            padding: "0 15rem",
            marginBottom: "5rem",
          }}
          justify="end"
        >
          <Space direction="vertical">
            <Space>
              <Title
                style={{
                  position: "absolute",
                  fontSize: "3rem",
                  top: "6rem",
                  left: "8rem",
                }}
              >
                Laetitia
              </Title>
              <img
                width={400}
                height={400}
                style={{
                  borderRadius: "100%",
                }}
                alt=""
                src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/08/t-4.jpg"
              />
            </Space>
            <div
              style={{
                position: "absolute",
                right: "-7rem",
                top: "9rem",
                width: "22rem",
                textAlign: "center",
                marginTop: "3rem",
              }}
            >
              <Title level={5}>
                Đánh giá cao cho trang web du lịch này! Nó cung cấp thông tin
                chi tiết về các điểm đến, khuyến nghị về nhà hàng, khách sạn và
                hoạt động vui chơi. Tôi khuyến nghị trang web này cho mọi người
                muốn có một kỳ nghỉ hoàn hảo và thú vị.
              </Title>
            </div>
          </Space>
        </Row>
      </Row>

      <Footer />
    </div>
  );
}

const data = [
  {
    imgUrl:
      "https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/08/t-3.jpg",
    description:
      "Morbi aliquet tempus tempor. Nunc vitae semper mauris. Nam sollicitudin risus dui, molestie congue massa ultricies eu. Duis pharetra, ligula in molestie congue mgravida.",
    authorName: "Sara & Tom",
  },
  {
    imgUrl:
      "https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/08/t-4.jpg",
    description:
      "Aenean eget nulla sagittis justo cursus porta at vitae tellus. Cras suscipit interdum sapien, ut consectetur libero volutpat sit amet. Sed dictum efficitur neque vitae placerat. Fusce gravida ultricies metus non luctus, vestibulum vel.",
    authorName: "Laetitia",
  },
];
