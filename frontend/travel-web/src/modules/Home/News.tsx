import {
  FacebookFilled,
  RedditSquareFilled,
  TwitterSquareFilled,
} from "@ant-design/icons";
import { Button, Card, Col, Row, Space, Typography } from "antd";
import React from "react";
import { Link } from "react-router-dom";

const { Title, Text } = Typography;

const News: React.FC = () => {
  return (
    <div>
      <Row
        style={{ textAlign: "center", marginTop: "5rem" }}
        className="news"
        justify="center"
      >
        <Space direction="vertical" size={30}>
          <Col>
            <Title className="m-0 fw-light" level={1}>
              CÁC TIN TỨC MỚI VỀ
            </Title>
            <Title
              type="danger"
              style={{ margin: 0, fontWeight: 100 }}
              level={1}
            >
              DU LỊCH
            </Title>
          </Col>

          <div style={{ width: "40rem" }}>
            <Title
              className="secondFont"
              type="secondary"
              style={{ margin: 0 }}
              level={5}
            >
              Cập nhập tin tức du lịch mới nhất Việt Nam và thế giới, chia sẻ
              kinh nghiệm, cẩm nang du lịch hữu ích cho tất cả mọi người
            </Title>
          </div>
        </Space>
      </Row>

      <Row
        justify="center"
        align="top"
        style={{ textAlign: "center", marginTop: "5rem", padding: "0 5rem" }}
        gutter={25}
      >
        <Col span={14}>
          <Card
            style={{ height: "41.6rem" }}
            className="border"
            hoverable
            cover={
              <img
                alt="example"
                src="https://www.icisequynhon.com/wp-content/uploads/2020/05/quynhon-binhdinh.jpg"
              />
            }
            onClick={() =>
              window.open(
                "/detail-place/b64fb4ab-581c-11ed-9c0f-0242ac1c000a",
                "_blank"
              )
            }
          >
            <Space direction="vertical">
              <Title className="m-0">Bình Định</Title>
              <Title className="m-0 fw-light">Nhất định phải đi 1 lần</Title>
              <Text type="secondary">Danh lam thắng cảnh</Text>
              <Text className="secondFont fs-6">
                Bình Định có gì? Để gợi ý cho những địa điểm vui chơi ở Bình
                Định nên đến là gì? Hãy cùng nhau khám phá những danh lam thắng
                cảnh ở Bình Định nhé!
              </Text>
              <Button type="text">XEM THÊM...</Button>
              <Space>
                <FacebookFilled
                  style={{ fontSize: "2rem", color: "#4267B2" }}
                  onClick={() =>
                    window.open("https://www.facebook.com/", "_blank")
                  }
                />

                <TwitterSquareFilled
                  style={{ fontSize: "2rem", color: "#1DA1F2" }}
                  onClick={() =>
                    window.open("https://twitter.com/elonmusk", "_blank")
                  }
                />

                <RedditSquareFilled
                  style={{ fontSize: "2rem", color: "#FF4500" }}
                  onClick={() =>
                    window.open("https://www.reddit.com/", "_blank")
                  }
                />
              </Space>
            </Space>
          </Card>
        </Col>

        <Col span={6}>
          <Space style={{ height: "40rem" }} direction="vertical" size={20}>
            <Card
              className="border p-0"
              hoverable
              cover={
                <img
                  alt="example"
                  src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/09/blog-post-h6-img2.jpg"
                />
              }
              onClick={() =>
                window.open(
                  "/detail-place/90a1727e-db80-11ed-9443-acde48001122",
                  "_blank"
                )
              }
            >
              <Text
                style={{ fontSize: "1rem", fontWeight: "700" }}
                className="m-0"
              >
                Nghĩa trang Nhơn Hải
              </Text>
            </Card>

            <Card
              hoverable
              cover={
                <img
                  alt="example"
                  src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/09/blog-post-h6-img2.jpg"
                />
              }
              onClick={() =>
                window.open(
                  "/detail-place/30891b14-581e-11ed-9c0f-0242ac1c000a",
                  "_blank"
                )
              }
            >
              <Text
                style={{ fontSize: "1rem", fontWeight: "700" }}
                className="m-0"
              >
                Khu Du Lịch Cửa Biển
              </Text>
            </Card>
          </Space>
        </Col>
      </Row>
    </div>
  );
};

export default News;
