import React, { useState } from "react";
import { Typography, Row, Col, Space, Modal } from "antd";

import {
  StarOutlined,
  MessageOutlined,
  EnvironmentOutlined,
} from "@ant-design/icons";
import { useParams } from "react-router-dom";
import reviewApi from "../../apis/reviewApi";

const { Title, Text } = Typography;

const Summary: React.FC = () => {
  const { id } = useParams();
  const [countReview, setCountReview] = useState(0);

  React.useEffect(() => {
    if (id) {
      reviewApi
        .getReviews(id, { page: 1, limit: 1 })
        .then((res) => {
          setCountReview(res.data.total_rows);
        })
        .catch((err) => console.log(err));
    }
  }, []);
  return (
    <div>
      <Space
        className="trapezoid pt-5"
        direction="vertical"
        align="center"
        size={20}
      >
        <Row className="mt-3" justify="center">
          <Space direction="vertical" align="center" size={3}>
            <EnvironmentOutlined className="text-white fs-3" />
            <Text className="text-white fw-lighter">Lượt check in</Text>
          </Space>
        </Row>

        <Row justify="center">
          <Space direction="vertical" align="center" size={3}>
            <MessageOutlined className="text-white fs-3" />
            <Text className="text-white fw-lighter">
              {countReview} lượt đánh giá
            </Text>
          </Space>
        </Row>

        <Row justify="center">
          <Space direction="vertical" align="center" size={3}>
            <StarOutlined className="text-white fs-3" />
            <Text className="text-white fw-lighter">Số sao</Text>
          </Space>
        </Row>
      </Space>
    </div>
  );
};

export default Summary;
