import { ArrowRightOutlined, EnvironmentFilled } from "@ant-design/icons";
import { Button, Col, Row, Space, Typography } from "antd";
import * as React from "react";

export interface IFavoriteItemProps {
  id: string;
  thumbnail: string;
  name: string;
  placeType: string;
}

const { Title, Text } = Typography;

export function FavoriteItem({
  id,
  thumbnail,
  name,
  placeType,
}: IFavoriteItemProps) {
  const openDetail = () => {
    window.open(`/detail-place/${id}`);
  };

  return (
    <Row className="rounded" align="middle">
      <Col className="position-relative">
        <img
          style={{
            borderRadius: "0.5rem",
            filter: "brightness(80%)",
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
            <Title className="text-white m-0" level={5}>
              {name}
            </Title>
            {/* <Title className="text-white" level={5}>
              {placeType}
            </Title> */}

            <Space>
              <Button
                className="text-white border-white"
                size="middle"
                type="text"
                shape="circle"
                icon={<ArrowRightOutlined />}
                onClick={openDetail}
              />

              <Button
                className="text-white border-white"
                size="middle"
                type="text"
                shape="circle"
                icon={<EnvironmentFilled />}
              />
            </Space>
          </Space>
        </div>
      </Col>
    </Row>
  );
}
