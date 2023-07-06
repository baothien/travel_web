import React from "react";

import { Typography, Row, Col, Space } from "antd";

import {
  InstagramOutlined,
  TwitterOutlined,
  RedditOutlined,
  FacebookOutlined,
  YoutubeOutlined,
} from "@ant-design/icons";

const { Title, Text } = Typography;
import mainLogo from "../assets/logo/mainLogo.jpg";

const Footer: React.FC = () => {
  return (
    <Row style={{ height: "20rem" }} className="px-5" justify="center">
      <Col span={12}>
        <Space direction="vertical" size={20}>
          {/* <img height={150} src={mainLogo} alt="" /> */}
          <p style={{ fontSize: "1.2rem", fontWeight: "700" }} className="m-0">
            VỀ VIRTUAL TRAVEL
          </p>

          <div style={{ width: "30rem" }}>
            <Text>
              Một nơi tuyệt vời để khám phá các địa điểm du lịch thú vị trên
              khắp thế giới. Trang web cung cấp thông tin về các điểm đến hấp
              dẫn, các hoạt động và trải nghiệm độc đáo mà bạn có thể tham gia
              khi đi du lịch
            </Text>
          </div>

          <Space size={30}>
            <FacebookOutlined />
            <InstagramOutlined />
            <TwitterOutlined />
            <RedditOutlined />
            <YoutubeOutlined />
          </Space>

          <p style={{ fontWeight: "700", fontSize: "1rem" }}>
            Copyright © 2023 VIRTUAL TRAVEL VIETNAM.
          </p>
        </Space>
      </Col>
      <Col span={4}>
        <Space direction="vertical">
          <Title level={3}>Điều hướng</Title>
          <Text>Trang chủ</Text>
          <Text>Điểm đến</Text>
          <Text>Khám phá</Text>
          <Text>Mốc thời gian</Text>
          <Text>Đánh giá</Text>
        </Space>
      </Col>
      <Col span={4}>
        <Space direction="vertical">
          <Title level={3}>Hỗ trợ</Title>
          <Text>FAQ</Text>
          <Text>Trung tâm hỗ trợ</Text>
          <Text>Bảo mật</Text>
        </Space>
      </Col>
      <Col span={4}>
        <Space direction="vertical">
          <Title level={3}>Đối tác</Title>
          <Text>Tài trợ</Text>
          <Text>Đăng ký</Text>
        </Space>
      </Col>
    </Row>
  );
};

export default Footer;
