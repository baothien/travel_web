import React, { useEffect, useState } from "react";
import { Row, Col, Space, Image } from "antd";
import { PlaceImg } from "../../models/place";
interface IProps {
  images: PlaceImg[];
}

const Images: React.FC<IProps> = ({ images }) => {
  const [arrImgUrl, setArrImgUrl] = useState<string[]>([]);

  useEffect(() => {
    const imgsUrl = images.map((item) => item.url);

    setArrImgUrl(imgsUrl);
  }, [images]);

  return (
    <Image.PreviewGroup>
      <Row
        style={{
          textAlign: "center",
          marginTop: "5rem",
          width: "100%",
          padding: "0 6rem",
        }}
        justify="center"
        gutter={[20, 20]}
      >
        {arrImgUrl.map((item, index) => (
          <Col key={index}>
            <Image style={{ height: "16rem" }} key={index} src={item} />
          </Col>
        ))}
      </Row>
    </Image.PreviewGroup>
  );
};

export default Images;
