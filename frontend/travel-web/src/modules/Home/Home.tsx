import React, { useEffect, useState } from "react";
import "./Home.scss";

import { Header, Footer } from "../../components";

import Slider from "./Slider";

import Introduce from "./Introduce";

import News from "./News";

import { Typography, Row, Card, Space, Button, Divider } from "antd";
import apiPlaces from "../../apis/placesApi";
import { Place } from "../../models/place";

import { placeType } from "../../assets/constants/placeType";

const { Title } = Typography;

import { CardPlaceType } from "./components/CardPlaceType";

const Homepage: React.FC = () => {
  const [suggestPlace, setSuggestPlace] = useState<Place[]>([]);

  useEffect(() => {
    apiPlaces
      .getPlaces()
      .then((res) => {
        setSuggestPlace(res.data.slice(0, 4));
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);
  return (
    <div className="container-home">
      <Header />

      <div className="px-5">
        <Slider />

        <Row
          style={{ marginTop: "5rem", padding: "0 3rem" }}
          justify="space-between"
          align="middle"
        >
          <div
            style={{ display: "flex", justifyContent: "center" }}
            className="w-100 mb-5"
          >
            <Title level={3}>XU HƯỚNG DU LỊCH</Title>
          </div>

          {placeType.map((item) => (
            <CardPlaceType
              title={item.title}
              description={item.description}
              imgURL={item.imgURL}
            />
          ))}
        </Row>

        <Row style={{ marginTop: "5rem" }} justify="center">
          <Title className="m-0" level={1}>
            ĐIỂM ĐẾN NỔI BẬT
          </Title>
        </Row>

        <Row style={{ marginTop: "2rem" }} justify="center">
          <Space style={{ alignItems: "flex-start" }} size={15}>
            {suggestPlace.map((item, index) => (
              <Card
                key={index}
                style={{ width: 250, textAlign: "center" }}
                hoverable
                cover={
                  <img
                    width={250}
                    height={150}
                    alt="example"
                    src={item.thumbnail}
                  />
                }
              >
                <Space size="small" direction="vertical">
                  {/* <img height={150} alt="example" src={peruImg} /> */}
                  <Title
                    style={{ minHeight: "4rem" }}
                    className="m-0"
                    level={3}
                  >
                    {item.name}
                  </Title>
                  <Button
                    style={{ width: "10rem" }}
                    className="rounded-0"
                    size="large"
                    onClick={() =>
                      window.open(`/detail-place/${item.id}`, "_blank")
                    }
                  >
                    KHÁM PHÁ
                  </Button>
                </Space>
              </Card>
            ))}
          </Space>
        </Row>

        <News />

        <Introduce />
      </div>

      <Divider />

      <Footer />
    </div>
  );
};

export default Homepage;
