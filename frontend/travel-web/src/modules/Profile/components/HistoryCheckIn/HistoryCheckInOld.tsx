import { Col, Row, Space } from "antd";
import React, { useState, useEffect } from "react";
import { HistoryCheckInItem } from "./HistoryCheckInItem";
import checkInApi from "../../../../apis/checkInApi";

const checkItems = [
  {
    imgURL:
      "https://media.vov.vn/sites/default/files/styles/large/public/2022-05/a4b4c128-53d9-4073-8c1d-98f677e40691_802064ff.jpg",
    title: "Khánh Hoà",
  },
  {
    imgURL:
      "https://images.unsplash.com/photo-1612441804231-77a36b284856?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8bW91bnRhaW4lMjBsYW5kc2NhcGV8ZW58MHx8MHx8&w=1000&q=80",
    title: "Bình Định",
  },
  {
    imgURL: "https://wallpapercave.com/wp/wp5336739.jpg",
    title: "Quảng Ngãi",
  },
  {
    imgURL: "https://wallpaperaccess.com/full/328061.jpg",
    title: "Thanh Hoá",
  },
  {
    imgURL:
      "https://e0.pxfuel.com/wallpapers/979/874/desktop-wallpaper-page-11-of-york-for-your-or-mobile-screen-new-york-landscape.jpg",
    title: "Phú Yên",
  },
];

const HistoryCheckIn: React.FC = () => {
  const [checkInItems, setCheckInItems] = useState<any>([]);

  const checkInImgUrls = [
    "https://gw.alipayobjects.com/zos/antfincdn/LlvErxo8H9/photo-1503185912284-5271ff81b9a8.webp",
    "https://gw.alipayobjects.com/zos/antfincdn/cV16ZqzMjW/photo-1473091540282-9b846e7965e3.webp",
    "https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp",
  ];

  useEffect(() => {
    const getData = async () => {
      await checkInApi.getCheckInList().then((res) => {
        const imgUrls = res.data.rows.map((item) => item.url);
        // setCheckInItems(res.data.rows);
        setCheckInItems((prev) => [
          ...prev,
          {
            title: res.data.rows[0].place.name,
            thumbnail: res.data.rows[0].place.thumbnail,
            checkInImgUrls: imgUrls,
          },
        ]);
      });
    };

    getData();
  }, []);
  return (
    <Row className="w-100" gutter={[16, 16]} justify="space-between">
      {checkInItems.map((item) => (
        <Col>
          <HistoryCheckInItem
            title={item.title}
            thumbnail={item.thumbnail}
            checkInImgUrls={item.checkInImgUrls}
          />
        </Col>
      ))}

      {checkItems.map((item) => (
        <Col>
          <HistoryCheckInItem
            thumbnail={item.imgURL}
            title={item.title}
            checkInImgUrls={checkInImgUrls}
          />
        </Col>
      ))}
    </Row>
  );
};

export default HistoryCheckIn;
