import React, { useEffect, useState } from "react";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import ReactReadMoreReadLess from "react-read-more-read-less";

import { Typography, Row, Col, Space, Avatar, Image } from "antd";
import ButtonLikeReplyCmt from "./ButtonLikeChildReview";

const { Text } = Typography;

interface IProps {
  id: string;
  username: string;
  avatar: string;
  review_img_url: string[];
  createdAt: string;
  content: string;
  status: -1 | 0 | 1;
}

const RelyComment: React.FC<IProps> = ({
  id,
  username,
  avatar,
  review_img_url,
  createdAt,
  status,
  content,
}) => {
  return (
    <>
      <Row
        style={{ width: "90%" }}
        id={id}
        className="mt-3"
        align="top"
        gutter={10}
      >
        <Col className="mt-3" span={3}>
          <Text
            style={{
              color: "#FD665E",
              fontSize: "5rem",
              lineHeight: "5rem",
            }}
          >
            “
          </Text>
        </Col>

        <Col span={21}>
          <Row className="mb-2 mt-3" gutter={15} align="middle">
            <Col>
              <Avatar size={40} src={avatar} />
            </Col>
            <Col>
              <Space direction="vertical" size={0}>
                <Text style={{ fontSize: "0.8rem" }} strong>
                  {username}
                </Text>
                <Text
                  style={{ fontSize: "0.8rem" }}
                  className="text-secondary m-0"
                >
                  {createdAt}
                </Text>
              </Space>
            </Col>
          </Row>

          <Row style={{ fontSize: "0.8rem" }} className="mb-2 text-secondary">
            <ReactReadMoreReadLess
              charLimit={150}
              readMoreText={
                <Text style={{ color: "#5449A3" }} strong>
                  Xem thêm
                </Text>
              }
              readLessText={
                <Text style={{ color: "#5449A3" }} strong>
                  Thu gọn
                </Text>
              }
            >
              {content}
            </ReactReadMoreReadLess>
          </Row>

          <Image.PreviewGroup>
            <Space>
              {review_img_url.map((url, index) => (
                <Image
                  key={index}
                  style={{ marginRight: "1rem" }}
                  width={80}
                  height={80}
                  src={url}
                />
              ))}
            </Space>
          </Image.PreviewGroup>

          <Row className="mt-2">
            {/* <ButtonLikeReplyCmt status={status} /> */}
          </Row>
        </Col>
      </Row>
    </>
  );
};

export default RelyComment;
