import { Row, Col, Avatar, Typography, Divider } from "antd";
import * as React from "react";
import notifyApi from "../../apis/notifyApi";
import { useAppDispatch } from "../../hooks";
import { notifyActions } from "./notifySlice";

export interface INotifyItemProps {
  id: string;
  user_name: string;
  avatar: string;
  title: string;
  isRead: boolean;
  date: string;
  destinationId: string;
  selectedNotify?: string;
}

export function NotifyItem({
  id,
  user_name,
  avatar,
  title,
  isRead,
  date,
  destinationId,
}: INotifyItemProps) {
  const dispatch = useAppDispatch();

  const [is_read, set_is_read] = React.useState(isRead);

  const handleReadNotify = () => {
    dispatch(notifyActions.setCmtNotifiedId(id));

    if (!is_read) {
      notifyApi
        .patchReadNotify(id)
        .then((res) => {
          dispatch(notifyActions.decreaseCount());

          set_is_read(true);
        })
        .catch((err) => {
          console.log("err", err);
        });
    }

    window.open(`/detail-place/${destinationId}`, "_blank");
  };

  return (
    <div>
      <Row
        style={{
          width: "100%",
          margin: "0.5rem 0",
        }}
        className="p-1 rounded"
        justify="space-between"
        onClick={handleReadNotify}
      >
        <Col span={4}>
          <Avatar style={{ width: "3rem", height: "3rem" }} src={avatar} />
        </Col>
        <Col span={19}>
          <Typography>
            <span style={{ fontWeight: "700" }}>{user_name}</span> {title}
          </Typography>
          <Typography style={{ fontSize: "0.8rem" }} className="text-muted">
            {date}
          </Typography>
        </Col>
        <Col className="align-items-center d-flex" span={1}>
          {is_read ? (
            <></>
          ) : (
            <div
              style={{
                width: "0.8rem",
                height: "0.8rem",
              }}
              className="rounded-circle bg-primary"
            ></div>
          )}
        </Col>
      </Row>
      <Divider style={{ backgroundColor: "#b8b8b8" }} className="m-0" />
    </div>
  );
}
