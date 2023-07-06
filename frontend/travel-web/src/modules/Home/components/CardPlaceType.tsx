/* eslint-disable @typescript-eslint/no-empty-interface */
import { Row, Space, Typography } from "antd";
import * as React from "react";

const { Title, Text } = Typography;

export interface ICardPlaceTypeProps {
  imgURL: string;
  title: string;
  description: string;
}

export function CardPlaceType({
  title,
  description,
  imgURL,
}: ICardPlaceTypeProps) {
  return (
    <Row justify="center">
      <Space style={{}} direction="vertical" align="center">
        <img src={imgURL} />

        <Text style={{ fontSize: "1rem" }} strong>
          {title}
        </Text>
        <Text style={{ fontSize: "0.8rem" }} className="text-secondary">
          {description}
        </Text>
      </Space>
    </Row>
  );
}
