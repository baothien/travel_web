import React, { useState } from "react";

import { Popup, Marker } from "react-leaflet";
import { Typography, Row, Col, Space, Button, Image } from "antd";

import { Icon } from "leaflet";
const { Title } = Typography;

import { ArrowRightOutlined, PictureFilled } from "@ant-design/icons";

import iconmarker from "../../../../assets/img/marker-icon.png";

const iconMarker = new Icon({
  iconUrl: iconmarker,
  iconSize: [25, 40],
});

interface IProps {
  id: string;
  position: [number, number];
  title: string;
  thumbnail: string;
  checkInImgUrls: string[];
}

const MapMarker: React.FC<IProps> = ({
  id,
  position,
  title,
  thumbnail,
  checkInImgUrls,
}) => {
  const [visible, setVisible] = useState(false);

  const openDetail = () => {
    window.open(`/detail-place/${id}`);
  };

  return (
    <>
      <Marker
        key={id}
        position={position}
        icon={iconMarker}
        eventHandlers={{
          mouseover: (event) => event.target.openPopup(),
          click: () => {
            setVisible(true);
          },
        }}
      >
        <Popup>
          <Row className="rounded" align="middle">
            <Col className="position-relative">
              <img
                style={{
                  borderRadius: "0.5rem",
                  filter: "brightness(80%)",
                  height: "100%",
                }}
                alt=""
                src={thumbnail}
                width="100%"
              />

              <div
                style={{ position: "absolute", bottom: 0, left: 10 }}
                className="p-2"
              >
                <Space direction="vertical">
                  <Title className="text-white m-0" level={5}>
                    {title}
                  </Title>

                  <Space>
                    <Button
                      className="text-white border-white"
                      size="middle"
                      type="text"
                      shape="circle"
                      icon={<ArrowRightOutlined />}
                      onClick={openDetail}
                    />

                    <Button
                      className="text-white border-white"
                      size="middle"
                      type="text"
                      shape="circle"
                      icon={<PictureFilled />}
                      onClick={() => setVisible(true)}
                    />
                  </Space>
                </Space>
              </div>
            </Col>
          </Row>
        </Popup>
      </Marker>

      <div style={{ display: "none" }}>
        <Image.PreviewGroup
          preview={{ visible, onVisibleChange: (vis) => setVisible(vis) }}
        >
          {checkInImgUrls.map((item, index) => (
            <Image key={index} src={item} />
          ))}
        </Image.PreviewGroup>
      </div>
    </>
  );
};

export default MapMarker;
