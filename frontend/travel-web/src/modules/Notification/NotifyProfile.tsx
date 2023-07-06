import * as React from "react";
import { NotifyItem } from "./NotifyItem";

import { Pagination, PaginationProps, Radio } from "antd";
import { Notify } from "../../models/notify";
import notifyApi from "../../apis/notifyApi";

// export interface INotificationProps {
//   width: string;
// }

export function NotifyProfile() {
  const [notifyList, setNotifyList] = React.useState<Notify[]>([]);

  const [valueRadio, setValueRadio] = React.useState("all");

  const [page, setPage] = React.useState<number>(1);

  const handleModeChange = (event: any) => {
    setValueRadio(event.target.value);
  };

  const onChangePage: PaginationProps["onChange"] = (page) => {
    setPage(page);
  };

  React.useEffect(() => {
    const paginate = {
      page: page,
      limit: 6,
    };

    notifyApi
      .getNotifyList(paginate)
      .then((res) => {
        if (valueRadio === "all") {
          setNotifyList(res.data.rows);
        } else {
          setNotifyList(res.data.rows.filter((item) => item.is_read === false));
        }
      })
      .catch((err) => console.log(err));
  }, [valueRadio, page]);

  return (
    <>
      {notifyList.length ? (
        <div style={{ width: "45rem" }}>
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
              id={item.id}
              user_name={item.from_user.user_name}
              avatar={item.from_user.avatar}
              title={item.title}
              isRead={item.is_read}
              date={item.created_at.slice(0, 19)}
              destinationId={item.destination_id}
            />
          ))}

          <Pagination
            className="w-100 mt-4"
            onChange={onChangePage}
            total={notifyList.length}
            showSizeChanger={false}
            hideOnSinglePage={true}
          />
        </div>
      ) : (
        <div>
          <Radio.Group
            onChange={handleModeChange}
            value={valueRadio}
            style={{ marginBottom: 8 }}
          >
            <Radio.Button value="all">Tất cả</Radio.Button>
            <Radio.Button value="unseen">Chưa xem</Radio.Button>
          </Radio.Group>

          <div className="d-flex justify-content-center h3">
            Chưa có thông báo
          </div>
        </div>
      )}
    </>
  );
}
