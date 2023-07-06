import { Button, Space } from "antd";
import React, { useState } from "react";
import { toast } from "react-toastify";
import { useAppSelector } from "../../../../hooks";
import { selectUser } from "../../../Authentication/authSlice";
import {
  DislikeFilled,
  DislikeOutlined,
  LikeFilled,
  LikeOutlined,
} from "@ant-design/icons";

interface IProps {
  status: -1 | 0 | 1;
}

const ButtonLike: React.FC<IProps> = ({ status }) => {
  const [newStatus, setNewStatus] = useState(status);

  const user = useAppSelector(selectUser);

  const handleLike = () => {
    if (user) {
      if (newStatus === 1) setNewStatus(0);
      else setNewStatus(1);
    } else {
      toast.info("Cần đăng nhập để thực hiện chức năng này");
    }
  };

  const handleDislike = () => {
    if (user) {
      if (newStatus === -1) setNewStatus(0);
      else setNewStatus(-1);
    } else {
      toast.info("Cần đăng nhập để thực hiện chứa năng này");
    }
  };

  return (
    // <Space direction="vertical">
    //   {newStatus === 1 ? (
    //     <Button
    //       style={{ width: "5.5rem", marginTop: "2rem" }}
    //       type="primary"
    //       size="small"
    //       shape="round"
    //       onClick={handleLike}
    //     >
    //       Like
    //     </Button>
    //   ) : (
    //     <Button
    //       style={{ width: "5.5rem", marginTop: "2rem" }}
    //       type="primary"
    //       size="small"
    //       shape="round"
    //       ghost
    //       onClick={handleLike}
    //     >
    //       Like
    //     </Button>
    //   )}
    //   {newStatus === -1 ? (
    //     <Button
    //       style={{ width: "5.5rem" }}
    //       type="primary"
    //       size="small"
    //       shape="round"
    //       danger
    //       onClick={handleDislike}
    //     >
    //       Dislike
    //     </Button>
    //   ) : (
    //     <Button
    //       style={{ width: "5.5rem" }}
    //       type="primary"
    //       size="small"
    //       shape="round"
    //       danger
    //       ghost
    //       onClick={handleDislike}
    //     >
    //       Dislike
    //     </Button>
    //   )}
    // </Space>
    <Space direction="vertical">
      {newStatus === 1 ? (
        <Button
          style={{ width: "4rem" }}
          size="small"
          shape="round"
          onClick={handleLike}
          icon={<LikeFilled style={{ color: "#1677FF" }} />}
        ></Button>
      ) : (
        <Button
          style={{ width: "4rem" }}
          size="small"
          shape="round"
          onClick={handleLike}
          icon={<LikeOutlined style={{ color: "grey" }} />}
        ></Button>
      )}

      {newStatus === -1 ? (
        <Button
          style={{ width: "4rem" }}
          size="small"
          shape="round"
          danger
          onClick={handleDislike}
          icon={<DislikeFilled />}
        ></Button>
      ) : (
        <Button
          style={{ width: "4rem" }}
          size="small"
          shape="round"
          onClick={handleDislike}
          icon={<DislikeOutlined style={{ color: "grey" }} />}
        ></Button>
      )}
    </Space>
  );
};

export default ButtonLike;
