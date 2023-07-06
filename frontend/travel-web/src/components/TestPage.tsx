import { useState } from "react";
import Cropper from "react-easy-crop";
import { Row, Col, Space, Image } from "antd";

const ImageGrid = ({ images }) => {
  return (
    <Row>
      <Image.PreviewGroup>
        {images.map((image) => (
          <Col xs={24} sm={12} md={8} lg={6} key={image.id}>
            <Image src={image.src} alt={image.alt} style={{ height: "5rem" }} />
          </Col>
        ))}
      </Image.PreviewGroup>
    </Row>
  );
};

export function TestPage() {
  return (
    <div className="">
      <ImageGrid
        images={[
          {
            id: 1,
            src: "https://images2.thanhnien.vn/Uploaded/minhnguyet/2022_01_25/blackpink-1-5384.jpg",
            alt: "Image 1",
          },
          // {
          //   id: 2,
          //   src: "https://media.baodansinh.vn/baodansinh/222978561005920256/2021/7/7/blackpink-1625630828685477636583.jpg",
          //   alt: "Image 2",
          // },
          {
            id: 3,
            src: "https://photo-cms-baophapluat.epicdn.me/w840/Uploaded/2023/gznrxgmabianhgzmath/2021_06_15/blackpink-4010.jpeg",
            alt: "Image 3",
          },
          {
            id: 4,
            src: "https://znews-photo.zingcdn.me/w660/Uploaded/unhuuak/2022_11_05/221016_BLACKPINK_ROSE_BALCKPINK_1.jpg",
            alt: "",
          },
          {
            id: 5,
            src: "https://images.unsplash.com/photo-1617634667039-8e4cb277ab46?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bmF0dXJlJTIwbGFuZHNjYXBlfGVufDB8fDB8fA%3D%3D&w=1000&q=80",
            alt: "",
          },
          {
            id: 6,
            src: "https://dvyvvujm9h0uq.cloudfront.net/com/articles/1528189468-127635-mountain-3351653-1280jpg.jpg",
            alt: "",
          },
          {
            id: 7,
            src: "https://shotkit.com/wp-content/uploads/bb-plugin/cache/night-landscape-photography-featured-landscape.jpg",
            alt: "",
          },
        ]}
      />
    </div>
  );
}
