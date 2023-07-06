import { Row, Space, Typography } from "antd";
import * as React from "react";

const { Text } = Typography;

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface ICardReviewProps {
  imgUrl: string;
  description: string;
  authorName: string;
}

export function CardReview({
  imgUrl,
  description,
  authorName,
}: ICardReviewProps) {
  return (
    <div className="w-100">
      <img style={{ width: "13rem" }} className="" src={imgUrl} alt="" />

      <Row style={{ width: "12rem" }} className="p-3" justify="center">
        <Text>
          {description}{" "}
          <span style={{ fontWeight: "700", fontSize: "1rem" }}>
            - {authorName}
          </span>
        </Text>
      </Row>
    </div>
  );
}
