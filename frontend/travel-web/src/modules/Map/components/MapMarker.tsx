import React, { useState } from "react";
import { Link } from "react-router-dom";

import { Popup, Marker } from "react-leaflet";
import { Typography, Row, Col, Space, Button } from "antd";

import { Icon } from "leaflet";
const { Title, Text } = Typography;

import {
  EnvironmentFilled,
  HeartFilled,
  HeartOutlined,
  StarFilled,
} from "@ant-design/icons";

import iconmarker from "../../../assets/img/marker-icon.png";

import apiPlaces from "../../../apis/placesApi";
import { FavPlaceParams } from "../../../models/place";
import { toast } from "react-toastify";

const iconMarker = new Icon({
  iconUrl: iconmarker,
  iconSize: [25, 40],
});

interface IProps {
  id: string;
  position: [number, number];
  title: string;
  thumbnail: string;
  address: string;
  placeType: string;
  isFavorite?: boolean;
}

const MapMarker: React.FC<IProps> = ({
  id,
  position,
  title,
  thumbnail,
  address,
  placeType,
  isFavorite,
}) => {
  const [is_favorite, set_is_favorite] = useState(isFavorite);
  console.log("is favorite", is_favorite);

  const handleFavorite = () => {
    const params: FavPlaceParams = {
      is_favorite: !is_favorite,
      place_id: id,
    };

    console.log(params);

    apiPlaces
      .postFavorite(params)
      .then((res) => {
        toast.success(res.message);
        set_is_favorite((prev) => !prev);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <>
      <Marker
        key={id}
        position={position}
        icon={iconMarker}
        eventHandlers={{
          mouseover: (event) => event.target.openPopup(),
          click: (event) => {
            window.open(`/detail-place/${id}`);
          },
        }}
      >
        <Popup>
          <Space style={{ width: "16rem" }} size={0} direction="vertical">
            <img
              style={{
                padding: 0,
                width: "100%",
                height: "6rem",
                borderTopRightRadius: "5px",
                borderTopLeftRadius: "5px",
              }}
              alt=""
              src={thumbnail}
            />

            <div className="px-3 py-2">
              <Row justify="space-between" align="middle">
                <Col span={16}>
                  <Title className="font-1" level={5}>
                    {title}
                  </Title>
                </Col>

                <Col className="mr-0" span={8}>
                  <Space size={10}>
                    {/* <Link to="/360">
                      <Button
                        size="middle"
                        shape="circle"
                        icon={
                          <EnvironmentFilled style={{ color: "#1A73E8" }} />
                        }
                      />
                    </Link> */}

                    <Button
                      size="middle"
                      shape="circle"
                      onClick={handleFavorite}
                      icon={
                        is_favorite ? (
                          <HeartFilled style={{ color: "#1A73E8" }} />
                        ) : (
                          <HeartOutlined style={{ color: "#1A73E8" }} />
                        )
                      }
                    />
                  </Space>
                </Col>
              </Row>

              {/* <Row justify="start">
                <Space size={5}>
                  <Text className="text-secondary">4.2</Text>
                  <Space size={1}>
                    {[0, 1, 2, 3, 4].map((item) => (
                      <StarFilled style={{ color: "#FCC526" }} />
                    ))}
                  </Space>
                  <Text className="text-secondary">(1.012)</Text>
                </Space>
              </Row> */}

              <Text className="text-secondary">{placeType}</Text>
            </div>
          </Space>
        </Popup>
      </Marker>
    </>
  );
};

export default MapMarker;
