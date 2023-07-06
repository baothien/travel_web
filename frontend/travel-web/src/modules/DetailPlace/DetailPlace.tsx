import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import "./DetailPlace.scss";

import { Header, Footer } from "../../components";

import Images from "./Images";
import Summary from "./Summary";
import Video from "./Video";

import { Typography, Row, Col, Space, Button } from "antd";

import apiPlaces from "../../apis/placesApi";
import { Place } from "../../models/place";
import ReviewsContainer from "./components/Review/ReviewsContainer";
import { toast } from "react-toastify";
import { useAppDispatch, useAppSelector } from "../../hooks";
import { selectUser } from "../Authentication/authSlice";
import { checkInActions } from "../CheckIn/checkInSlice";
import checkInApi from "../../apis/checkInApi";

const { Title, Text } = Typography;

interface IDetailPlaceProps {
  replyCmtIdNotified?: string;
}

const DetailPlace: React.FC = (props: IDetailPlaceProps) => {
  const { id } = useParams();

  const dispath = useAppDispatch();

  const user = useAppSelector(selectUser);

  const [isCheckIn, setIsCheckIn] = useState(false);

  useEffect(() => {
    if (user) {
      checkInApi.getCheckInList().then((res) => {
        const placeIds = res.data.rows.map((item) => item.place_id);
        setIsCheckIn(placeIds.includes(id));
      });
    }
  }, [user]);

  const navigate = useNavigate();

  const [renderReview, setRenderReview] = useState<boolean>(false);

  const [place, setPlace] = useState<Place>({
    id: "",
    thumbnail: "",
    name: "",
    place_type_id: "",
    place_type: {
      id: "",
      code: "",
      name: "",
    },
    lat: 0,
    lng: 0,
    address: "",
    place_img: [],
    review: [],
  });

  useEffect(() => {
    const getDetailPlace = async () => {
      if (id) {
        apiPlaces
          .getPlace(id)
          .then((res) => {
            setPlace((prev) => (prev = res.data));
          })
          .catch((err) => console.log(err));
      }
    };
    getDetailPlace();
  }, [id]);

  const clickCheckIn = () => {
    if (user) {
      dispath(
        checkInActions.setBgImgURL(
          "https://roadtovrlive-5ea0.kxcdn.com/wp-content/uploads/2014/09/Venice.Still001.jpeg"
        )
      );

      dispath(checkInActions.setCurrentPlaceId(id));

      navigate("/check-in");
    } else {
      navigate("/login");
    }
  };

  return (
    <div style={{ minHeight: "200rem" }} className="primary-font">
      <Header />

      <Row style={{}} className="mb-5">
        <Row
          style={{
            marginTop: "5rem",
            backgroundImage: `url(${place.thumbnail}), linear-gradient(to bottom, rgba(245, 246, 252, 0.52), rgba(117, 19, 93, 0.73))`,
          }}
          className="bg-img rounded"
          justify="center"
        >
          <Col span={12} className="position-relative">
            <div
              style={{
                width: "70%",
                backgroundColor: "#292D33",
                padding: "2rem",
                zIndex: 1,
                top: 80,
              }}
              className="h-50 ms-5 p-5 position-absolute"
            >
              <Text style={{ fontSize: "1.6rem" }} className="text-white">
                {place.place_type?.name}
              </Text>
              <Title
                style={{ fontSize: "2.8rem" }}
                className="text-white fw-light m-0 mt-4"
              >
                {place.name}
              </Title>
            </div>
            <div
              style={{
                backgroundColor: "#F86449",
                width: "14rem",
                height: "90%",
                top: 40,
                left: 410,
                zIndex: 0,
              }}
              className="position-absolute"
            ></div>
          </Col>

          <Col className="position-relative" span={12}>
            <div
              style={{
                width: "85%",
                bottom: 100,
              }}
              className="bg-white position-absolute shadow p-5"
            >
              <Space direction="vertical" size={25}>
                <Title className="fw-light" level={2}>
                  Khám phá các trải nghiệm du lịch độc đáo và mới lạ
                </Title>
                <Text className="fs-6">
                  Một nơi tuyệt vời để khám phá các địa điểm du lịch thú vị trên
                  khắp thế giới. Trang web cung cấp thông tin về các điểm đến
                  hấp dẫn, các hoạt động và trải nghiệm độc đáo mà bạn có thể
                  tham gia khi đi du lịch
                </Text>
                <Button
                  style={{ width: "10rem", height: "2.5rem" }}
                  className="rounded-0"
                >
                  TÌM HIỂU THÊM
                </Button>
              </Space>
            </div>
          </Col>
        </Row>

        <Row
          style={{ marginTop: "5rem", padding: "0 6rem" }}
          className="w-100"
          justify="center"
          align="middle"
        >
          <Col span={6}>
            <Space style={{}} className="" direction="vertical" size={15}>
              <Text style={{ fontSize: "2rem" }} className="" strong>
                {place.name}
              </Text>
              <Text
                style={{ fontSize: "1.2rem", color: "#FF7424" }}
                className="mt-0"
                strong
              >
                {place.place_type?.name}
              </Text>

              <Text style={{ fontSize: "0.9rem" }}>{place.address}</Text>

              <Button
                style={{ width: "15rem", height: "3rem" }}
                className="bg-black mt-2 rounded-0"
                onClick={clickCheckIn}
              >
                <Text className="text-white fs-6">Check in</Text>
                <i
                  style={{ width: "2rem" }}
                  className="fa-solid fa-arrow-right text-white "
                ></i>
              </Button>
            </Space>
          </Col>

          <Col
            style={{
              height: "60%",
              width: "9rem",
            }}
            className="d-flex justify-content-center"
            span={6}
          >
            <Summary />
          </Col>

          <Col className="position-relative">
            <Video />
          </Col>
        </Row>

        <Images images={place.place_img} />

        <div
          style={{
            backgroundColor: "#F7FAFB",
            margin: "3rem 4rem",
            padding: "2rem 0rem",
          }}
          className="w-100 "
        >
          <Space direction="vertical">
            <Title className="m-0">Bình luận và</Title>
            <Title style={{ color: "#FF7424" }}>Đánh giá</Title>
          </Space>

          <ReviewsContainer isCheckIn={isCheckIn} />
        </div>
      </Row>

      <Footer />
    </div>
  );
};

export default DetailPlace;
