import * as React from "react";
import { NotifyItem } from "./NotifyItem";

import { Badge, Button, Popover, Radio } from "antd";
import { Notify } from "../../models/notify";
import notifyApi from "../../apis/notifyApi";
import { Link } from "react-router-dom";
import { BellOutlined } from "@ant-design/icons";
import { useAppDispatch, useAppSelector } from "../../hooks";
import { notifyActions, selectNotifyCount } from "./notifySlice";

// export interface INotificationProps {
//   width: string;
// }

export function NotifyHeader() {
  const dispatch = useAppDispatch();

  const notifyCount = useAppSelector(selectNotifyCount);

  const [notifyList, setNotifyList] = React.useState<Notify[]>([]);

  const [valueRadio, setValueRadio] = React.useState("all");

  const handleModeChange = (event: any) => {
    setValueRadio(event.target.value);
  };

  React.useEffect(() => {
    const paginate = {
      page: 1,
      limit: 5,
    };

    notifyApi
      .getNotifyList(paginate)
      .then((res) => {
        notifyApi
          .getNotifyCount()
          .then((res) => {
            dispatch(notifyActions.setCount(res.data));
          })
          .catch((err) => {
            console.log(err);
          });

        if (valueRadio === "all") setNotifyList(res.data.rows);
        else {
          setNotifyList(res.data.rows.filter((item) => item.is_read === false));
        }
      })
      .catch((err) => console.log(err));
  }, [valueRadio, notifyCount]);

  return (
    <Popover
      placement="bottom"
      content={
        notifyList.length === 0 ? (
          <>Chưa có thông báo</>
        ) : (
          <div style={{ width: "20rem" }}>
            <Radio.Group
              onChange={handleModeChange}
              value={valueRadio}
              style={{ marginBottom: 8 }}
            >
              <Radio.Button value="all">Tất cả</Radio.Button>
              <Radio.Button value="unseen">Chưa xem</Radio.Button>
            </Radio.Group>

            {notifyList.map((item) => (
              <NotifyItem
                key={item.id}
                id={item.child_id}
                user_name={item.from_user.user_name}
                avatar={item.from_user.avatar}
                title={
                  item.title.length < 70
                    ? item.title
                    : item.title.slice(0, 60).concat("...")
                }
                isRead={item.is_read}
                date={
                  item.created_at.slice(0, 10) +
                  " " +
                  item.created_at.slice(11, 19)
                }
                destinationId={item.destination_id}
              />
            ))}

            <Link
              style={{
                display: "flex",
                justifyContent: "center",
              }}
              to="/profile/notification"
            >
              <Button style={{}} className="" type="link">
                Xem thêm
              </Button>
            </Link>
          </div>
        )
      }
      title="Thông báo"
      trigger="click"
    >
      <Badge count={notifyCount}>
        <Button
          shape="circle"
          type="text"
          icon={<BellOutlined style={{ fontSize: "1.2rem" }} />}
        />
      </Badge>
    </Popover>
  );
}
