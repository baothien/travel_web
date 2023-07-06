import {
  DislikeFilled,
  DislikeOutlined,
  LikeFilled,
  LikeOutlined,
} from "@ant-design/icons";
import { Button, Space } from "antd";
import React, { useState } from "react";

interface IProps {
  status: -1 | 0 | 1;
}

const ButtonLikeReplyCmt: React.FC<IProps> = ({ status }) => {
  const [newStatus, setNewStatus] = useState(status);

  const handleLike = () => {
    if (newStatus === 1) setNewStatus(0);
    else setNewStatus(1);
  };

  const handleDislike = () => {
    if (newStatus === -1) setNewStatus(0);
    else setNewStatus(-1);
  };
  return (
    <Space>
      {newStatus === 1 ? (
        <Button
          size="small"
          shape="round"
          onClick={handleLike}
          icon={<LikeFilled style={{ color: "#1677FF" }} />}
        ></Button>
      ) : (
        <Button
          size="small"
          shape="round"
          onClick={handleLike}
          icon={<LikeOutlined style={{ color: "grey" }} />}
        ></Button>
      )}

      {newStatus === -1 ? (
        <Button
          size="small"
          shape="round"
          danger
          onClick={handleDislike}
          icon={<DislikeFilled />}
        ></Button>
      ) : (
        <Button
          size="small"
          shape="round"
          onClick={handleDislike}
          icon={<DislikeOutlined style={{ color: "grey" }} />}
        ></Button>
      )}
    </Space>
  );
};

export default ButtonLikeReplyCmt;
