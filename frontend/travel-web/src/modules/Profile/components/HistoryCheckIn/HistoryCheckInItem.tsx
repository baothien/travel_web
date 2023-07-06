/* eslint-disable @typescript-eslint/no-empty-interface */
import { Button, Col, Row, Space, Typography, Image } from "antd";
import * as React from "react";

const { Title, Text } = Typography;

import { ArrowRightOutlined, PictureFilled } from "@ant-design/icons";

export interface IHistoryCheckInItemProps {
  thumbnail: string;
  title: string;
  checkInImgUrls: string[];
}

export function HistoryCheckInItem({
  thumbnail,
  title,
  checkInImgUrls,
}: IHistoryCheckInItemProps) {
  const [visible, setVisible] = React.useState(false);

  return (
    <Row className="rounded" align="middle">
      <Col className="position-relative">
        <img
          style={{
            borderRadius: "0.5rem",
            filter: "brightness(75%)",
            height: "14rem",
          }}
          alt=""
          src={thumbnail}
          width="100%"
        />

        <div
          style={{ position: "absolute", bottom: 0, left: 10 }}
          className="p-2"
        >
          <Space direction="vertical">
            <Title style={{ color: "white" }} className="m-0" level={5}>
              {title}
            </Title>
            <Space>
              <Button
                className="text-white border-white"
                size="middle"
                type="text"
                shape="circle"
                icon={<PictureFilled />}
                onClick={() => setVisible(true)}
              />

              <Button
                className="text-white border-white"
                size="middle"
                type="text"
                shape="circle"
                icon={<ArrowRightOutlined />}
                onClick={() => window.open(`/detail-place/`)}
              />
            </Space>
          </Space>
        </div>
      </Col>

      <div style={{ display: "none" }}>
        <Image.PreviewGroup
          preview={{ visible, onVisibleChange: (vis) => setVisible(vis) }}
        >
          {checkInImgUrls.map((item, index) => (
            <Image key={index} src={item} />
          ))}
        </Image.PreviewGroup>
      </div>
    </Row>
  );
}
